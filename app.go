package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"log"
)

func main() {
	url := "https://storage.googleapis.com/golang/go1.4.1.windows-amd64.zip"
	srcFile := "go1.4.1.windows-amd64.zip"
	dest := "."

	cleanup(srcFile)
	cleanup("go")

	downloadFromUrl(url)

	unzip(srcFile, dest)
}

func cleanup(fileOrFolder) {

	err := os.Remove(fileOrFolder)

      if err != nil {
          fmt.Println(err)
          return
      }

}

func downloadFromUrl(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}

func unzip(src, dest string) error {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer r.Close()

    for _, f := range r.File {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()

        fpath := filepath.Join(dest, f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(fpath, f.Mode())
        } else {
            var fdir string
            if lastIndex := strings.LastIndex(fpath,string(os.PathSeparator)); lastIndex > -1 {
                fdir = fpath[:lastIndex]
            }

            err = os.MkdirAll(fdir, f.Mode())
            if err != nil {
                log.Fatal(err)
                return err
            }
            f, err := os.OpenFile(
                fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer f.Close()

            _, err = io.Copy(f, rc)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

