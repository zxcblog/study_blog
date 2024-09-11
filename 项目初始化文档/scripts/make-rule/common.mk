# 通用的配置文件，用来配置所有makefile文件都会使用到的配置信息

SHELL := /bin/bash
GIT := git

# 包含common.mk文件的位置, 是用该位置做为根目录
COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/../.. && pwd -P))
endif

# Makefile配置，在执行命令时会应为切换目录而输出提示信息，这里设置不输出
ifndef V
MAKEFLAGS += --no-print-directory
endif

# 默认最小覆盖率60%
ifeq ($(origin COVERAGE),undefined)
COVERAGE := 60
endif


# 复制 githooks 文件夹下的钩子到 .git/hooks 文件夹下
# 在执行make 命令时会自动执行
COPY_GITHOOK:=$(shell cp -f githooks/* .git/hooks/)

# linux 命令
FIND := find . ! -path './third_party/*' ! -path './vendor/*'
XARGS := xargs -r
GIT := git

# 指定需要安装的工具: BLOCKER_TOOLS, CRITICAL_TOOLS, TRIVIAL_TOOLS.
# BLOCKER_TOOLS 必须，可能导致make all不能正常运行的工具
# CRITICAL_TOOLS 重要，可能导致一些重要命令失效的工具
# TRIVIAL_TOOLS 可选，有没有都可以的工具
#BLOCKER_TOOLS ?= gsemver  go-junit-report  addlicense  codegen
#CRITICAL_TOOLS ?= swagger mockgen gotests git-chglog github-release coscmd go-mod-outdated protoc-gen-go cfssl go-gitlint
#TRIVIAL_TOOLS ?= depth go-callvis gothanks richgo rts kube-score
BLOCKER_TOOLS ?= golines goimports golangci-lint go-junit-report
CRITICAL_TOOLS ?=
TRIVIAL_TOOLS ?=