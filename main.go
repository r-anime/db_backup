package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"golang.org/x/term"
)

const NumberSoftDelete = 1

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	dbService, dbName, backupType, compressionLevel, backupDir, minSaveSize := ParseArgs()

	date := time.Now().Format(time.DateOnly)
	newBackupName := fmt.Sprintf("%s.pgdump.%s.%s.backup", dbName, date, backupType)
	log.Println("Starting DB backup for", newBackupName)

	filesToSoftDelete, filesToDelete := findBackups(backupDir, minSaveSize, backupType)

	deleteFiles(filesToDelete)
	filesToSoftDelete = softDeleteFiles(filesToSoftDelete)

	tempBackup := strings.TrimSuffix(newBackupName, ".backup") + ".in_progress.backup"
	duration, err := runBackup(dbService, dbName, tempBackup, compressionLevel)
	if err != nil {
		log.Println("Finished DB backup for", newBackupName, "in", duration)
		log.Println("Reverting softDeletedFiles", filesToSoftDelete)
		unSoftDeleteFiles(filesToSoftDelete)
		return
	}

	renameFile(filepath.Join(backupDir, tempBackup), filepath.Join(backupDir, newBackupName))

	deleteFiles(filesToSoftDelete)
	log.Println("Finished DB backup for", newBackupName, "in", duration)
}

func findBackups(dir string, minSaveSize int64, backupType BackupType) (softDeleteFiles []string, hardDeleteFiles []string) {
	files, err := filepath.Glob(filepath.Join(dir, fmt.Sprintf("*.%s.backup", backupType)))
	if err != nil {
		return []string{}, []string{}
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i] > files[j]
	})

	var soft []string
	var hard []string

	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			log.Fatal(err)
		}

		if info.Size() < minSaveSize || len(soft) >= NumberSoftDelete {
			hard = append(hard, file)
		} else {
			soft = append(soft, file)
		}
	}

	return soft, hard
}

func deleteFiles(files []string) {
	for _, file := range files {
		log.Println("Deleting", file)
		err := os.Remove(file)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func softDeleteFiles(files []string) []string {
	for i, file := range files {
		newFile := strings.TrimSuffix(file, ".backup") + ".delete.backup"
		renameFile(file, newFile)
		files[i] = newFile
	}
	return files
}

func unSoftDeleteFiles(files []string) []string {
	for i, file := range files {
		newFile := strings.TrimSuffix(file, ".delete.backup") + ".backup"
		renameFile(file, newFile)
		files[i] = newFile
	}
	return files
}

func renameFile(oldFile string, newFile string) {
	log.Println("Renaming", oldFile, "->", newFile)
	err := os.Rename(oldFile, newFile)
	if err != nil {
		log.Fatal(err)
	}
}

func runBackup(dbContainer string, dbName string, tempBackup string, compressionLevel int8) (duration time.Duration, err error) {
	start := time.Now()

	command := []string{"docker", "exec"}
	if term.IsTerminal(int(os.Stdin.Fd())) {
		command = append(command, "-it")
	}
	command = append(command,
		dbContainer,
		"pg_dump",
		"-U", dbName,
		fmt.Sprintf("--compress=zstd:level=%d,long=1", compressionLevel),
		"-Fc",
		"-f", "/mnt/backup/"+tempBackup,
		dbName,
	)
	humanCommand := strings.Join(command, " ")

	log.Printf("Running command: '%s'\n", humanCommand)
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	if err != nil {
		log.Println("Error running", humanCommand)
		log.Println(err)
	}
	return time.Since(start), err
}
