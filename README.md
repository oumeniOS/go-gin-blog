# go-gin-blog


##### 使用docker启动MySQL容器
1、拉取MySQL镜像  
  
  	docker pull mysql  
待镜像安装完成后，通过下面的命令查看：  
  
  	docker images   //查看当前的所有镜像  
![](https://github.com/oumeniOS/go-gin-blog/blob/master/img-folder/001.jpg?raw=true '查看所有镜像')
接着我们启动一个mysql镜像的容器，输入下面的命令：  
  		
  	docker run -itd -P mysql bash 
  	 
* docker run : 启动一个容器  
* -itd：i是交互式操作，t是一个终端，d指的是在后台运行  
* -P 指在本地生成一个随机端口，用来映射mysql的3306端口  
* mysql：刚才下的mysql镜像名  
* bash：指创建一个交互式shell   

然后我们查看已经运行的镜像：    

	docker ps -a  

![](https://github.com/oumeniOS/go-gin-blog/blob/master/img-folder/002.jpg?raw=true '查看所有容器')
接下来进入容器：

	docker exec -it mysql bash 

* docker exec: docker进入某个容器的命令
* mysql 是容器的NAME
这样我们就进入mysql的容器了  

想要退出容器，执行命令：  

	exit;  

参考链接: [https://www.jianshu.com/p/83ecd99cf3eb] (https://www.jianshu.com/p/83ecd99cf3eb)

#####docker的基本命令
1.查看正在运行的容器

	docker ps

2.查看所有的容器（包括正在运行的和停止运行的容器）

	docker ps -a

3.停止运行一个容器
	
	docker stop <container_id>
	
4.删除一个容器(只能删除已经停止运行的容器)

	docker rm <container_id>

5.查看所有的镜像

	docker images
	
6.删除指定ID的镜像（删除前需要停止对应的容器）

	docker rmi <image_id>
	docker rmi -f <image_id>//强制删除
7.进入指定ID的容器

	docker exec -it <container_id> bash

 
* docker exec: docker进入某个容器的命令
* bash 可选
* -it 可选  

#####Mac下MySQL命令行操作
1.连接到本机的mysql

	mysql -uroot -ppassword

* root为本机mysql的用户名
* password为本机mysql root用户的密码

2.退出mysql命令
	
	exit;
###### 数据库的相关操作
1. 创建数据库  

		create database <database_name>;

2. 显示所有的数据库

		show databases;

3. 切换数据库

		use <database_name>;
	
4. 删除数据库
	
		drop database <database_name>;

#######数据表操作
1.创建数据表
	
	use <database_name>;
	create table <table_name> <(字段设定列表)>;
	
2.显示所有的数据表

	use <database_name>;
	show tables;

3.某表的数据结构
	
	describe <table_name>;

4.显示表中所有的记录

	select * from <table_name>;

5.删除表

	drop table <table_name>;

6.删除表数据
	
	delete from <table_name>;
	truncate table <table_name>;

(不带where参数的delete语句可以删除mysql表中所有内容；)

* 使用truncate table也可以清空mysql表中所有内容；
* 但是使用delete清空表中的记录，内容的ID仍然从删除点的ID继续建立，而不是从1开始。
* 而truncate相当于保留了表的结构而重新建立了一张同样的新表。
* 效率上truncate比delete快。但truncate删除后不记录mysql日志，不可以恢复数据。
* delete的效果有点像将mysql表中所有记录一条一条删除到删完。

参考链接：[https://blog.csdn.net/qq_19484963/article/details/80431703](https://blog.csdn.net/qq_19484963/article/details/80431703)

外部访问docker中的服务
docker-machine 192.168.99.100 go-gin-blog  
* 192.168.99.100 外部访问的ip
* go-gin-blog 容器名字
* docker-machine 命令
