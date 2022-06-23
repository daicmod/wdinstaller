package main

import (
	"log"
	"os"

	"github.com/daicmod/wdinstaller"
	"github.com/sclevine/agouti"
)

func main() {
	p, _ := os.Getwd()
	opt := wdinstaller.Option{DestPath: p}

	// Sample Edge
	if err := wdinstaller.EdgeDriverInstaller(opt); err != nil {
		panic(err)
	}

	edgeDriver := agouti.EdgeDriver(agouti.Browser("msedge"))
	defer edgeDriver.Stop()
	if err := edgeDriver.Start(); err != nil {
		panic(err)
	}
	edgePage, _ := edgeDriver.NewPage()
	edgePage.Navigate("https://www.google.com")

	// Sample Chrome
	if err := wdinstaller.ChromeDriverInstaller(opt); err != nil {
		panic(err)
	}

	chromeDriver := agouti.ChromeDriver(agouti.Browser("chrome"))
	defer chromeDriver.Stop()
	if err := chromeDriver.Start(); err != nil {
		panic(err)
	}
	chromePage, _ := chromeDriver.NewPage()
	chromePage.Navigate("https://www.google.com")

	// Sample Firefox
	if err := wdinstaller.GeckoDriverInstaller(opt); err != nil {
		panic(err)
	}

	geckoDriver := agouti.GeckoDriver(agouti.Browser("firefox"))
	defer geckoDriver.Stop()
	if err := geckoDriver.Start(); err != nil {
		log.Println("start", err)
	}
	firefoxPage, _ := geckoDriver.NewPage()
	firefoxPage.Navigate("https://www.google.com")
}
