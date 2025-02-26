![alt icon](https://github.com/xiao-k8/OllamaScanner/blob/main/image/icon.webp)
<img src="https://github.com/xiao-k8/OllamaScanner/blob/main/image/icon.webp" width="35%" style="text-align: center">
# 说明

主要用于扫描内网或者互联网开放的Ollama服务（默认11434端口）

---
# 使用

使用方法
> OllamaScanner -ip 192.168.1.1/24 -p 11400-11500 -t 100

结果默认输出到当前目录，开放Ollama端口的结果输出在`Success_xxx.log`

| 命令  | 说明                                                                        |
| --- | ------------------------------------------------------------------------- |
| -ip | 设置需要扫描的IP地址或IP段，支持nmap风格输入，例:-ip 192.168.1.1/24,10.10.10.1-254,172.16.1.* |
| -p  | 设置需要扫描的端口范围，默认11434，也支持nmao风格输入，例: -p 11400-11450,8080,1-1000             |
| -t  | 设置扫描线程，默认20                                                               |
| -h  | 显示帮助信息                                                                    |
