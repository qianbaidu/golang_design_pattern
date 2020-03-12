package main

//订单
//订单报表
type OrderMainDao interface { //订单记录
	SaveOrderMain()
	//DeleteOrderMain()
	//SearchOrderMain()
}

type OrderDetailDao interface { //订单详情
	SaveOrderDetail()
}

type DaoFactory interface { //抽象工厂接口
	CreateOrderMainDao() OrderMainDao
	CreateOrderDetailMainDao() OrderDetailDao
}
