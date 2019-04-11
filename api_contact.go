package godingtalk

import (
	"fmt"
	"net/url"
)

type User struct {
	OAPIResponse
	Userid     string `json:"userid"`
	Name       string `json:"name"`
	Mobile     string `json:"mobile"`
	Tel        string `json:"tel"`
	Remark     string `json:"remark"`
	Order      int    `json:"order"`
	IsAdmin    bool
	IsBoss     bool
	IsLeader   bool
	IsSys      bool `json:"is_sys"`
	SysLevel   int  `json:"sys_level"`
	Active     bool
	Department []int
	Position   string
	HiredDate  string `json:"hiredDate"`
	Jobnumber  string `json:"jobnumber"`
	Email      string `json:"email"`
	Avatar     string
	Extattr    interface{}
}

type UserList struct {
	OAPIResponse
	HasMore  bool   `json:"hasMore"`
	Userlist []User `json:"userlist"`
}

type Department struct {
	OAPIResponse
	Id                    int    `json:"id"`
	Name                  string `json:"name"`
	ParentId              int    `json:"parentid"`
	Order                 int
	DeptPerimits          string
	UserPerimits          string
	OuterDept             bool
	OuterPermitDepts      string
	OuterPermitUsers      string
	OrgDeptOwner          string
	DeptManagerUseridList string
}

type DepartmentList struct {
	OAPIResponse
	Departments []Department `json:"department"`
}

// DepartmentList is 获取部门列表
func (c *DingTalkClient) DepartmentList() (DepartmentList, error) {
	var data DepartmentList
	err := c.httpRPC("department/list", nil, nil, &data)
	return data, err
}

//DepartmentDetail is 获取部门详情
func (c *DingTalkClient) DepartmentDetail(id int) (Department, error) {
	var data Department
	params := url.Values{}
	params.Add("id", fmt.Sprintf("%d", id))
	err := c.httpRPC("department/get", params, nil, &data)
	return data, err
}

//UserList is 获取部门成员
func (c *DingTalkClient) UserList(departmentID, offset, size int) (UserList, error) {
	var data UserList
	if size > 100 {
		return data, fmt.Errorf("size 最大100")
	}

	params := url.Values{}
	params.Add("department_id", fmt.Sprintf("%d", departmentID))
	params.Add("offset", fmt.Sprintf("%d", offset))
	params.Add("size", fmt.Sprintf("%d", size))
	err := c.httpRPC("user/list", params, nil, &data)
	return data, err
}

//CreateChat is
func (c *DingTalkClient) CreateChat(name string, owner string, useridlist []string) (string, error) {
	var data struct {
		OAPIResponse
		Chatid string
	}
	request := map[string]interface{}{
		"name":       name,
		"owner":      owner,
		"useridlist": useridlist,
	}
	err := c.httpRPC("chat/create", nil, request, &data)
	return data.Chatid, err
}

//UserInfoByCode 校验免登录码并换取用户身份
func (c *DingTalkClient) UserInfoByCode(code string) (User, error) {
	var data User
	params := url.Values{}
	params.Add("code", code)
	err := c.httpRPC("user/getuserinfo", params, nil, &data)
	return data, err
}

//UseridByUnionId 通过UnionId获取玩家Userid
func (c *DingTalkClient) UseridByUnionId(unionid string) (string, error) {
	var data struct {
		OAPIResponse
		UserID string `json:"userid"`
	}

	params := url.Values{}
	params.Add("unionid", unionid)
	err := c.httpRPC("user/getUseridByUnionid", params, nil, &data)
	if err != nil {
		return "", err
	}

	return data.UserID, err
}

//UserDetail is 获取用户详情
func (c *DingTalkClient) UserDetail(id string) (User, error) {
	var data User
	params := url.Values{}
	params.Add("userid", id)
	err := c.httpRPC("user/get", params, nil, &data)
	return data, err
}

type DepartmentChild struct {
	OAPIResponse
	SubDeptIdList []int `json:"sub_dept_id_list"`
}

//DepartmentChildIds is 获取子部门ID列表
func (c *DingTalkClient) DepartmentChildIds(departmentID int) (DepartmentChild, error) {
	var data DepartmentChild
	params := url.Values{}
	params.Add("id", fmt.Sprintf("%d", departmentID))
	err := c.httpRPC("department/list_ids", params, nil, &data)
	return data, err
	return data.UserID, err
}
