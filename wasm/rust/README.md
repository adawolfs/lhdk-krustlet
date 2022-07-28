# RUST

## Create package
```
cargo new main
```

## Add wasm32-wasi target
```
rustup target add wasm32-wasi
```

## Compile to wasi
```
cargo build --target wasm32-wasi
```

## Extract module
```
mv target/wasm32-wasi/main.wasm target/main.wasm
```