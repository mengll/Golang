package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	path := os.Args[1]
	path_dir := strings.Split(path, "\\")
	file_name := strings.Join(path_dir[len(path_dir)-1:], "\\")
	unzip(path, "./")

	ucompress_path := strings.Join(path_dir[:len(path_dir)-1], "\\")
	CompressZip(path, "./temp/"+file_name, "i am the best")

	//删除解压目录
	os.RemoveAll(ucompress_path)
}

func CompressZip(path_name, file_name, comment string) {
	File, _ := os.Create(file_name)
	PS := strings.Split(path_name, "\\")
	PathName := strings.Join(PS[:len(PS)-1], "\\")
	os.Chdir(PathName)
	Path := PS[len(PS)-1]
	defer File.Close()
	Zip := zip.NewWriter(File)
	defer Zip.Close()

	walk := func(Path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		Src, _ := os.Open(Path)
		defer Src.Close()

		h := &zip.FileHeader{Name: Path, Method: zip.Deflate, Flags: 0x800}
		FileName, _ := Zip.CreateHeader(h)
		Zip.SetComment(comment)

		io.Copy(FileName, Src)
		Zip.Flush()
		return nil
	}
	if err := filepath.Walk(Path, walk); err != nil {
		fmt.Println(err)
	}
	Zip.Flush()
}

//解压文件
func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}
