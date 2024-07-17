<h1 align="center">Wocker</h1>

<!-- File: README.md -->
<!-- Author: YJ -->
<!-- Email: yj1516268@outlook.com -->
<!-- Created Time: 2024-06-30 16:24:15 -->

---

<p align="center">
  <a href="https://github.com/YHYJ/wocker/actions/workflows/release.yml"><img src="https://github.com/YHYJ/wocker/actions/workflows/release.yml/badge.svg" alt="Go build and release by GoReleaser"></a>
</p>

---

## Table of Contents

<!-- vim-markdown-toc GFM -->

* [Install](#install)
  * [一键安装](#一键安装)
* [Usage](#usage)
* [Compile](#compile)
  * [当前平台](#当前平台)
  * [交叉编译](#交叉编译)
    * [Linux](#linux)

<!-- vim-markdown-toc -->

---

<!---------------------------------------->
<!--                     _              -->
<!-- __      _____   ___| | _____ _ __  -->
<!-- \ \ /\ / / _ \ / __| |/ / _ \ '__| -->
<!--  \ V  V / (_) | (__|   <  __/ |    -->
<!--   \_/\_/ \___/ \___|_|\_\___|_|    -->
<!---------------------------------------->

---

一个 Docker 包装器

## Install

### 一键安装

```bash
curl -fsSL https://raw.githubusercontent.com/YHYJ/wocker/main/install.sh | sudo bash -s
```

## Usage

- `image`子命令

  管理 docker 镜像，可以指定镜像或交互式操作

- `volume`子命令

  管理 docker 卷，可以指定卷或交互式操作

- `version`子命令

  查看程序版本信息

- `help`子命令

  查看程序帮助信息

## Compile

### 当前平台

```bash
go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/wocker/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/wocker/general.BuildTime=`date +%s` -X github.com/yhyj/wocker/general.BuildBy=$USER" -o build/wocker main.go
```

### 交叉编译

使用命令`go tool dist list`查看支持的平台

#### Linux

```bash
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/wocker/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/wocker/general.BuildTime=`date +%s` -X github.com/yhyj/wocker/general.BuildBy=$USER" -o build/wocker main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64
