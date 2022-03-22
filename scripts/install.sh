#!/bin/sh -e

if [ -n "${HELM_LINTER_PLUGIN_NO_INSTALL_HOOK}" ]; then
    echo "Development mode: not downloading versioned release."
    exit 0
fi

version="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
echo "Downloading and installing helm-teller v${version} ..."

arch=""
case $(uname -m) in
  x86_64)
    arch="amd64"
    ;;
  armv6*)
    arch="armv6"
    ;;
  armv7*)
    arch="armv7"
    ;;
  aarch64 | arm64)
    arch="arm64"
    ;;
  *)
    echo "Failed to detect target architecture"
    exit 1
    ;;
esac

if [ "$(uname)" = "Darwin" ]; then
    url="https://github.com/spectralops/helm-teller/releases/download/v${version}/helm-teller_${version}_darwin_${arch}.tar.gz"
elif [ "$(uname)" = "Linux" ] ; then
    url="https://github.com/spectralops/helm-teller/releases/download/v${version}/helm-teller_${version}_linux_${arch}.tar.gz"
else
    url="https://github.com/spectralops/helm-teller/releases/download/v${version}/helm-teller_${version}_windows_${arch}.tar.gz"
fi

echo "Downloding helm-teller from: $url"

mkdir -p "bin"
mkdir -p "releases/v${version}"

curl -sSL "${url}" -o "releases/v${version}.tar.gz"
tar xzf "releases/v${version}.tar.gz" -C "releases/v${version}"
mv "releases/v${version}/helm-teller" "bin/teller" || \
    mv "releases/v${version}/helm-teller.exe" "bin/teller"