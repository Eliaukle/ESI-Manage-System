package controller

import (
	"blog_server/common"
	"blog_server/model"
	"blog_server/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

type SubjectController struct {
	DB *gorm.DB
}

type ISubjectController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Search(c *gin.Context)
	List(c *gin.Context)
}

func (s SubjectController) Create(c *gin.Context) {
	var subjectRequest model.Subject
	// 数据验证
	if err := c.ShouldBindJSON(&subjectRequest); err != nil {
		response.Fail(c, nil, "数据错误")
		return
	}
	// 添加学科信息
	subject := model.Subject{
		SubName: subjectRequest.SubName,
		SubType: subjectRequest.SubType,
	}
	if err := s.DB.Create(&subject).Error; err != nil {
		response.Fail(c, nil, "学科名称已存在")
		return
	}
	response.Success(c, gin.H{"id": subject.ID}, "添加成功")
}

func (s SubjectController) Update(c *gin.Context) {
	var subjectRequest model.Subject
	// 数据验证
	if err := c.ShouldBindJSON(&subjectRequest); err != nil {
		response.Fail(c, nil, "数据错误")
		return
	}
	// 获取path中的id
	subjectId := c.Params.ByName("id")
	// 查找学科信息
	var subject model.Subject
	if s.DB.Where("id = ?", subjectId).First(&subject).RecordNotFound() {
		response.Fail(c, nil, "学科不存在")
		return
	}
	// 修改学科信息
	if err := s.DB.Model(&subject).Update(subjectRequest).Error; err != nil {
		response.Fail(c, nil, "修改失败")
		return
	}
	response.Success(c, nil, "修改成功")
}

func (s SubjectController) Delete(c *gin.Context) {
	// 获取path中的id
	subjectId := c.Params.ByName("id")
	// 查找学科信息
	var subject model.Subject
	if s.DB.Where("id = ?", subjectId).First(&subject).RecordNotFound() {
		response.Fail(c, nil, "大学不存在")
		return
	}
	// 删除学科信息
	if err := s.DB.Delete(&subject).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}
	response.Success(c, nil, "删除成功")
}

func (s SubjectController) Search(c *gin.Context) {
	// 获取学科类别、学科名称、分页参数
	subType := c.DefaultQuery("subType", "null")
	subName := c.DefaultQuery("subName", "")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "7"))
	var query []string
	var args []string
	// 若学科类别存在
	if subType != "null" {
		query = append(query, "sub_type = ?")
		args = append(args, subType)
	}
	// 若学科名称存在
	if subName != "" {
		query = append(query, "sub_name LIKE ?")
		args = append(args, "%"+subName+"%")
	}
	// 拼接字符串
	var querystr string
	if len(query) > 0 {
		querystr = strings.Join(query, " AND ")
	}
	var subject []model.Subject
	// 学科总数
	var count int
	// 查询学科
	switch len(args) {
	case 0:
		s.DB.Table("subjects").Select("*").Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&subject)
		s.DB.Model(model.Subject{}).Count(&count)
	case 1:
		s.DB.Table("subjects").Select("*").Where(querystr, args[0]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&subject)
		s.DB.Model(model.Subject{}).Where(querystr, args[0]).Count(&count)
	case 2:
		s.DB.Table("subjects").Select("*").Where(querystr, args[0], args[1]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&subject)
		s.DB.Model(model.Subject{}).Where(querystr, args[0], args[1]).Count(&count)
	case 3:
		s.DB.Table("subjects").Select("*").Where(querystr, args[0], args[1], args[2]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&subject)
		s.DB.Model(model.Subject{}).Where(querystr, args[0], args[1], args[2]).Count(&count)
	}
	// 展示学科信息列表
	response.Success(c, gin.H{"subject": subject, "count": count}, "查找成功")
}

func (s SubjectController) List(c *gin.Context) {
	// 学科信息
	var subject []model.Subject
	s.DB.Table("subjects").Select("id, sub_name").Find(&subject)
	// 展示学科信息
	response.Success(c, gin.H{"subject": subject}, "查找成功")
}

func NewSubjectController() ISubjectController {
	db := common.GetDB()
	db.AutoMigrate(model.Subject{})
	return SubjectController{DB: db}
}
