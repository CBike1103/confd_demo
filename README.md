# confd_demo
### conf.d文件夹中是多个confd的配置文件。confd启动的时候会读取所有的toml文件，根据toml中所声明的key向etcd注册，并且记录在内存中。toml文件之后可以用脚本自动生成。
### 当收到来自etcd的变化通知，confd会用一个 sync.watchGroup 运行每一个toml文件指定的操作。操作包括
 - 创建一个log file，在其中记录变化的kv
 - 之后运行一个 命令行命令
### demo中的命令行命令是运行一个python脚本。这个脚本会读取之前创建的log来获取变化的kv，并写入一个文件中。

# 待改进的地方
### etcd的功能过于单一化，与我们的需求并不完全统一。并且还有以下缺点
 - 收到etcd通知的时候，不会记录下变化的kv。需要另外向etcd发起请求获得kv.
 - 每次都一定要生成一个文件，无法跳过
 - 同时集成了zookeeper，redis等等很多client, 有些多余

### 虽然也可以去改动confd的代码，但是考虑到confd多余的代码很多，最理想的状况是把研究重点放在etcd的client，自己写一个client，并且在其中实现每次获得watch反馈时运行脚本。