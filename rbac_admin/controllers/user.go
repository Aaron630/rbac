package controllers

import (
	"encoding/json"
	"fmt"
	"rbac_admin/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) UserInfo() {
	adminUser := &models.AdminUser{}
	adminUser.GetUserInfoById(1) // TODO
	fmt.Println(adminUser)
	jData := `{
    "id": "4291d7da9005377ec9aec4a71ea837f",
    "name": "天野远子",
    "username": "admin",
    "password": "",
    "avatar": "/avatar2.jpg",
    "status": 1,
    "telephone": "",
    "lastLoginIp": "27.154.74.117",
    "lastLoginTime": 1534837621348,
    "creatorId": "admin",
    "createTime": 1497160610259,
    "merchantCode": "TLif2btpzg079h15bk",
    "deleted": 0,
    "roleId": "admin",
    "role": {
      "id": "admin",
      "name": "管理员",
      "describe": "拥有所有权限",
      "status": 1,
      "creatorId": "system",
      "createTime": 1497160610259,
      "deleted": 0,
      "permissions": [{
        "roleId": "admin",
        "permissionId": "dashboard",
        "permissionName": "仪表盘",
        "actionEntitySet": [{
          "action": "add",
          "describe": "新增",
          "defaultCheck": false
        }, {
          "action": "query",
          "describe": "查询",
          "defaultCheck": false
        }, {
          "action": "get",
          "describe": "详情",
          "defaultCheck": false
        }, {
          "action": "update",
          "describe": "修改",
          "defaultCheck": false
        }, {
          "action": "delete",
          "describe": "删除",
          "defaultCheck": false
        }],
        "actionList": null,
        "dataAccess": null
      }, {
        "roleId": "admin",
        "permissionId": "exception",
        "permissionName": "异常页面权限",
        "actionEntitySet": [{
          "action": "add",
          "describe": "新增",
          "defaultCheck": false
        }, {
          "action": "query",
          "describe": "查询",
          "defaultCheck": false
        }, {
          "action": "get",
          "describe": "详情",
          "defaultCheck": false
        }, {
          "action": "update",
          "describe": "修改",
          "defaultCheck": false
        }, {
          "action": "delete",
          "describe": "删除",
          "defaultCheck": false
        }],
        "actionList": null,
        "dataAccess": null
      }, {
        "roleId": "admin",
        "permissionId": "result",
        "permissionName": "结果权限",
        "actionEntitySet": [{
          "action": "add",
          "describe": "新增",
          "defaultCheck": false
        }, {
          "action": "query",
          "describe": "查询",
          "defaultCheck": false
        }, {
          "action": "get",
          "describe": "详情",
          "defaultCheck": false
        }, {
          "action": "update",
          "describe": "修改",
          "defaultCheck": false
        }, {
          "action": "delete",
          "describe": "删除",
          "defaultCheck": false
        }],
        "actionList": null,
        "dataAccess": null
      }, {
        "roleId": "admin",
        "permissionId": "profile",
        "permissionName": "详细页权限",
        "actionEntitySet": [{
          "action": "add",
          "describe": "新增",
          "defaultCheck": false
        }, {
          "action": "query",
          "describe": "查询",
          "defaultCheck": false
        }, {
          "action": "get",
          "describe": "详情",
          "defaultCheck": false
        }, {
          "action": "update",
          "describe": "修改",
          "defaultCheck": false
        }, {
          "action": "delete",
          "describe": "删除",
          "defaultCheck": false
        }],
        "actionList": null,
        "dataAccess": null
      }, {
        "roleId": "admin",
        "permissionId": "table",
        "permissionName": "表格权限",
        "actionEntitySet": [{
          "action": "add",
          "describe": "新增",
          "defaultCheck": false
        }, {
          "action": "import",
          "describe": "导入",
          "defaultCheck": false
        }, {
          "action": "get",
          "describe": "详情",
          "defaultCheck": false
        }, {
          "action": "update",
          "describe": "修改",
          "defaultCheck": false
        }],
        "actionList": null,
        "dataAccess": null
      }, {
        "roleId": "admin",
        "permissionId": "form",
        "permissionName": "表单权限",
        "actionEntitySet": [{
          "action": "add",
          "describe": "新增",
          "defaultCheck": false
        }, {
          "action": "get",
          "describe": "详情",
          "defaultCheck": false
        }, {
          "action": "query",
          "describe": "查询",
          "defaultCheck": false
        }, {
          "action": "update",
          "describe": "修改",
          "defaultCheck": false
        }, {
          "action": "delete",
          "describe": "删除",
          "defaultCheck": false
        }],
        "actionList": null,
        "dataAccess": null
      }, {
        "roleId": "admin",
        "permissionId": "order",
        "permissionName": "订单管理",
        "actionEntitySet": [{
          "action": "add",
          "describe": "新增",
          "defaultCheck": false
        }, {
          "action": "query",
          "describe": "查询",
          "defaultCheck": false
        }, {
          "action": "get",
          "describe": "详情",
          "defaultCheck": false
        }, {
          "action": "update",
          "describe": "修改",
          "defaultCheck": false
        }, {
          "action": "delete",
          "describe": "删除",
          "defaultCheck": false
        }],
        "actionList": null,
        "dataAccess": null
      }, {
        "roleId": "admin",
        "permissionId": "permission",
        "permissionName": "权限管理",
        "actionEntitySet": [{
          "action": "add",
          "describe": "新增",
          "defaultCheck": false
        }, {
          "action": "get",
          "describe": "详情",
          "defaultCheck": false
        }, {
          "action": "update",
          "describe": "修改",
          "defaultCheck": false
        }, {
          "action": "delete",
          "describe": "删除",
          "defaultCheck": false
        }],
        "actionList": null,
        "dataAccess": null
      }, {
        "roleId": "admin",
        "permissionId": "role",
        "permissionName": "角色管理",
        "actionEntitySet": [{
          "action": "add",
          "describe": "新增",
          "defaultCheck": false
        }, {
          "action": "get",
          "describe": "详情",
          "defaultCheck": false
        }, {
          "action": "update",
          "describe": "修改",
          "defaultCheck": false
        }, {
          "action": "delete",
          "describe": "删除",
          "defaultCheck": false
        }],
        "actionList": null,
        "dataAccess": null
      }, {
        "roleId": "admin",
        "permissionId": "table",
        "permissionName": "桌子管理",
        "actionEntitySet": [{
          "action": "add",
          "describe": "新增",
          "defaultCheck": false
        }, {
          "action": "get",
          "describe": "详情",
          "defaultCheck": false
        }, {
          "action": "query",
          "describe": "查询",
          "defaultCheck": false
        }, {
          "action": "update",
          "describe": "修改",
          "defaultCheck": false
        }, {
          "action": "delete",
          "describe": "删除",
          "defaultCheck": false
        }],
        "actionList": null,
        "dataAccess": null
      }, {
        "roleId": "admin",
        "permissionId": "user",
        "permissionName": "用户管理",
        "actionEntitySet": [{
          "action": "add",
          "describe": "新增",
          "defaultCheck": false
        }, {
          "action": "import",
          "describe": "导入",
          "defaultCheck": false
        }, {
          "action": "get",
          "describe": "详情",
          "defaultCheck": false
        }, {
          "action": "update",
          "describe": "修改",
          "defaultCheck": false
        }, {
          "action": "delete",
          "describe": "删除",
          "defaultCheck": false
        }, {
          "action": "export",
          "describe": "导出",
          "defaultCheck": false
        }],
        "actionList": null,
        "dataAccess": null
      },
      {
        "roleId": "admin",
        "permissionId": "support",
        "permissionName": "超级模块",
        "actionEntitySet": [{
          "action": "add",
          "describe": "新增",
          "defaultCheck": false
        }, {
          "action": "import",
          "describe": "导入",
          "defaultCheck": false
        }, {
          "action": "get",
          "describe": "详情",
          "defaultCheck": false
        }, {
          "action": "update",
          "describe": "修改",
          "defaultCheck": false
        }, {
          "action": "delete",
          "describe": "删除",
          "defaultCheck": false
        }, {
          "action": "export",
          "describe": "导出",
          "defaultCheck": false
        }],
        "actionList": null,
        "dataAccess": null
      }]
    }
  }`
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jData), &data)
	if err != nil {
		fmt.Println(err)
	}
	c.Success(data, "获取用户信息成功")
}
