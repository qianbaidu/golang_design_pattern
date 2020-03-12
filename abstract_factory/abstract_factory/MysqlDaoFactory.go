package main

import "fmt"


// mysql针对接口的实现
type MysqlMainDao struct {

}

func (*MysqlMainDao) SaveOrderMain()  {
	fmt.Println("mysql main save")
}

// mysql detail接口实现
type MysqlDetailDao struct {

}

func (*MysqlDetailDao) SaveOrderDetail()  {
	fmt.Println("mysql detail save")
}


