#!/bin/bash

set -e

INSTALL_DIR="/usr/local/bin/"
FILENAME=""
RELEASE="0.1.2"

flags() {
  while [[ $# -gt 0 ]]; do
    key="$1"

    case $key in
    -d | --install-dir)
      INSTALL_DIR="$2"
      shift
      shift
      ;;
    -r | --release)
      RELEASE="$2"
      shift
      shift
      ;;
    *)
      echo "Unrecognized argument $key"
      exit 1
      ;;
    esac
  done
}

filename() {
  local OS

  OS=$(uname -s)
  CPU=$(uname -m)

  echo "[DENOMON]"
  echo "Thanks for using denomon."

  echo ""
  echo "> Instaling version $RELEASE"
  echo ""

  if [ "$OS" == "Linux" ]; then
    FILENAME="denomon_${RELEASE}_${OS}_${CPU}"
  elif [ "$OS" == "Darwin" ]; then
    FILENAME="denomon_${RELEASE}_${OS}_${CPU}"
  else
    echo "OS $OS is not supported."
    echo "If you think that's a bug - please file an issue to https://github.com/muhibbudins/denomon/issues"
    exit 1
  fi
}

download() {
  echo ""

  if [ "$RELEASE" == "latest" ]; then
    URL=https://github.com/muhibbudins/denomon/releases/latest/download/$FILENAME.tar.gz
  else
    URL=https://github.com/muhibbudins/denomon/releases/download/v$RELEASE/$FILENAME.tar.gz
  fi
  
  DOWNLOAD_DIR=$(mktemp -d)

  echo "Downloading $URL"

  mkdir -p $INSTALL_DIR &>/dev/null
  curl -L $URL -o $DOWNLOAD_DIR/$FILENAME.tar.gz

  if [ 0 -ne $? ]; then 
    echo "Download failed.  Check that the release/filename are correct."
    exit 1
  fi;

  echo ""
  echo "Installing denomon"
  echo ""

  tar -xzvf $DOWNLOAD_DIR/$FILENAME.tar.gz -C $DOWNLOAD_DIR
  mv $DOWNLOAD_DIR/denomon $INSTALL_DIR/denomon
  chmod u+x $INSTALL_DIR/denomon

  echo ""
  echo "Successfully installing denomon"
  echo ""
  echo "`denomon --help`"
  echo ""
}

check() {
  echo "Checking dependencies for the installation script:"
  echo ""

  if hash curl 2>/dev/null; then
    echo "> curl is available"
  else
    echo "> curl is missing!"
    SHOULD_EXIT="true"
  fi

  if hash tar 2>/dev/null; then
    echo "> tar is available"
  else
    echo "> tar is missing!"
    SHOULD_EXIT="true"
  fi

  if [ "$SHOULD_EXIT" = "true" ]; then
    exit 1
  fi
}

flags "$@"
filename
check
download