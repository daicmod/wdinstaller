package wdinstaller

import (
	"errors"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

type Option struct {
	DestPath string

	os             string
	arch           string
	filename       string
	driverVersion  string
	browserVersion string
	driverURL      string
}

func (opt *Option) init(filename string) error {
	if opt.DestPath == "" {
		return errors.New("")
	}
	if opt.os == "" {
		opt.os = runtime.GOOS
	}
	if opt.arch == "" {
		if strings.Contains(runtime.GOARCH, "64") {
			opt.arch = "64"
		} else {
			opt.arch = "32"
		}
	}
	if opt.os == "windows" {
		opt.os = "win"
		opt.filename = filename + ".exe"
	} else {
		opt.os = runtime.GOOS
		opt.filename = filename
	}

	if filename == "MicrosoftWebDriver" {
		opt.driverURL = "https://msedgedriver.azureedge.net/"
		if opt.os != "win" {
			opt.arch = "64"
		}
	} else if filename == "chromedriver" {
		opt.driverURL = "https://chromedriver.storage.googleapis.com/"
		if opt.os == "win" {
			opt.arch = "32"
		} else {
			opt.arch = "64"
		}
	} else {
		return errors.New("NOT SUPPORTED DRIVER")
	}
	return nil
}

func downloadDriver(filepath string, url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)

	return err
}
