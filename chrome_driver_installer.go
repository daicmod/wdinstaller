package wdinstaller

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/daicmod/extractzip"
)

func ChromeDriverInstaller(opt Option) error {
	if err := opt.init("chromedriver"); err != nil {
		return err
	}

	opt.baseURL = "https://chromedriver.storage.googleapis.com/"
	if opt.os == "win" {
		opt.arch = "32"
		opt.browserVersion = getVersionExec("powershell", "(Get-Item \"C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe\").VersionInfo")
		opt.driverVersion = getVersionExec("powershell", filepath.Join(opt.DestPath, opt.filename), "--version")
	} else {
		opt.arch = "64"
		opt.browserVersion = getVersionExec("google-chrome", "--version")
		opt.driverVersion = getVersionExec("chromedriver", "--version")
	}

	driverLatestVersion, _ := getChromeDriverVersion("http://chromedriver.storage.googleapis.com/LATEST_RELEASE_" + strings.Split(opt.browserVersion, ".")[0])
	opt.driverURL = opt.baseURL + driverLatestVersion + "/chromedriver_" + opt.os + opt.arch + ".zip"

	// No download required
	if (opt.browserVersion != "" && opt.browserVersion == opt.driverVersion) ||
		opt.driverVersion == driverLatestVersion {
		return nil
	}

	// Download driver zip/tar
	compressedPath := filepath.Join(opt.DestPath, strings.Split(opt.driverURL, "/")[len(strings.Split(opt.driverURL, "/"))-1])
	if err := downloadDriver(compressedPath, opt.driverURL); err != nil {
		return err
	}

	// Extract driver
	err := extractzip.ExtractFromZip(opt.filename, compressedPath, opt.DestPath)
	if err != nil {
		return err
	}
	os.RemoveAll(compressedPath)

	return nil
}

func getChromeDriverVersion(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	b, _ := io.ReadAll(res.Body)
	driverLatestVersion := string(b)
	return driverLatestVersion, nil
}
