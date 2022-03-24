#!/usr/bin/env bash
GOPATH=$(pwd)
export GOBIN=/usr/local/go/bin
export PATH=$PATH:$GOBIN
export PGO_BASEOS=debian
export PGO_IMAGE_PREFIX=zhonghl003
export PGO_VERSION=3.0.0
export CCP_PGVERSION=14
export CCP_BACKREST_VERSION=2.36
export CCP_PG_FULLVERSION=14.2
export export CCP_POSTGIS_VERSION=3.1
export CCPROOT=$GOPATH
export PGO_IMAGE_TAG=${PGO_BASEOS}-${PGO_VERSION}