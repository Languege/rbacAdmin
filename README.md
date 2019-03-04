# 概述
rbacAdmin是使用beego开发的后台权限管理模板，实现RBAC权限管理方式，前端采用layui2.x。

# 安装
下载该项目后
```bash
git pull https://github.com/Languege/rbacAdmin
```

创建项目数据库，然后执行rbac.sql初始化权限管理表


默认登录账户```admin@admin.com``` 密码```12345678```

修改app.conf数据库连接信息


# 运行
进入项目根目录，运行
```$xslt
bee run
```

打开 http://127.0.0.1:8080

# 使用

为你的接口建立权限-->将权限添加到角色-->将角色绑定到用户

