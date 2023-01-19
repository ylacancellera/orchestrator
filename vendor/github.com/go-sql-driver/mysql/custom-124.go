package mysql

import (
	"runtime"
)

func getCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func (mc *mysqlConn) Addr() string {
	return mc.cfg.Addr
}
