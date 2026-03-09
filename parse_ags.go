package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/pflag"
)

type BackupType = string

func ParseArgs() (dockerContainer string, dbName string, backupType BackupType, compressionLevel int8, backupDir string, minSaveSize int64) {
	dockerContainerPtr := pflag.StringP("docker-container", "d", "modbot-db", "the docker (not compose) container name for the db")
	dbNamePtr := pflag.StringP("db-name", "n", "r_anime", "the name of the database")
	backupTypePtr := pflag.StringP("backup-type", "t", "manual", "the backup series label (yearly, monthly, weekly, daily, manual)")
	compressionLevelPtr := pflag.Int8P("compression-level", "c", 3, "the compression level (1 - 19)")
	backupDirPtr := pflag.StringP("backup-dir", "b", "/srv/prod/backup/", "the directory the backups are located in")
	minSaveSizePtr := pflag.Uint16P("min-save-size", "m", 2*1024, "the minimum file size to consider for saving in MiB")

	pflag.Parse()

	runValidations(*dockerContainerPtr, *backupTypePtr, *compressionLevelPtr, *backupDirPtr)

	return *dockerContainerPtr, *dbNamePtr, *backupTypePtr, *compressionLevelPtr, *backupDirPtr, int64(*minSaveSizePtr) * 1024 * 1024
}

func runValidations(dockerContainer string, backupType BackupType, compressionLevel int8, backupDir string) {
	var errs []error

	runDockerContainerValidation(&errs, dockerContainer)
	runBackupTypeValidation(&errs, backupType)
	runCompressionLevelValidation(&errs, compressionLevel)
	runBackupDirValidation(&errs, backupDir)

	if len(errs) > 0 {
		log.Println("Invalid flags")
		for _, err := range errs {
			log.Println(err)
		}
		os.Exit(1)
	}
}

func runDockerContainerValidation(errs *[]error, dockerContainer string) {
	cmd := exec.Command("docker", "ps", "-q", "--filter", "name=^"+dockerContainer+"$", "--filter", "status=running")

	out, err := cmd.Output()
	if err != nil {
		*errs = append(*errs, err)
		return
	}

	if strings.TrimSpace(string(out)) == "" {
		*errs = append(*errs, fmt.Errorf("%s is not running", dockerContainer))
	}
}

func runBackupTypeValidation(errs *[]error, backupType BackupType) {
	switch backupType {
	case "yearly", "monthly", "weekly", "daily", "manual":
	default:
		*errs = append(*errs, fmt.Errorf("%s is not a valid backup type ((yearly, monthly, weekly, daily, manual])", backupType))
	}
}

func runCompressionLevelValidation(errs *[]error, compressionLevel int8) {
	if compressionLevel < 1 {
		*errs = append(*errs, fmt.Errorf("%d is too low of a compression level (1-19)", compressionLevel))

	} else if compressionLevel > 19 {
		*errs = append(*errs, fmt.Errorf("%d is too high of a compression level (1-19)", compressionLevel))
	}
}

func runBackupDirValidation(errs *[]error, backupDir string) {
	info, err := os.Stat(backupDir)
	if err != nil {
		*errs = append(*errs, fmt.Errorf("%s is not a valid directory: %w", backupDir, err))
		return
	}

	if !info.IsDir() {
		*errs = append(*errs, fmt.Errorf("%s is not a valid directory", backupDir))
	}
}
