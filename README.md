# wdinstaller

## Webdriver installer for Golang

This is the webdriver installer module.
If webdriver is out of date or does not match your browser version, install the latest webdriver.

For now support:

[ChromeDriver](https://chromedriver.chromium.org/)  
[GeckoDriver](https://github.com/mozilla/geckodriver)  
[MicrosoftEdgeDriver](https://developer.microsoft.com/en-us/microsoft-edge/tools/webdriver/)  

## Usage

Use it with [agouti](https://github.com/sclevine/agouti).

```go
func main() {
  opt := wdinstaller.Option{DestPath: "./"}
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
}
  
```

## Instalation

```
$ go install github.com/daicmod/wdinstaller@latest
```

## License

MIT
