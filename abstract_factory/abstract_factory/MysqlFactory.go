package main

type MysqlFactory struct {

}

func (*MysqlFactory)CreateOrderMainDao() OrderMainDao {
	return &MysqlMainDao{}
}


func (*MysqlFactory)CreateOrderDetailMainDao() OrderDetailDao {
	return &MysqlDetailDao{}
}


