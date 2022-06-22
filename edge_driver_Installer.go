package wdinstaller

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/daicmod/extractzip"
)

func EdgeDriverInstaller(opt Option) error {
	if err := opt.init("MicrosoftWebDriver"); err != nil {
		return err
	}

	opt = initEdgeOption(opt)

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
	if err := os.Rename(filepath.Join(opt.DestPath, "msedgedriver.exe"), filepath.Join(opt.DestPath, opt.filename)); err != nil {
		return err
	}

	// Remove zip file
	os.RemoveAll(compressedPath)

	return nil
}

func initEdgeOption(opt Option) Option {
	// get driver & browser version
	if opt.os == "win" {
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
		opt.driverURL = opt.driverURL + opt.browserVersion + "/edgedriver_" + opt.os + opt.arch + ".zip"
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
		opt.driverURL = opt.driverURL + opt.browserVersion + "/edgedriver_" + opt.os + "64.zip"
	}
	return opt
}
