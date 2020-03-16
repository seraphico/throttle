# throttle

#### 编译：
```go
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -ldflags '-s -w' \-o throttle app/app.go  && scp throttle root@192.168.1.8:/root/
```

#### 令行帮助手册
##### 主命令
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
##### 子命令
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