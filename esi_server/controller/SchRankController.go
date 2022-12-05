package controller

import (
	"blog_server/common"
	"blog_server/model"
	"blog_server/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

type SchRankController struct {
	DB *gorm.DB
}

type ISchRankController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Search(c *gin.Context)
}

func (sr SchRankController) Create(c *gin.Context) {
	var schRankRequest model.SchRank
	// 数据验证
	if err := c.ShouldBindJSON(&schRankRequest); err != nil {
		response.Fail(c, nil, "数据错误")
		return
	}
	// 判断大学排名是否重复添加
	var existSchRank []model.SchRank
	sr.DB.Where("(year = ? AND sch_name = ?) OR (year = ? AND sch_count_rank = ?) OR (year=? AND sch_world_rank = ?)", schRankRequest.Year, schRankRequest.SchName, schRankRequest.Year, schRankRequest.SchCountRank, schRankRequest.Year, schRankRequest.SchWorldRank).Find(&existSchRank)
	if len(existSchRank) > 0 {
		response.Fail(c, nil, "大学排名已存在")
		return
	}
	// 添加大学排名
	schRank := model.SchRank{
		Year:         schRankRequest.Year,
		SchName:      schRankRequest.SchName,
		SchCountRank: schRankRequest.SchCountRank,
		SchWorldRank: schRankRequest.SchWorldRank,
	}
	// 添加大学排名
	if err := sr.DB.Create(&schRank).Error; err != nil {
		response.Fail(c, nil, "添加失败")
		return
	}
	response.Success(c, gin.H{"id": schRank.ID}, "添加成功")
}

func (sr SchRankController) Update(c *gin.Context) {
	var schRankRequest model.SchRank
	// 数据验证
	if err := c.ShouldBindJSON(&schRankRequest); err != nil {
		response.Fail(c, nil, "数据错误")
		return
	}
	// 获取path中的id
	schRankId := c.Params.ByName("id")
	// 查找大学排名
	var schRank model.SchRank
	if sr.DB.Where("id = ?", schRankId).First(&schRank).RecordNotFound() {
		response.Fail(c, nil, "大学排名不存在")
		return
	}
	// 判断大学排名是否重复添加
	var existSchRank []model.SchRank
	sr.DB.Where("(year = ? AND sch_count_rank = ?) OR (year=? AND sch_world_rank = ?)", schRank.Year, schRankRequest.SchCountRank, schRank.Year, schRankRequest.SchWorldRank).Find(&existSchRank)
	fmt.Println(existSchRank)
	if len(existSchRank) > 1 {
		response.Fail(c, nil, "大学排名已存在")
		return
	}
	// 修改大学排名信息
	if err := sr.DB.Model(&schRank).Update(schRankRequest).Error; err != nil {
		response.Fail(c, nil, "修改失败")
		return
	}
	response.Success(c, nil, "修改成功")
}

func (sr SchRankController) Delete(c *gin.Context) {
	// 获取path中的id
	schRankId := c.Params.ByName("id")
	// 查找大学排名信息
	var schRank model.SchRank
	if sr.DB.Where("id = ?", schRankId).First(&schRank).RecordNotFound() {
		response.Fail(c, nil, "大学排名不存在")
		return
	}
	// 删除大学排名信息
	if err := sr.DB.Delete(&schRank).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}
	response.Success(c, nil, "删除成功")
}

func (sr SchRankController) Search(c *gin.Context) {
	// 获取年份、大学、分页参数
	year := c.DefaultQuery("year", "null")
	schName := c.DefaultQuery("schName", "")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "8"))
	var query []string
	var args []string
	// 若年份存在
	if year != "null" {
		query = append(query, "year = ?")
		args = append(args, year)
	}
	// 若大学名称存在
	if schName != "" {
		query = append(query, "sch_name LIKE ?")
		args = append(args, "%"+schName+"%")
	}
	// 拼接字符串
	var querystr string
	if len(query) > 0 {
		querystr = strings.Join(query, " AND ")
	}
	var schRank []model.SchRank
	// 大学排名总数
	var count int
	// 查询大学排名
	switch len(args) {
	case 0:
		sr.DB.Table("sch_ranks").Select("*").Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&schRank)
		sr.DB.Model(model.SchRank{}).Count(&count)
	case 1:
		sr.DB.Table("sch_ranks").Select("*").Where(querystr, args[0]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&schRank)
		sr.DB.Model(model.SchRank{}).Where(querystr, args[0]).Count(&count)
	case 2:
		sr.DB.Table("sch_ranks").Select("*").Where(querystr, args[0], args[1]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&schRank)
		sr.DB.Model(model.SchRank{}).Where(querystr, args[0], args[1]).Count(&count)
	}
	// 展示大学排名列表
	response.Success(c, gin.H{"schRank": schRank, "count": count}, "查找成功")
}

func NewSchRankController() ISchRankController {
	db := common.GetDB()
	db.AutoMigrate(model.SchRank{})
	return SchRankController{DB: db}
}
