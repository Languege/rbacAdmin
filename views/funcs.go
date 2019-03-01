package views

import "rbacAdmin/models"

/**
 *@author LanguageY++2013
 *2019/3/1 7:29 PM
 **/
 //检查是否包含该权限
func CheckPermission(p_id int, myPL []models.AdminPermissions) bool {

	for _, v := range myPL {
		if v.Id == p_id {
			return true
		}
	}

	return false
}