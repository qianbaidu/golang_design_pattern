package main

import "fmt"


// Oracl 针对接口的实现
type OracleMainDao struct {

}

func (*OracleMainDao) SaveOrderMain()  {
	fmt.Println("Oracl main save")
}

// mysql detail接口实现
type OraclDetailDao struct {

}

func (*OraclDetailDao) SaveOrderDetail()  {
	fmt.Println("Oracl detail save")
}


