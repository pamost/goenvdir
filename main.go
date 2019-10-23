package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Process files in the environment directory
func processEnvDir(envDir string) (valEnv []string, err error) {
	// Calculation of the absolute path
	absPath, err := filepath.Abs(envDir)
	if err != nil {
		log.Printf("absolute path calculation error %v", envDir)
		return nil, err
	}

	// Read contents of dir
	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		log.Printf("can't find env dir %v", absPath)
		return nil, err
	}

	// Read files
	valEnv = make([]string, 0, len(files))

	for _, file := range files {
		// Process only files, exclude directories, symbolic links, etc.
		if !file.Mode().IsRegular() {
			continue
		}

		value, err := ioutil.ReadFile(filepath.Join(absPath, file.Name()))
		if err != nil {
			log.Printf("missing value for variable %v", file.Name())
			return nil, err
		}

		valEnv = append(valEnv, fmt.Sprintf("%s=%s", file.Name(), string(value)))
	}

	return valEnv, nil
}

// Execute an external program with environment variables
func execEnvCmd(dir, prog string) error {
	valEnv, err := processEnvDir(dir)
	if err != nil {
		return err
	}

	cmd := exec.Command(prog)

	cmd.Env = append(os.Environ(), valEnv...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	return nil
}

func main() {
	// goenvdir /path/to/env/dir some_prog
	if len(os.Args) < 3 {
		log.Fatal("arguments are incorrect")
	}

	dir := os.Args[1]
	prog := os.Args[2]

	if err := execEnvCmd(dir, prog); err != nil {
		log.Fatal(err)
	}
}
