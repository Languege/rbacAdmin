package repositories

import (
	"rbacAdmin/models"
	"github.com/astaxie/beego/orm"
	"strings"
)

/**
 *@author LanguageY++2013
 *2019/2/28 7:29 PM
 **/
func AdminRoles_Pagination(query map[string]string, fields []string, sortby []string, order []string,
	page int64, pageSize int64) (pagination Page, err error){

	offset := (page - 1) * pageSize
	limit := pageSize

	l, err := models.GetAllAdminRoles(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}
	count, err := AdminRoles_GetCount(query)

	pagination = PageUtil(count, page, pageSize, l)

	return
}

func AdminRoles_GetCount(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.AdminRoles))
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