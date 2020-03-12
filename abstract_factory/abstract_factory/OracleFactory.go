package main

type OracleFactory struct {

}

func (*OracleFactory)CreateOrderMainDao() OrderMainDao {
	return &MysqlMainDao{}
}


func (*OracleFactory)CreateOrderDetailMainDao() OrderDetailDao {
	return &MysqlDetailDao{}
}


