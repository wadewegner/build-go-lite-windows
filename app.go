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
	"strconv"
)

func main() {
	srcFile := "go1.4.1.windows-386.zip"
	url := "https://storage.googleapis.com/golang/" + srcFile
	
	dest := "."
	goPath := "go"

	err := os.Remove(srcFile)
	err = os.RemoveAll(goPath)

	if err != nil {
		fmt.Println(err)
	}

	downloadFromUrl(url)

	unzip(srcFile, dest)

	for index,element := range alwaysRemove {

		fullPath := goPath + "/" + element
		
		fmt.Println("Removing file " + strconv.Itoa(index) + ": " + fullPath)
		err = os.RemoveAll(goPath + "/" + element)
	
		if err != nil {
			fmt.Println(err)
		}
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

var alwaysRemove = []string{
	".git",
	".gitattributes",
	".gitignore",
	"CONTRIBUTING.md",
	"VERSION.cache",
	"favicon.ico",
	"robots.txt",

	".hgignore", // waw - added
	".hgtags", // waw - added

	"api",
	"blog", // waw - added
	"include",
	"lib",
	"misc",
	"pkg/obj",
	"test",

	"bin/dist",

	"pkg/GOOS_GOARCH/cmd",

	"pkg/tool/GOOS_GOARCH/dist",
	"pkg/tool/GOOS_GOARCH/fix",
	"pkg/tool/GOOS_GOARCH/nm",
	"pkg/tool/GOOS_GOARCH/objdump",
	"pkg/tool/GOOS_GOARCH/yacc",

	"src/cmd/dist",
	"src/cmd/fix",
	"src/cmd/nm",
	"src/cmd/objdump",
	"src/cmd/yacc",

	"src/cmd/5a",
	"src/cmd/5g",
	"src/cmd/5l",
	"src/cmd/6a",
	"src/cmd/6g",
	"src/cmd/6l",
	"src/cmd/8a",
	"src/cmd/8g",
	"src/cmd/8l",
	"src/cmd/9a",
	"src/cmd/9g",
	"src/cmd/9l",
	"src/cmd/cc",
	"src/cmd/gc",

	"src/cmd/link",
	"src/cmd/ld",
	"src/cmd/pack", // ?

	"src/cmd/go",

	"src/all.bash",
	"src/all.bat",
	"src/all.rc",
	"src/androidtest.bash",
	"src/clean.bash",
	"src/clean.bat",
	"src/clean.rc",
	"src/make.Dist",
	"src/make.bash",
	"src/make.bat",
	"src/make.rc",
	"src/nacltest.bash",
	"src/race.bash",
	"src/race.bat",
	"src/run.bash",
	"src/run.bat",
	"src/run.rc",

	"src/lib9",
	"src/libbio",
	"src/liblink",

	// TODO(adg): option to preserve pprof
	"pkg/tool/GOOS_GOARCH/pprof",
	"src/cmd/pprof",

	// TODO(adg): option to preserve gofmt
	"bin/gofmt",
	"src/cmd/gofmt",

	// TODO(adg): option to preserve docs
	"doc",
}

