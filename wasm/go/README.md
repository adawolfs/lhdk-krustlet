# Go

# TinyGo

## Install on Debian
This is already done on devcontainer
```
wget https://github.com/tinygo-org/tinygo/releases/download/v0.23.0/tinygo_0.23.0_amd64.deb
sudo dpkg -i tinygo_0.23.0_amd64.deb
```

## Create module
```
mkdir src target
touch src/main.go
```

## Build
```
tinygo build -wasm-abi=generic -target=wasi -o target/main.wasm src/main.go
```

## Build for web
```
GOOS=js GOARCH=wasm go build -o static/main.wasm src/web/main.go
```
    ### Extras
- https://binx.io/2022/04/22/golang-webassembly/

