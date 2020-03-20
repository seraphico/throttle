# throttle

## 项目介绍：

    一个基于TC的网络限速简化工具,基于Golang开发

## 编译：

```bash
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -ldflags '-s -w' \-o throttle app/app.go 
```

## 令行帮助手册

### 主命令

``` shell script
    throttle --help
```

``` shell script
    NAME:
       throttle - throttle -h,--help
    
    USAGE:
       throttle [global options] command [command options] [arguments...]
    
    VERSION:
       release 1.0.0
    
    AUTHOR:
       seraphic <dongdong1260@gmail.com>
    
    COMMANDS:
       add      创建规则
       ls       查看规则
       rm       删除规则
       help, h  Shows a list of commands or help for one command
    
    GLOBAL OPTIONS:
       --help, -h     show help
       --version, -v  print the version
    
    COPYRIGHT:
       ©2010-2020 Seraphic Corporation,All Rights Reserved

```
---
### 子命令
```shell script
    [root@c7node03 ~]# ./throttle  ls  --help 

```
```shell script
    NAME:
       throttle ls - 查看规则
    
    USAGE:
       throttle ls [command options] [arguments...]
    
    OPTIONS:
       --interface value, -i value  指定查询规则的设备 (default: "eth0")
       --network value, -t value    指定查询的IP地址段 (default: "all")
       
    [root@c7node03 ~]# 

```

## License

MIT License

Copyright (c) 2020 邢东东

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

## 特别鸣谢
   - [Cli](https://github.com/urfave/cli)
   - [ASCII Table Writer](https://github.com/olekukonko/tablewriter)
