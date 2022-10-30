# MMMoney

> 一个关于 money 的工具集, 寄托一个美好的心愿(make more money)~

---

## Install 

```shell
# go install 
go install github.com/evercyan/mmmoney

# homebrew
brew install evercyan/tap/mmmoney
```

---

## Usage

```shell
mmmoney help
# mmmoney: make more money
# 
# Usage:
#   mmmoney [command]
# 
# Available Commands:
#   completion  Generate the autocompletion script for the specified shell
#   help        Help about any command
#   loan        计算房贷本金利息
#   tax         计算五险一金和个税
# 
# Flags:
#   -h, --help      help for mmmoney
#   -v, --version   version for mmmoney
# 
# Use "mmmoney [command] --help" for more information about a command.
```

- 计算房贷本金利息 `mmmoney loan`

![loan](https://cdn.jsdelivr.net/gh/evercyan/repository/resource/d9/d969c3c96ece6224ef547cbf0893fb4b.png)

- 计算五险一金和个税 `mmmoney tax`

![loan](https://cdn.jsdelivr.net/gh/evercyan/repository/resource/7d/7d6c6276120a7579a691dbeb59dc04e5.png)
