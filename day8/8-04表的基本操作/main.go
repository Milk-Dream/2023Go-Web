package main

func main() {
	//查看自己在哪个表
	//select database();
	//查看表的结构
	//describe student或者desc student;

	//查看创建表的语法
	//show create table student;


	/*
		### 8-4 表的基本操作

		- 先选择一个库，然后在这个库中操作表
		- 使用一个库

		~~~mYsql
		use my_db1;
		~~~

		- 查看所在库

		~~~mysql
		select database();
		~~~

		- 新建表

		~~~mysql
		create table t1(id int, name char(6));	# 新建表t1, 有两个字段:id、name
		~~~

		- 查看表

		~~~mysql
		show tables;			# 查看当前库所有表
		describe t1;	 		# 查看指定表，可简写为 desc t1;
		show create table t1;	# 查看指定表
		~~~

		- 修改表

		~~~mysql
		alter table t1 add age int;					# t1表增加age字段
		alter table t1 modify name char(16);		# 修改t1表name字段属性为char(16) 
		alter table t1 drop name;					# 删除字段 drop
		~~~

		- 删除表

		~~~mysql
		drop table t1;			# 清空表
		~~~


	
	*/
}