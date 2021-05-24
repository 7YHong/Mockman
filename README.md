## Mockman
一个简单的mockserver
- 运行

直接双击，或者
~~~
bin/mockman api
~~~

- 参数

~~~
# bin/mockman api --help
~~~
~~~
Command Options:
  -a, --addr            Listen to the specified address
  -d, --daemon          Run in the background
  -l, --logname         Set log file name as this
~~~
例如
~~~
bin/mockman api -d -a :8888 -l mylogname
~~~

- http访问路径
~~~
GET /hello
ANY /echo  
ANY /api/*
~~~

- 运行配置
~~~
修改目录下.env文件，或者直接设置环境变量
~~~

- mock接口配置
~~~
修改conf/config.yml文件，添加的path会挂到/api/下，按照示例进行修改即可
~~~

## 鸣谢
Mix Go [https://github.com/mix-go/mix](https://github.com/mix-go/mix)