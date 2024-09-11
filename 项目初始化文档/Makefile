
.DEFAULT_GOAL := all

.PHONY: all
all: tidy format lint cover


# 项目包名，用于包自动导入，和go.mod文件 module 相同
ROOT_PACKAGE=项目初始化文档

# 引入其他makefile文件
include scripts/make-rule/common.mk
include scripts/make-rule/golang.mk
include scripts/make-rule/tools.mk

## tidy: go mod tidy
.PHONY: tidy
tidy:
	@$(GO) mod tidy

## tools: 安装需要的工具
.PHONY: tools
tools:
	@$(MAKE) tools.install

## format: 格式化代码，在 git bash 中进行操作
.PHONY: format
format: tools.verify.golines tools.verify.goimports
	@echo "===========> Formating codes"
	$(FIND) -type f -name '*.go' | $(XARGS) gofmt -s -w
	$(FIND) -type f -name '*.go' | $(XARGS) goimports -w -local $(ROOT_PACKAGE)
	$(FIND) -type f -name '*.go' | $(XARGS) golines -w --max-len=120 --no-reformat-tags --shorten-comments --ignore-generated .
	$(GO) mod edit -fmt
	$(GIT) add -A .

## lint: 静态代码检查 windows使用命令 "make lint ROOT_DIR=项目路径"
.PHONY: lint
lint:
	@$(MAKE) go.lint

## test: 单元测试
.PHONY: test
test:
	@$(MAKE) go.test

## cover: 运行单元测试并获取测试覆盖率
.PHONY: cover
cover:
	@$(MAKE) go.test.cover

## check-updates: 检查包依赖是否过时
.PHONY: check-updates
check-updates:
	@$(MAKE) go.updates

## help: 显示帮助信息
.PHONY: help
help: Makefile
	 @echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	 @sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@sed -n 's/^#end//p' $<

#end
#endOptions:
#end  V                设置为1启用详细生成, 默认值为0