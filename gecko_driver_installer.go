package wdinstaller

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/daicmod/extractzip"
)

func GeckoDriverInstaller(opt Option) error {
	if err := opt.init("geckodriver"); err != nil {
		return err
	}

	opt.baseURL = "https://github.com/mozilla/geckodriver/releases/download/"
	driverLatestVersion, _ := getGeckoDriverVersion("https://github.com/mozilla/geckodriver/releases/latest")
	opt.driverURL = opt.baseURL + driverLatestVersion + "/geckodriver-" + driverLatestVersion + "-" + opt.os + opt.arch
	if opt.os == "win" {
		opt.arch = "32"
		opt.browserVersion = getVersionExec("powershell", "(Get-Item \"C:\\Program Files\\Mozilla Firefox\\firefox.exe\").VersionInfo")
		opt.driverVersion = getVersionExec("powershell", filepath.Join(opt.DestPath, opt.filename), "--version")
		opt.driverVersion = "V" + opt.driverVersion
		opt.driverURL = opt.driverURL + ".zip"
	} else {
		opt.arch = "64"
		opt.browserVersion = getVersionExec("firefox", "--version")
		opt.driverVersion = getVersionExec("geckodriver", "--version")
		opt.driverVersion = "V" + opt.driverVersion
		opt.driverURL = opt.driverURL + ".tar.gz"
	}

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
	if opt.os == "win" {
		err := extractzip.ExtractFromZip(opt.filename, compressedPath, opt.DestPath)
		if err != nil {
			return err
		}
	} else {
		err := extractzip.ExtractFromTar(opt.filename, compressedPath, opt.DestPath)
		if err != nil {
			return err
		}
	}
	// Remove zip file
	os.RemoveAll(compressedPath)

	return nil
}

func getGeckoDriverVersion(url string) (string, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, _ := client.Get(url)
	defer res.Body.Close()

	tmp := strings.Split(res.Header.Get("Location"), "/")
	driverLatestVersion := tmp[len(tmp)-1]

	return driverLatestVersion, nil
}
