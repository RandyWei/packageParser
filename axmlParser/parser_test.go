package axmlParser

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	var filename = "/Users/wei/Downloads/仕途联线_1.0.0.apk"

	listener := new(AndroidListener)
	_, err := ParseApk(filename, listener)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Init package is", listener.PackageName,
		"Activity is", listener.ActivityName)
}
