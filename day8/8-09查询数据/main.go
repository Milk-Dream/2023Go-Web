package main


func main() {
	//查询表操作语法
	//SELECT DISTINCT field_list
	//FROM table
	//WHERE <where_condition>
	//GROUP BY<group_by_list>
	//HAVING <having_condition>
	//ORDER BY <order_by_condition>
	//LIMIT <limit_number>

	//and not or 逻辑判断
	//范围 between 1 and 10 1~10之间
	//like 模糊查询 like "%小白%"包含小白
	//select distinct age from student//去除重复查询
	//select * from student order by score 默认升序，
	//select * from student order by score asc 升序
	//select * from student order by score desc 降序

	//select * from student limit 3;只展示3条数据
	//select * from student limit 0,5; 1,2,3,4,5
	//select * from student limit 5,5; 6,7,8,9,10

	//select post as "部门", max(salary) as "最高薪资" from 
	//as 取别名
}