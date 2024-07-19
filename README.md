<h1 align="center">Wocker</h1>
<h3 align="center">一个 Docker 包装器</h3>

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

* [适配](#适配)
* [安装](#安装)
  * [一键安装](#一键安装)
  * [编译安装](#编译安装)
    * [当前平台](#当前平台)
    * [交叉编译](#交叉编译)
* [用法](#用法)

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

## 适配

- Linux: 适配
- macOS: 适配
- Windows: 适配

## 安装

### 一键安装

```bash
curl -fsSL https://raw.githubusercontent.com/YHYJ/wocker/main/install.sh | sudo bash -s
```

也可以从 [GitHub Releases](https://github.com/YHYJ/wocker/releases) 下载解压后使用

### 编译安装

#### 当前平台

```bash
go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/wocker/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/wocker/general.BuildTime=`date +%s` -X github.com/yhyj/wocker/general.BuildBy=$USER" -o build/wocker main.go
```

#### 交叉编译

> 使用命令`go tool dist list`查看支持的平台
>
> Linux 和 macOS 使用命令`uname -m`，Windows 使用命令`echo %PROCESSOR_ARCHITECTURE%` 确认系统架构
>
> - 例如 x86_64 则设 GOARCH=amd64
> - 例如 aarch64 则设 GOARCH=arm64
> - ...

设置如下系统变量后使用 [编译安装](#编译安装) 的命令即可进行交叉编译：

- CGO_ENABLED: 不使用 CGO，设为 0
- GOOS: 设为 linux, darwin 或 windows
- GOARCH: 根据当前系统架构设置

## 用法

- `image`子命令

  管理 docker 镜像，可以指定镜像或交互式操作

- `volume`子命令

  管理 docker 数据卷，可以指定卷或交互式操作

- `version`子命令

  查看程序版本信息

- `help`子命令

  查看程序帮助信息
