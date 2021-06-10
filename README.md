# gin-static
gin 开发的静态网页，配置好静态资源目录即可

```shell
./main
```

或者添加参数：

```shell
./main -w ./website -s ./website/static -o 8040
```

 默认静态html文件地址 website，默认 静态资源目录public，默认端口 8040



注意：每次更新静态页面文件名或目录，需要重新启动一次应用，重新生成路由



如果存在中文名称，在linux系统中可能会出现问题，使用convmv工具转换文件目录名

```shell
yum install convmv

convmv -f gbk -t utf-8 -r --notest .
```

