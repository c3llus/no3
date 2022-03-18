package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	username      = os.Getenv("USER")
	baseDirectory = fmt.Sprintf(baseDirectoryFormat, username)
	animalTypes   = []string{
		land,
		water,
	}
	fileNameFromDir = regexp.MustCompile(regexFileName)
)

// Create base directory if not exist,
// download zip folder if not exist
func CreateBaseDirectory() {
	baseDirectory = fmt.Sprintf(baseDirectoryFormat, username)
	os.MkdirAll(baseDirectory, os.ModePerm)

	// copy ZIP
	source := strToPath(".", zipFileName)
	destination := strToPath(baseDirectory, zipFileName)
	input, err := ioutil.ReadFile(source)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(destination, input, 0644)
	if err != nil {
		panic(err)
	}
}

// JAWABAN A
// Create sub directory
func CreateSubDirectory() {
	for i, animalType := range animalTypes {
		os.MkdirAll(baseDirectory+"/"+animalType, os.ModePerm)
		if i < len(animalTypes)-1 {
			time.Sleep(3 * time.Second) // sesuai requirement soal
		}

	}
}

// JAWABAN B
func UnzipFile() {
	archive, err := zip.OpenReader(strToPath(baseDirectory, zipFileName))
	if err != nil {
		return
	}
	defer archive.Close()

	for _, f := range archive.File {

		// JAWABAN C
		// Seperate land and water animal; delete non land or water
		filePath := AnimalSorter(f.FileInfo().Name())
		if filePath == "" {
			continue
		}
		fmt.Println("PATH: " + filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(baseDirectory)+string(os.PathSeparator)) {
			return
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}
}

// JAWABAN D
func RemoveBirds() {
	files, _ := ioutil.ReadDir(strToPath(baseDirectory, land))
	for _, file := range files {
		fileName := fileNameFromDir.FindStringSubmatch(file.Name())[0]
		if IsBird(&fileName) {

			fmt.Println("RM: " + file.Name())
			err := os.Remove(strToPath(baseDirectory, land, file.Name()))
			if err != nil {
				panic(err)
			}
		}
	}
}

// JAWABAN E
func ListAnimalsToTXT() {

	fileTXT, err := os.Create(listTXTFileName)
	if err != nil {
		panic(err)
	}
	defer fileTXT.Close()

	for _, animalType := range animalTypes {
		files, _ := ioutil.ReadDir(strToPath(baseDirectory, animalType))
		for _, file := range files {
			fileTXT.WriteString(GetStringTxt(file))
		}
	}
}
