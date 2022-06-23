package wdinstaller

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/daicmod/extractzip"
)

func EdgeDriverInstaller(opt Option) error {
	if err := opt.init("MicrosoftWebDriver"); err != nil {
		return err
	}
	fileNamerow := "msedgedriver"
	opt.baseURL = "https://msedgedriver.azureedge.net/"
	if opt.os == "win" {
		fileNamerow = fileNamerow + ",exe"
		opt.browserVersion = getVersionExec("powershell", "(Get-Item \"C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe\").VersionInfo")
		opt.driverVersion = getVersionExec("powershell", filepath.Join(opt.DestPath, opt.filename), "--version")
	} else {
		opt.arch = "64"
		opt.browserVersion = getVersionExec("microsoft-edge", "--version")
		opt.driverVersion = getVersionExec("MicrosftWebDriver", "--version")
	}
	opt.driverURL = opt.baseURL + opt.browserVersion + "/edgedriver_" + opt.os + opt.arch + ".zip"

	// No download required
	if opt.browserVersion != "" && opt.browserVersion == opt.driverVersion {
		return nil
	}

	// Download driver zip/tar
	compressedPath := filepath.Join(opt.DestPath, strings.Split(opt.driverURL, "/")[len(strings.Split(opt.driverURL, "/"))-1])
	if err := downloadDriver(compressedPath, opt.driverURL); err != nil {
		return err
	}

	// Extract driver
	err := extractzip.ExtractFromZip(fileNamerow, compressedPath, opt.DestPath)
	if err != nil {
		return err
	}
	if err := os.Rename(filepath.Join(opt.DestPath, fileNamerow), filepath.Join(opt.DestPath, opt.filename)); err != nil {
		return err
	}
	// Remove zip file
	os.RemoveAll(compressedPath)

	return nil
}
