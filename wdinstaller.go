package wdinstaller

import (
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/daicmod/extractzip"
)

type Option struct {
	DestPath string
	OS       string
	Arch     string

	filename       string
	driverVersion  string
	browserVersion string
	driverURL      string
}

func (opt *Option) init(filename string) error {
	if opt.DestPath == "" {
		return errors.New("")
	}
	if opt.OS == "" {
		opt.OS = runtime.GOOS
	}
	if opt.Arch == "" {
		if strings.Contains(runtime.GOARCH, "64") {
			opt.Arch = "64"
		} else {
			opt.Arch = "32"
		}
	}
	if opt.OS == "windows" {
		opt.OS = "win"
		opt.filename = filename + ".exe"
	} else {
		opt.OS = runtime.GOOS
		opt.filename = filename
	}

	if filename == "MicrosoftWebDriver" {
		opt.driverURL = "https://msedgedriver.azureedge.net/"
	}
	return nil
}

func EdgeDriverInstaller(opt Option) error {
	if err := opt.init("MicrosoftWebDriver"); err != nil {
		return err
	}

	// get driver & browser version
	if opt.OS == "win" {
		out, _ := exec.Command("powershell", "(Get-Item \"C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe\").VersionInfo").Output()
		if len(out) != 0 {
			r := regexp.MustCompile(`[\d*\.]+`)
			opt.browserVersion = r.FindAllStringSubmatch(string(out), 1)[0][0]
		} else {
			opt.browserVersion = ""
		}
		out, _ = exec.Command("powershell", "(Get-Item "+filepath.Join(opt.DestPath, opt.filename)+").VersionInfo").Output()
		if len(out) != 0 {
			r := regexp.MustCompile(`[\d*\.]+`)
			opt.driverVersion = r.FindAllStringSubmatch(string(out), 1)[0][0]
		} else {
			opt.driverVersion = ""
		}
		opt.driverURL = opt.driverURL + opt.browserVersion + "/edgedriver_" + opt.OS + opt.Arch + ".zip"
	} else {
		out, _ := exec.Command("microsoft-edge", "--version").Output()
		if len(out) != 0 {
			r := regexp.MustCompile(`[\d*\.]+`)
			opt.browserVersion = r.FindAllStringSubmatch(string(out), 1)[0][0]
		} else {
			opt.browserVersion = ""
		}
		out, _ = exec.Command("MicrosftWebDriver", "--version").Output()
		if len(out) != 0 {
			r := regexp.MustCompile(`[\d*\.]+`)
			opt.driverVersion = r.FindAllStringSubmatch(string(out), 1)[0][0]
		} else {
			opt.driverVersion = ""
		}
		opt.driverURL = opt.driverURL + opt.browserVersion + "/edgedriver_" + opt.OS + "64.zip"
	}

	// No download required
	if opt.browserVersion == opt.driverVersion {
		return nil
	}

	// Download driver zip/tar
	compressedPath := filepath.Join(opt.DestPath, strings.Split(opt.driverURL, "/")[len(strings.Split(opt.driverURL, "/"))-1])
	if err := downloadDriver(compressedPath, opt.driverURL); err != nil {
		return err
	}

	// Extract driver
	err := extractzip.ExtractFromZip("msedgedriver.exe", compressedPath, opt.DestPath)
	if err != nil {
		return err
	}
	if err := os.Rename("msedgedriver.exe", opt.filename); err != nil {
		return err
	}
	os.RemoveAll(compressedPath)

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
