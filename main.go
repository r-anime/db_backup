package main

import (
	"fmt"
)

func main() {
	dbService, dbName, backupType, compressionLevel := ParseArgs()

	fmt.Println("Test")
	fmt.Println("dbService", dbService)
	fmt.Println("dbName", dbName)
	fmt.Println("backupType", backupType)
	fmt.Println("compressionLevel", compressionLevel)
}
