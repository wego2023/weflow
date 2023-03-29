// Code generated by hertz generator. DO NOT EDIT.

package Hello

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	hello "github.com/wego2023/weflow/internal/biz/handler/hello"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/hello", append(_hellomethodMw(), hello.HelloMethod)...)
}
