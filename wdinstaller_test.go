package wdinstaller

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func prepareOption() Option {
	p, _ := os.Getwd()
	opt := Option{}
	opt.DestPath = p
	return opt
}

func prepareFileName(fileName string) string {
	if runtime.GOOS == "windows" {
		return fileName + ".exe"
	}
	return fileName
}

func TestEdgeDriverInstaller(t *testing.T) {
	opt := prepareOption()
	fileName := prepareFileName("MicrosoftWebDriver")

	if err := EdgeDriverInstaller(opt); err != nil {
		t.Errorf("EdgeDriverInstaller error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(opt.DestPath, fileName)); os.IsNotExist(err) {
		t.Error("Not Found MicrosoftWebDriver", filepath.Join(opt.DestPath, fileName))
	}

	if err := os.Remove(filepath.Join(opt.DestPath, fileName)); err != nil {
		panic(err)
	}
}

func TestChromeDriverInstaller(t *testing.T) {
	opt := prepareOption()
	fileName := prepareFileName("chromedriver")

	if err := ChromeDriverInstaller(opt); err != nil {
		t.Errorf("ChromeDriverInstaller error: %v", err)
	}

	if _, err := os.Stat(filepath.Join(opt.DestPath, fileName)); os.IsNotExist(err) {
		t.Error("Not Found chromedriver", filepath.Join(opt.DestPath, fileName))
	}

	if err := os.Remove(filepath.Join(opt.DestPath, fileName)); err != nil {
		panic(err)
	}
}

func TestGeckoDriverInstaller(t *testing.T) {
	opt := prepareOption()
	fileName := prepareFileName("geckodriver")

	if err := GeckoDriverInstaller(opt); err != nil {
		t.Errorf("GeckoDriverInstaller error: %v", err)
	}

	if _, err := os.Stat(filepath.Join(opt.DestPath, fileName)); os.IsNotExist(err) {
		t.Error("Not Found geckodriver", filepath.Join(opt.DestPath, fileName))
	}

	if err := os.Remove(filepath.Join(opt.DestPath, fileName)); err != nil {
		panic(err)
	}
}
