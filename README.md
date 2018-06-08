# confd_demo
 -  conf.d文件夹中是多个confd的配置文件。confd启动的时候会读取所有的toml文件，根据toml中所声明的key向etcd注册，并且记录在内存中。toml文件之后可以用脚本自动生成。
 - 当收到来自etcd的变化通知，confd会根据变化的key运行相应的toml文件指定的操作。 读取templates文件夹中相应的tmpl文件作为模板，在demo_dir生成对应的php缓存文件。