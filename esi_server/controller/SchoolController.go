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

type SchoolController struct {
	DB *gorm.DB
}

type ISchoolController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Search(c *gin.Context)
	List(c *gin.Context)
}

func (s SchoolController) Create(c *gin.Context) {
	var schoolRequest model.School
	// 数据验证
	if err := c.ShouldBindJSON(&schoolRequest); err != nil {
		response.Fail(c, nil, "数据错误")
		return
	}
	// 添加大学信息
	school := model.School{
		ID:        schoolRequest.ID,
		SchName:   schoolRequest.SchName,
		SchType:   schoolRequest.SchType,
		SchManage: schoolRequest.SchManage,
	}
	if err := s.DB.Create(&school).Error; err != nil {
		response.Fail(c, nil, "大学编号或大学名称已存在")
		return
	}
	response.Success(c, gin.H{"id": school.ID}, "添加成功")
}

func (s SchoolController) Update(c *gin.Context) {
	var schoolRequest model.School
	// 数据验证
	if err := c.ShouldBindJSON(&schoolRequest); err != nil {
		response.Fail(c, nil, "数据错误")
		return
	}
	// 获取path中的id
	schoolId := c.Params.ByName("id")
	// 查找大学信息
	var school model.School
	if s.DB.Where("id = ?", schoolId).First(&school).RecordNotFound() {
		response.Fail(c, nil, "大学信息不存在")
		return
	}
	// 修改大学信息
	if err := s.DB.Model(&school).Update(schoolRequest).Error; err != nil {
		response.Fail(c, nil, "修改失败")
		return
	}
	response.Success(c, nil, "修改成功")
}

func (s SchoolController) Delete(c *gin.Context) {
	// 获取path中的id
	schoolId := c.Params.ByName("id")
	// 查找大学信息
	var school model.School
	if s.DB.Where("id = ?", schoolId).First(&school).RecordNotFound() {
		response.Fail(c, nil, "大学信息不存在")
		return
	}
	// 删除大学信息
	if err := s.DB.Delete(&school).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}
	response.Success(c, nil, "删除成功")
}

func (s SchoolController) Search(c *gin.Context) {
	// 获取大学名称、类别、性质、分页参数
	schName := c.DefaultQuery("schName", "")
	schType := c.DefaultQuery("schType", "null")
	schManage := c.DefaultQuery("schManage", "null")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "7"))
	var query []string
	var args []string
	// 若大学名称存在
	if schName != "" {
		query = append(query, "sch_name LIKE ?")
		args = append(args, "%"+schName+"%")
	}
	// 若类别存在
	if schType != "null" {
		query = append(query, "sch_type = ?")
		args = append(args, schType)
	}
	// 若性质存在
	if schManage != "null" {
		query = append(query, "sch_manage =?")
		args = append(args, schManage)
	}
	// 拼接字符串
	var querystr string
	if len(query) > 0 {
		querystr = strings.Join(query, " AND ")
	}
	var school []model.School
	// 大学总数
	var count int
	// 查询大学信息
	switch len(args) {
	case 0:
		s.DB.Table("schools").Select("*").Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&school)
		s.DB.Model(model.School{}).Count(&count)
	case 1:
		s.DB.Table("schools").Select("*").Where(querystr, args[0]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&school)
		s.DB.Model(model.School{}).Where(querystr, args[0]).Count(&count)
	case 2:
		s.DB.Table("schools").Select("*").Where(querystr, args[0], args[1]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&school)
		s.DB.Model(model.School{}).Where(querystr, args[0], args[1]).Count(&count)
	case 3:
		s.DB.Table("schools").Select("*").Where(querystr, args[0], args[1], args[2]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&school)
		s.DB.Model(model.School{}).Where(querystr, args[0], args[1], args[2]).Count(&count)
	}
	// 展示大学信息列表
	response.Success(c, gin.H{"school": school, "count": count}, "查找成功")
}

func (s SchoolController) List(c *gin.Context) {
	// 大学信息
	var school []model.School
	s.DB.Table("schools").Select("id, sch_name").Find(&school)
	// 展示大学信息
	response.Success(c, gin.H{"school": school}, "查找成功")
}

func NewSchoolController() ISchoolController {
	db := common.GetDB()
	db.AutoMigrate(model.School{})
	return SchoolController{DB: db}
}
