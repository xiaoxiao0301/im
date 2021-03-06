# WebIM

## 简介
> vue + golang + xorm + websocket 实现单聊和私聊

本项目主要实现了以下功能点：

- 用户的注册与登录
- 图片上传
- 单聊
- 群聊
- 创建群聊 
- 添加好友
- 添加群
- 单聊和群聊中包括
    - 发送文字信息
    - 发送图片信息
    - 发送emoj表情
    - 发送语音聊天信息
    

## 打包部署
<span style="color:red;">视图文件和资源文件不需要打包，只需要复制到打包后的对应目录下就行</span>

### windows
build.bat
```bash
::remove dir
rd /s/q release
::make dir 
md release
::go build -ldflags "-H windowsgui" -o chat.exe
go build -o chat.exe
::
COPY chat.exe release\
COPY favicon.ico release\favicon.ico
::
XCOPY asset\*.* release\asset\  /s /e
XCOPY view\*.* release\view\  /s /e 
```

### linux
build.sh
```bash
#!/bin/sh
rm -rf ./release
mkdir  release
go build -o chat
chmod +x ./chat
cp chat ./release/
cp favicon.ico ./release/
cp -arf ./asset ./release/
cp -arf ./view ./release/
```



## Nginx反向代理
```nginx
	upstream wsbackend {
			server 192.168.0.102:8080;
			server 192.168.0.100:8080;
			hash $request_uri;
	}
	map $http_upgrade $connection_upgrade {
    default upgrade;
    ''      close;
	}
    server {
	  listen  80;
	  server_name localhost;
	  location / {
	   proxy_pass http://wsbackend;
	  }
	  location ^~ /chat {
	   proxy_pass http://wsbackend;
	   proxy_connect_timeout 500s;
       proxy_read_timeout 500s;
	   proxy_send_timeout 500s;
	   proxy_set_header Upgrade $http_upgrade; # 表示这是websocket
       proxy_set_header Connection "Upgrade";
	  }
	 }

}
```

## 扩展
xorm默认输出的json格式化时间戳格式形式是`2021-06-23T14:54:22+08:00` 下面来优化一下
1. 自定义一个time.Time类型
2. 实现MarshalJSON方法
下面来看代码
```go
import "time"

type Mytimes time.Time

func (this Mytimes) MarshalJSON() ([]byte, error)  {
	if time.Time(this).IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + time.Time(this).Format("2006-01-02 15:04:05") + `"`), nil
}

```   
```go
type TestTime struct {
	Id int64 `xorm:"bigint(20) pk autoincr" json:"id"`
	Name string `xorm:"varchar(10)" json:"name"`
	CreatedAt Mytimes `xorm:"timestamp created" json:"created_at"`
	UpdatedAt time.Time `xorm:"timestamp updated" json:"updated_at"`
}
```