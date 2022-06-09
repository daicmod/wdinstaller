package wdinstaller

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEdgeDriverInstaller(t *testing.T) {
	opt := Option{}
	opt.DestPath = "./"

	if err := EdgeDriverInstaller(opt); err != nil {
		t.Errorf("EdgeDriverInstaller error: %v", err)
	}

	if _, err := os.Stat(filepath.Join(opt.DestPath, "MicrosoftWebDriver.exe")); os.IsNotExist(err) {
		t.Error("Not Found MicrosoftWebDriver.exe", filepath.Join(opt.DestPath, "MicrosoftWebDriver.exe"))
	}

	if err := os.Remove(filepath.Join(opt.DestPath, "MicrosoftWebDriver.exe")); err != nil {
		panic(err)
	}
}
