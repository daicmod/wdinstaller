package wdinstaller

import (
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"regexp"
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
	baseURL        string
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

func getVersionExec(command string, args ...string) string {
	out, _ := exec.Command(command, args...).Output()
	if len(out) != 0 {
		r := regexp.MustCompile(`[\d*\.]+`)
		return r.FindAllStringSubmatch(string(out), 1)[0][0]
	} else {
		return ""
	}
}
