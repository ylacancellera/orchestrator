package sqlutils

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func getCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func getUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

func wrapErrBadConn(db *sql.DB, err error) error {
	if err == nil {
		return nil
	}
	if !errors.Is(err, driver.ErrBadConn) {
		return err
	}
	conn := reflect.ValueOf(db).Elem().FieldByName("connector")
	cfg := reflect.NewAt(conn.Type(), unsafe.Pointer(conn.UnsafeAddr())).Elem().Interface()

	config := reflect.ValueOf(cfg)
	addr := config.Elem().FieldByName("cfg").Elem().FieldByName("Addr")

	return fmt.Errorf("func: %s, addr: %s, err: %w", getCurrentFuncName(), addr, err)
}
