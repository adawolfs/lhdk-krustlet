# WebAssembly + Kubernetes = Krustlet

# WebAssembly
https://www.fermyon.com/wasm-languages/webassembly-language-support

## WASI

## Wasmtime

### Install
```
curl https://wasmtime.dev/install.sh -sSf | bash
```

### Execute module
```
wasmtime --dir . main.wasm
```
## WASMER
### Install
```
curl https://get.wasmer.io -sSfL | sh
```
### Execute it
As you might notice when I fist installed it and try to execute it thows an error
```
➜  wasm+k8s wasmer wasm/c/module/main.wasm            
wasmer: error while loading shared libraries: libtinfo.so.5: cannot open shared object file: No such file or dir$
ctory                                                                                           
```

So I had to google the error click the first result and read a bit.
```
➜  wasm+k8s sudo apt install libtinfo5 
Reading package lists... Done
```

Test all the packages
```
$ wasmer wasm/golang/main.wasm 
Hello from GoLang
$ wasmer wasm/rust/module/target/wasm32-wasi/debug/module.wasm 
Hello from Rust
$ wasmer wasm/c/module/main.wasm 
Hello from C
```

# WASM to OCI

```
wget https://github.com/engineerd/wasm-to-oci/releases/download/v0.1.2/linux-amd64-wasm-to-oci
mv linux-amd64-wasm-to-oci wasm-to-oci
chmod +x wasm-to-oci
sudo cp wasm-to-oci /usr/local/bin
```

```
export CR_PAT=ghp_***********************************
echo $CR_PAT | docker login ghcr.io -u adawolfs --password-stdin
wasm-to-oci push wasm/c/module/main.wasm ghcr.io/adawolfs/wasm-c:latest
```
# Kubernetes
## KinD
I found it cleaner to create a virtual machine with Centos Stream 8 with Docker and KinD

### KinD requires docker "of course!"
```
[root@localhost ~]# sudo yum install docker-ce docker-ce-cli containerd.io docker-compose-plugin --allowerasing
[root@localhost ~]# systemctl start docker
```

```
[root@localhost ~]# vim kind-config.yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  image: kindest/node:v1.21.12@sha256:f316b33dd88f8196379f38feb80545ef3ed44d9197dca1bfd48bcb1583210207

```
```
[root@localhost ~]# kind create cluster --config=kind-config.yaml       
Creating cluster "kind" ...
 ✓ Ensuring node image (kindest/node:v1.21.12) 🖼
 ✓ Preparing nodes 📦                                                                             
 ✓ Writing configuration 📜                                                                        ✓ Starting control-plane 🕹️ 
 ✓ Installing CNI 🔌                                                                               ✓ Installing StorageClass 💾 
Set kubectl context to "kind-kind"
You can now use your cluster with: 

kubectl cluster-info --context kind-kind

Have a nice day! 👋
```
```
[root@localhost ~]# kubectl cluster-info --context kind-kind
Kubernetes control plane is runnin
g at https://127.0.0.1:43015
CoreDNS is running at https://127.0.0.1:43015/api/v1/namespaces/kube-system/services/kube-dns:dns/
proxy
```

```
[root@localhost ~]# kubectl get nodes
NAME                 STATUS   ROLES           AGE     VERSION
kind-control-plane   Ready    control-plane   2m38s   v1.21.12
```

## Krustlet
https://krustlet.dev/

Create bootstrap config
```
bash ./krustlet/bootstrap.sh
```

Start Krustlet
```
KUBECONFIG=~/.krustlet/config/kubeconfig
./krustlet/krustlet-wasi --node-ip 172.17.0.1 --node-name=krustlet --bootstrap-file=~/.krustlet/config/bootstrap.conf
```

Now on another terminal lets approve the certificate 
```
kubectl certificate approve localhost.localdomain-tls
```