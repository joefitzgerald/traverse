# Detailed Installation Instructions

<!-- @import "[TOC]" {cmd="toc" depthFrom=2 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [Detailed Installation Instructions](#detailed-installation-instructions)
  - [Manual install](#manual-install)
    - [Download the latest binary](#download-the-latest-binary)
    - [wget](#wget)
  - [MacOS / Linux via Homebrew install](#macos--linux-via-homebrew-install)
  - [Docker image pull](#docker-image-pull)
    - [One-shot container use](#one-shot-container-use)
    - [Run container commands interactively](#run-container-commands-interactively)
  - [Go install](#go-install)

<!-- /code_chunk_output -->


## Manual install

### [Download the latest binary](https://github.com/joefitzgerald/traverse/releases/latest)

### wget
Use wget to download the pre-compiled binaries:

```bash
wget https://github.com/joefitzgerald/traverse/releases/download/${VERSION}/${BINARY}.tar.gz -O - |\
  tar xz && mv ${BINARY} /usr/bin/traverse
```

For instance, VERSION=v0.3.1 and BINARY=traverse_${VERSION}_linux_amd64

## MacOS / Linux via Homebrew install

Using [Homebrew](https://brew.sh/)  

```bash
brew tap joefitzgerald/tap
brew install traverse
```

## Docker image pull

```bash
docker pull ghcr.io/joefitzgerald/traverse
```

### One-shot container use

```bash
docker run --rm -v "${PWD}":/workdir ghcr.io/joefitzgerald/traverse [flags]
```


### Run container commands interactively

```bash
docker run --rm -it -v "${PWD}":/workdir --entrypoint sh ghcr.io/joefitzgerald/traverse
```

It can be useful to have a bash function to avoid typing the whole docker command:

```bash
traverse() {
  docker run --rm -i -v "${PWD}":/workdir ghcr.io/joefitzgerald/traverse "$@"
}
```


## Go install

```bash
GO111MODULE=on go get github.com/joefitzgerald/traverse/cmd/traverse
```