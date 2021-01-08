# goinfras-tool
一个配合goinfras构建项目的自动生成工具

### 1.安装

```
// 拉取并安装
go get -u github.com/bb-orz/gt

```

### 2.查看命令：

> gt -h

```
NAME:
   gt - A generation tool of go app scaffold which base on bb-orz/goinfras.

USAGE:
   gt [option] [command] [args]

VERSION:
   1.0.0

COMMANDS:
   init     Go Web Application Initialization
   model    Add core model
   domain   Add core domain in project
   service  Add Application Service
   restful  Add Restful API
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

```

#### 2.1 Command Init

init 命令用于初始化应用脚手架，目前有sample/account|grpc|micro 四种模板可选

- Sample：https://github.com/bb-orz/goapp-sample 简单的restful应用脚手架
- Account：https://github.com/bb-orz/goapp-acount 实现基本账户接口的restful应用脚手架
- Grpc：https://github.com/bb-orz/goapp-grpc 简单的grpc应用脚手架
- Micro：https://github.com/bb-orz/goapp-micro 简单的go-micro rpc微服务应用脚手架

```
NAME:
  gt init - Go Web Application Initialization

USAGE:
  gt init [--name|-n=][project_name] [--git|-g=true|false] [--mod|-m=true|false] 

DESCRIPTION:
  The init command create a new go web application in current directory，this command will generate some necessary folders and files, crete a project.

OPTIONS:
  --name value, -n value    [[--name|-n=]ProjectName] (default: "goapp")
  --sample value, -s value  [--sample|-s=[sample|account|grpc|micro]] (default: "sample")
  --mod, -m                 [--mod|-m=true|false] (default: true)
  --git, -g                 [--git|-g=true|false] (default: true)
  --help, -h                show help (default: false)
```

#### 2.2 Command Model
model 命令用于根据数据库表schema生成相应的go struct model 和 dto

```
NAME:
   gt model - Add core model

USAGE:
   gt model [command options] [arguments...] ...

DESCRIPTION:
   The model command create a new core model with go struct，this command will generate some necessary files or dir in core directory .

OPTIONS:
   --driver value, -D value           (default: "mysql")
   --host value, -H value             (default: "localhost")
   --port value, -P value             (default: 3306)
   --database value, -d value         
   --table value, -t value            
   --user value, -u value             (default: "dev")
   --password value, -p value         (default: "123456")
   --output_path value, -o value      (default: "./core")
   --dto_output_path value, -O value  (default: "./dtos")
   --formatter value, -f value        (default: "gorm")
   --help, -h                         show help (default: false)

```

#### 2.3 Command Domain

domain 命令用于根据领域驱动开发创建相应的领域模块代码，如传入数据库连接参数，会相应生成简单的表curd 代码

```
NAME:
   gt domain - Add core domain in project

USAGE:
   gt domain [--name|-n=][DomainName] ...

DESCRIPTION:
   The domain command create a new core domain with go struct，this command will generate some necessary files or dir in core directory .

OPTIONS:
   --name value, -n value             (default: "example")
   --driver value, -D value           (default: "mysql")
   --host value, -H value             (default: "localhost")
   --port value, -P value             (default: 3306)
   --database value, -d value         
   --table value, -t value            
   --user value, -u value             (default: "dev")
   --password value, -p value         (default: "123456")
   --output_path value, -o value      (default: "./core")
   --dto_output_path value, -O value  (default: "./dtos")
   --formatter value, -f value        (default: "gorm")
   --help, -h                         show help (default: false)
   

```

#### 2.4 Command Service

service 命令用于创建服务层代码范式及相应的dto范式

```
NAME:
   gt service - Add Application Service

USAGE:
   gt service [--name|-n=][ServiceName] ...

DESCRIPTION:
   The service command create a new service go interface，this command will generate some necessary files in service directory.

OPTIONS:
   --name value, -n value                   (default: "example")
   --version value, -v value                (default: "V1")
   --interface_output_path value, -o value  (default: "./services")
   --implement_output_path value, -c value  (default: "./core")
   --dto_output_path value, -d value        (default: "./dtos")
   --help, -h                               show help (default: false)
  
```

#### 2.5 Command Restful

restful 命令用于创建restful接口层的代码范式

```
USAGE:
   gt restful [--name|-n=][RestfulName] ...

DESCRIPTION:
   The restful command create a new restful api with go struct，this command will generate some necessary files or dir in restful director .

OPTIONS:
   --name value, -n value         (default: "example")
   --engine value, -e value       (default: "gin")
   --output_path value, -o value  (default: "./restful")
   --help, -h                     show help (default: false)
   

```