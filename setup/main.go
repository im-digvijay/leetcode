package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getProjectRootDirectory() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	projectRootDirRegex := `(?:^/(?:[^/]+/)*leetcode/)|(?:^(?:[A-Z]:\\)?(?:[^/\\]+\\)*leetcode\\)`
	compiledRegex := regexp.MustCompile(projectRootDirRegex)
	return compiledRegex.FindString(cwd), nil
}

func main() {
	var problemName string
	fmt.Print("Enter Problem name (this will become directory name): ")
	_, err := fmt.Scanln(&problemName)
	if err != nil {
		log.Fatal("Not able to read from standard input\n", err)
	}

	projectRootDir, err := getProjectRootDirectory()
	if err != nil {
		log.Fatal("Unable to identify project root directory.\n", err)
	}

	problemDir := projectRootDir + "golang/" + problemName

	ok, err := exists(problemDir)
	if err != nil {
		log.Fatal("Some error occurred.\n", err)
	}
	if !ok {
		fmt.Printf("Problem directory %s doesn't exists.\nCreating...\n", problemDir)
		if err := os.Mkdir(problemDir, os.ModeDir|os.ModePerm); err != nil {
			log.Fatal("Unable to create directory.\n", err)
		}
	} else {
		fmt.Printf("Problem directory %s exists.\n", problemDir)
	}
}
