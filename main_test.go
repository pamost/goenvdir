package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestProcessEnvDir(t *testing.T) {
	envdir, test := envdir()
	defer os.RemoveAll(envdir) // clean up

	result, err := processEnvDir(envdir)
	if err != nil {
		return
	}

	if !reflect.DeepEqual(result, test) {
		t.Errorf("error in process env dir %v not equal %v", result, test)
	}
}

func TestExecEnvCmd(t *testing.T) {
	envdir, _ := envdir()
	defer os.RemoveAll(envdir) // clean up

	err := execEnvCmd(envdir, "env")
	if err != nil {
		t.Errorf("error in execute cmd %v", err)
	}
}

// Create tmp env dir and files
func envdir() (envdir string, test []string) {
	items := []struct {
		file  string
		value string
	}{
		{"A_ENV", "123"},
		{"B_VAR", "another_val"},
		{"C_INT", "val456"},
	}

	// Create tmp dir
	envdir, err := ioutil.TempDir("", "envdir")
	if err != nil {
		log.Fatal(err)
	}

	// Create tmp files
	test = make([]string, 0, len(items))

	for _, env := range items {
		tmpFile := filepath.Join(envdir, env.file)
		test = append(test, fmt.Sprintf("%s=%s", env.file, env.value))

		if err := ioutil.WriteFile(tmpFile, []byte(env.value), 0666); err != nil {
			log.Fatal(err)
		}
	}

	return envdir, test
}
