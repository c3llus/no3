package main

import (
	"fmt"
	"io/fs"
	"strings"
)

func strToPath(args ...string) (resp string) {

	for i, v := range args {
		resp = resp + v
		if i < len(args)-1 {
			resp = resp + "/"
		}
	}

	return resp
}

func AnimalSorter(fileName string) (fileDest string) {

	for _, animalType := range animalTypes {
		if strings.Contains(strings.ToLower(fileName), animalType) {
			return strToPath(baseDirectory, animalType, fileName)
		}
	}

	return ""
}

func IsBird(fileName *string) (resp bool) {

	if strings.Contains(strings.ToLower(*fileName), bird) {
		return true
	}

	return false
}

func GetFilePermission(file fs.FileInfo) string {
	return fmt.Sprintf("%04o", file.Mode().Perm())
}

func GetStringTxt(file fs.FileInfo) string {
	return fmt.Sprintf(writeTXTStringFormat, GetFilePermission(file), file.Name()+newLine)
}
