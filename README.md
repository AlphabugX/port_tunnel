# port_tunnel
Msfconsole也有类似功能，但是效果很差。此工具可以反向端口转发，用于复杂的内网环境，让内网臣服于红队吧！（2021年06月20日，昨天在部门内部做了技术分享，今天就来发布吧。）

```go
Usage of ./port_tunnel_darwin:
  -l string
    	0.0.0.0:222 [本地网卡监听端口] (default "0.0.0.0:222")
  -r string
    	8.8.8.8:1234 [NC、CobaltStrike、Metasploit等服务监听端口] (default "8.8.8.8:1234")
```

![port_tunnel](https://image.3001.net/images/20210620/1624182547678.png)

## 基础功能演示

https://www.yuque.com/docs/share/ff7e9ffb-98c9-48ab-81b0-b985f85cbb68?# 《port_tunnel (Go开发)》

## 场景1

### 不出网打MS17-010 反向上线

https://www.yuque.com/docs/share/c0fe12b2-7a76-4c4e-9b8e-adec2c2ceef7?# 《MS17-010 port_tunnel上线案例》

# 更新日志

- 2021年06月20日17:41:00 这个工具，源码将在1.0版本发布。