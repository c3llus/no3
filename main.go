package main

func main() {

	// Create directories
	CreateBaseDirectory()
	CreateSubDirectory()

	// Unzip file ACCORDINGLY
	UnzipFile()

	// Remove files that contain "bird"
	RemoveBirds()

	// List animals
	ListAnimalsToTXT()
}
