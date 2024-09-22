package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	subModuleName := flag.String("subModuleName", "sys", "set submodule name")
	// 替换为你的文件夹路径
	folderPath := "rpc/" + *subModuleName + "/proto"
	// 新文件的路径
	outputFilePath := "rpc/" + *subModuleName + "/" + *subModuleName + ".proto"

	files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var fileContents []byte
	var startContents []byte
	sb := strings.Builder{}

	sb.WriteString("syntax = \"proto3\";\n")
	sb.WriteString("package " + *subModuleName + "client;\n")
	sb.WriteString("option go_package = \"./" + *subModuleName + "client\";\n\n")

	for _, file := range files {
		if file.IsDir() || file.Name() == "main.go" {
			continue
		}

		fileName := filepath.Join(folderPath, file.Name())
		fileData, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Error reading file %s: %s\n", fileName, err)
			continue
		}
		fileData = fileData[20:]

		_, err = sb.Write(fileData)
		if err != nil {
			panic(err)
		}
	}
	startContents = append(startContents, fileContents...)

	err = os.WriteFile(outputFilePath, []byte(sb.String()), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Merged files to", outputFilePath)
}
