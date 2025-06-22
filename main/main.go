package main

import (
	"Go/main/shapes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
)

type FileMetadata struct {
	Name    string
	Size    int64
	Content string
}

type FileError struct {
	filename string
	Err      error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("error with file %s: %v", e.filename, e.Err)
}

func ReadFileMetadata(filename string) (*FileMetadata, *FileError) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, &FileError{
			filename: filename,
			Err:      err}
	}
	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return nil, &FileError{
			filename: filename,
			Err:      err}
	}

	fileInfo, err := file.Stat()

	if err != nil {
		return nil, &FileError{
			filename: filename,
			Err:      err}
	}

	return &FileMetadata{
		Name:    fileInfo.Name(),
		Size:    fileInfo.Size(),
		Content: string(data),
	}, nil
}

func add(a int, b int) int {
	return a + b
}

func printShapeInfo(s shapes.Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	newUser := &User{Username: "john_doe", Email: "john.doe@example.com"}
	json, err := json.Marshal(newUser)
	if err == nil {
		fmt.Println(string(json))
	}
	filename := "main/example.txt"

	metadata, err := ReadFileMetadata(filename)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("File Name: %s\n", metadata.Name)
	fmt.Printf("File Size: %d bytes\n", metadata.Size)
	fmt.Println("File Content:")
	fmt.Println(metadata.Content)

	c := shapes.Circle{Radius: 5}
	r := shapes.Rectangle{Width: 4, Height: 6}

	printShapeInfo(c)
	printShapeInfo(r)

	fmt.Println(reflect.TypeOf(c).Kind())

}
