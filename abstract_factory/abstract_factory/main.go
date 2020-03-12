package main

func main()  {
	f := &MysqlFactory{}
	f.CreateOrderMainDao().SaveOrderMain()
	f.CreateOrderDetailMainDao().SaveOrderDetail()


	f1 := &OracleFactory{}
	f1.CreateOrderMainDao().SaveOrderMain()
	f1.CreateOrderDetailMainDao().SaveOrderDetail()
}
