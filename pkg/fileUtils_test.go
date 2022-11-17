package pkg

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetAllFiles(t *testing.T) {
	files, _ := GetAllDirs("/Users/Mr_J/Downloads")
	for _, file := range files {
		fmt.Println(file)
	}

	contains := strings.Contains("/Users/Mr_J/Desktop/synchron-space/conf/conf.yaml", "/Users/Mr_J/Desktop/")
	fmt.Println(contains)
}
