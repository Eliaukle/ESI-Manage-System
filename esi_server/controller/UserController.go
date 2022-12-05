package controller

import (
	"blog_server/common"
	"blog_server/model"
	"blog_server/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

// Login 登录
func Login(c *gin.Context) {
	db := common.GetDB()
	// 获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	userName := requestUser.UserName
	password := requestUser.Password
	// 数据验证
	var user model.User
	db.Where("user_name =?", userName).First(&user)
	if user.ID == 0 {
		response.Fail(c, nil, "用户不存在")
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(c, nil, "密码错误")
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Fail(c, nil, "系统错误")
		return
	}
	// 返回结果
	response.Success(c, gin.H{"token": token}, "登录成功")
}

// GetInfo 登录后获取信息
func GetInfo(c *gin.Context) {
	// 获取上下文中的用户信息
	user, _ := c.Get("user")
	// 返回用户信息
	response.Success(c, gin.H{"user": user}, "登录获取信息成功")
}

func CreateUser(c *gin.Context) {
	db := common.GetDB()
	// 获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	userName := requestUser.UserName
	password := requestUser.Password
	identify := requestUser.Identify
	// 数据验证
	var user model.User
	db.Where("user_name =?", userName).First(&user)
	if user.ID != 0 {
		response.Fail(c, nil, "用户已存在")
		return
	}
	// 密码加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 创建用户
	newUser := model.User{
		UserName: userName,
		Password: string(hashedPassword),
		Identify: identify,
	}
	db.Create(&newUser)
	// 返回结果
	response.Success(c, nil, "添加成功")
}

func UpdateUser(c *gin.Context) {
	db := common.GetDB()
	// 获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	identify := requestUser.Identify
	// 获取path中的id
	userId := c.Params.ByName("id")
	// 查找用户信息
	var user model.User
	if db.Where("id = ?", userId).First(&user).RecordNotFound() {
		response.Fail(c, nil, "用户信息不存在")
		return
	}
	newUser := model.User{
		UserName: user.UserName,
		Password: user.Password,
		Identify: identify,
	}
	// 修改用户信息
	if err := db.Model(&user).Update(newUser).Error; err != nil {
		response.Fail(c, nil, "修改失败")
		return
	}
	response.Success(c, nil, "修改成功")
}

func DeleteUser(c *gin.Context) {
	db := common.GetDB()
	// 获取path中的id
	userId := c.Params.ByName("id")
	// 查找用户信息
	var user model.User
	if db.Where("id = ?", userId).First(&user).RecordNotFound() {
		response.Fail(c, nil, "用户信息不存在")
		return
	}
	// 删除用户信息
	if err := db.Delete(&user).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}
	response.Success(c, nil, "删除成功")
}

func SearchUser(c *gin.Context) {
	db := common.GetDB()
	// 获取用户名、身份、分页参数
	userName := c.DefaultQuery("userName", "")
	Identify := c.DefaultQuery("identify", "null")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "7"))
	var query []string
	var args []string
	// 若用户名存在
	if userName != "" {
		query = append(query, "user_name LIKE ?")
		args = append(args, "%"+userName+"%")
	}
	// 若身份存在
	if Identify != "null" {
		query = append(query, "identify = ?")
		args = append(args, Identify)
	}
	// 拼接字符串
	var querystr string
	if len(query) > 0 {
		querystr = strings.Join(query, " AND ")
	}
	var user []model.User
	// 用户总数
	var count int
	// 查询用户信息
	switch len(args) {
	case 0:
		db.Table("users").Select("*").Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&user)
		db.Model(model.User{}).Count(&count)
	case 1:
		db.Table("users").Select("*").Where(querystr, args[0]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&user)
		db.Model(model.User{}).Where(querystr, args[0]).Count(&count)
	case 2:
		db.Table("users").Select("*").Where(querystr, args[0], args[1]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&user)
		db.Model(model.User{}).Where(querystr, args[0], args[1]).Count(&count)
	}
	// 展示用户信息列表
	response.Success(c, gin.H{"user": user, "count": count}, "查找成功")
}
