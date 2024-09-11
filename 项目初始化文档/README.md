# 项目初始化文档
本项目是对现有日记账进行重构

## 功能特性
本项目的主要功能是

## 架构设计
### 目录结构设计
```text
|- githooks       // git hooks  
    |- commit-msg // git提交钩子  
    |- pre-commit // git分支切换钩子  
|- scripts        // 执行各种构建，安装，分析等操作的脚本   
    |- make-rule  // makefile 规则文件 
```
 
## windows 安装make
使用choco命令行安装make, 需要使用管理员打开powershell

```shell
# 设置chocolatey 安装目录, 默认安装到 C:\ProgramData\chocolatey
[System.Environment]::SetEnvironmentVariable("ChocolateyInstall", "D:\ProgramData\chocolatey", "Machine")

# 安装choco
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

# 查看choco版本
choco -v

# 安装make
choco install make
```











