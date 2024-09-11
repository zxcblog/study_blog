# 安装项目依赖工具, 需要安装的工具在common.mk中定义

TOOLS ?=$(BLOCKER_TOOLS) $(CRITICAL_TOOLS) $(TRIVIAL_TOOLS)

# 安装所有工具
.PHONY: tools.install
tools.install: $(addprefix tools.install., $(TOOLS))

# 使用%通配符，对所有install工具进行统一输出和控制
.PHONY: tools.install.%
tools.install.%:
	@echo "==================> Installing $*"
	@$(MAKE) install.$*

# 使用%通配符校验工具是否安装，没有安装则进行安装
.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

# 安装 格式化工具
.PHONY: install.golines
install.golines:
	@$(GO) install github.com/segmentio/golines@latest


# 安装 包导入工具
.PHONY: install.goimports
install.goimports:
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

# 安装静态代码检查工具
.PHONY: install.golangci-lint
install.golangci-lint:
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3
	@golangci-lint completion bash > $(HOME)/.golangci-lint.bash
	@if ! grep -q .golangci-lint.bash $(HOME)/.bashrc; then echo "source \$$HOME/.golangci-lint.bash" >> $(HOME)/.bashrc; fi

# 安装 将 go test 输出转换为xml工具
.PHONY: install.go-junit-report
install.go-junit-report:
	@$(GO) install github.com/jstemmer/go-junit-report@latest

# 安装检查依赖是否过时工具
.PHONY: install.go-mod-outdated
install.go-mod-outdated:
	@$(GO) install github.com/psampaz/go-mod-outdated@latest
