# Tessor.network - The Fantasy Sports Blockchain

![GitHub release](https://img.shields.io/github/release/tessornetwork/fury.svg)
[![CircleCI](https://circleci.com/gh/tessornetwork/fury/tree/master.svg?style=shield)](https://circleci.com/gh/tessornetwork/fury/tree/master)
[![codecov](https://codecov.io/gh/tessornetwork/fury/branch/master/graph/badge.svg)](https://codecov.io/gh/tessornetwork/fury)
[![Go Report Card](https://goreportcard.com/badge/github.com/tessornetwork/fury)](https://goreportcard.com/report/github.com/tessornetwork/fury)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/tessornetwork/fury.svg)
![GitHub](https://img.shields.io/github/license/tessornetwork/fury.svg)

Tessor.network, as known as *The Fantasy Sports Blockchain* is the easiest way for players to manage their 
Fantasy Sports Teams using blockchain technology. 

## References
* [Official website](https://tessor.network/)
* [Documentation](https://docs.tessor.network/)
* [Telegram group](https://t.me/tessornetwork)

## Version

Current Software Version is `v0.1.0-FanfuryTestnet.1`

Current Chain Version is `Furytestnet-1000`
## Quick Start

To compile our software (Debian/Ubuntu SO)

```bash
apt update && apt upgrade -y
apt install -y git gcc make unzip
snap install --classic go


echo 'export GOPATH="$HOME/go"' >> ~/.profile
echo 'export PATH="$GOPATH/bin:$PATH"' >> ~/.profile
echo 'export PATH="$PATH:/snap/bin"' >> ~/.profile

source ~/.profile

git remote clone https://github.com/tessornetwork/fury.git
git checkout tags/v4.2.0
go mod verify
make install
```

You can verify your installation with

```bash
furyd version
```

