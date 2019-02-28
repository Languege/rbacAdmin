package repositories

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"rbacAdmin/models"
)

/**
 *@author LanguageY++2013
 *2019/2/28 4:38 PM
 **/
func GetAllAdminUsersCount(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.AdminUsers))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}


	count, err =  qs.Count()
	return
}

func AdminUsers_Pagination(query map[string]string, fields []string, sortby []string, order []string,
	page int64, pageSize int64) (pagination Page, err error){

	offset := (page - 1) * pageSize
	limit := pageSize

	l, err := models.GetAllAdminUsers(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}
	count, err := GetAllAdminUsersCount(query)

	pagination = PageUtil(count, page, pageSize, l)

	return
}