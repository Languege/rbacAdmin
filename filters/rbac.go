package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/casbin/beego-orm-adapter"
	"github.com/casbin/casbin"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/authz"
)

/**
 *@author LanguageY++2013
 *2019/2/27 7:05 PM
 **/
func init() {
	beego.InsertFilter("/admin/*", beego.BeforeRouter, func(context *context.Context) {
		//校验用户名和密码
		sess, _ := beego.GlobalSessions.SessionStart(context.ResponseWriter, context.Request)
		defer sess.SessionRelease(context.ResponseWriter)

		role := ""

		sessRole := sess.Get("role")
		if sessRole != nil {
			role = sessRole.(string)
		}

		//角色
		context.Request.SetBasicAuth(role, "")
	})

	//beego-orm-adapter
	a := beegoormadapter.NewAdapter("mysql", beego.AppConfig.String("mysqldsn"), true)
	e := casbin.NewEnforcer("rbac_model.conf", a)

	e.LoadPolicy()

	logs.Info("Policies:", e.GetPolicy())

	beego.InsertFilter("/admin/*", beego.BeforeRouter, authz.NewAuthorizer(e))
}