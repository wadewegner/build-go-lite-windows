package main

import (
"fmt"
"io"
"net/http"
"os"
"strings"
"os/exec"
)

func main() {
	srcFile := "go1.4.1.windows-386.zip"
	url := "https://storage.googleapis.com/golang/" + srcFile
	goPath := "go"
	goZipName := "go1.4.1.windows-386-waw.zip"

	os.RemoveAll(srcFile)
	os.RemoveAll(goPath)
	os.RemoveAll(goZipName)

	downloadFromUrl(url)

	_, err := exec.LookPath("7z")
	if err != nil {
		fmt.Println("Make sure 7zip is install and include your path.")
		return
	}

	commandString := fmt.Sprintf("unzip %s", srcFile)
	executeCmd(commandString)

	for _,element := range alwaysRemove {

		fullPath := goPath + "/" + element
		
		fmt.Println("Removing file: " + fullPath)
		err = os.RemoveAll(goPath + "/" + element)
		
		if err != nil {
			fmt.Println(err)
		}
	}

	commandString = fmt.Sprintf(`7z a -tzip %s %s`, goZipName, goPath)
	executeCmd(commandString)
}

func executeCmd(commandString string) {
	commandSlice := strings.Fields(commandString)
	fmt.Println(commandString)
	c := exec.Command(commandSlice[0], commandSlice[1:]...)
	err := c.Run()

	if err != nil {
		fmt.Println(err)
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

