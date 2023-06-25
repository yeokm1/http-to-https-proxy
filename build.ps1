$env:GOOS="darwin"
$env:GOARCH="amd64"
go build -o http-to-https-proxy-darwin-amd64

$env:GOOS="darwin"
$env:GOARCH="arm64"
go build -o http-to-https-proxy-darwin-arm64

$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o http-to-https-proxy-linux-amd64

$env:GOOS="linux"
$env:GOARCH="arm"
go build -o http-to-https-proxy-linux-arm

$env:GOOS="linux"
$env:GOARCH="arm64"
go build -o http-to-https-proxy-linux-arm64

$env:GOOS="windows"
$env:GOARCH="386"
go build -o http-to-https-proxy-win-386.exe

$env:GOOS="windows"
$env:GOARCH="amd64"
go build -o http-to-https-proxy-win-amd64.exe
