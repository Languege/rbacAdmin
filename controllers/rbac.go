package controllers

import (
	"rbacAdmin/repositories"
)

// RbacController operations for Rbac
type RbacController struct {
	BaseController
}

// URLMapping ...
func (c *RbacController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Rbac
// @Param	body		body 	models.Rbac	true		"body for Rbac content"
// @Success 201 {object} models.Rbac
// @Failure 403 body is empty
// @router / [post]
func (c *RbacController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Rbac by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Rbac
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RbacController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Rbac
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Rbac
// @Failure 403
// @router / [get]
func (c *RbacController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Rbac
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Rbac	true		"body for Rbac content"
// @Success 200 {object} models.Rbac
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RbacController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Rbac
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RbacController) Delete() {

}

//@router /permission_list	[get]
func (c *RbacController) PermissionList() {

	query, fields, sortby, order, page, pageSize, err := c.PageParams()
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	pageData, err := repositories.AdminPermissions_Pagination(query, fields, sortby, order, page, pageSize)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	} else {

		c.TplName = "rbac/permission_list.html"
		c.Layout = "_layout/iframe_layout.html"
		c.Data["page"] = pageData
	}
}

//@router /role_list	[get]
func (c *RbacController) RoleList()  {
	query, fields, sortby, order, page, pageSize, err := c.PageParams()

	pageData, err := repositories.AdminRoles_Pagination(query, fields, sortby, order, page, pageSize)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	} else {

		c.TplName = "rbac/role_list.html"
		c.Layout = "_layout/iframe_layout.html"
		c.Data["page"] = pageData
	}
}
