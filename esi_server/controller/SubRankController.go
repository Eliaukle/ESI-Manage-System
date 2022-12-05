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

type SubRankController struct {
	DB *gorm.DB
}

type ISubRankController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Search(c *gin.Context)
}

func (sr SubRankController) Create(c *gin.Context) {
	var subRankRequest model.SubRank
	// 数据验证
	if err := c.ShouldBindJSON(&subRankRequest); err != nil {
		response.Fail(c, nil, "数据错误")
		return
	}
	// 判断学科排名是否重复添加
	var existSubRank []model.SubRank
	sr.DB.Where("(year = ? AND sch_name = ? AND sub_name = ?) OR (year = ? AND sub_Count_Rank = ?) OR (year = ? AND sub_World_Rank = ?)", subRankRequest.Year, subRankRequest.SchName, subRankRequest.SubName, subRankRequest.Year, subRankRequest.SubCountRank, subRankRequest.Year, subRankRequest.SubWorldRank).Find(&existSubRank)
	if len(existSubRank) > 0 {
		response.Fail(c, nil, "学科排名已存在")
		return
	}
	// 添加学科排名
	schRank := model.SubRank{
		Year:         subRankRequest.Year,
		SchName:      subRankRequest.SchName,
		SubName:      subRankRequest.SubName,
		SubCountRank: subRankRequest.SubCountRank,
		SubWorldRank: subRankRequest.SubWorldRank,
	}
	if err := sr.DB.Create(&schRank).Error; err != nil {
		response.Fail(c, nil, "添加失败")
		return
	}
	response.Success(c, gin.H{"id": schRank.ID}, "添加成功")
}

func (sr SubRankController) Update(c *gin.Context) {
	var subRankRequest model.SubRank
	// 数据验证
	if err := c.ShouldBindJSON(&subRankRequest); err != nil {
		response.Fail(c, nil, "数据错误")
		return
	}
	// 获取path中的id
	subRankId := c.Params.ByName("id")
	// 查找学科排名
	var subRank model.SubRank
	if sr.DB.Where("id = ?", subRankId).First(&subRank).RecordNotFound() {
		response.Fail(c, nil, "学科排名不存在")
		return
	}
	// 判断学科排名是否重复添加
	var existSubRank []model.SubRank
	sr.DB.Where("(year = ? AND sub_Count_Rank = ?) OR (year = ? AND sub_World_Rank = ?)", subRank.Year, subRankRequest.SubCountRank, subRank.Year, subRankRequest.SubWorldRank).Find(&existSubRank)
	if len(existSubRank) > 1 {
		response.Fail(c, nil, "学科排名已存在")
		return
	}
	// 修改学科排名信息
	if err := sr.DB.Model(&subRank).Update(subRankRequest).Error; err != nil {
		response.Fail(c, nil, "修改失败")
		return
	}
	response.Success(c, nil, "修改成功")
}

func (sr SubRankController) Delete(c *gin.Context) {
	// 获取path中的id
	subRankId := c.Params.ByName("id")
	fmt.Println(subRankId)
	// 查找学科排名信息
	var subRank model.SubRank
	if sr.DB.Where("id = ?", subRankId).First(&subRank).RecordNotFound() {
		response.Fail(c, nil, "学科排名不存在")
		return
	}
	// 删除学科排名信息
	if err := sr.DB.Delete(&subRank).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}
	response.Success(c, nil, "删除成功")
}

func (sr SubRankController) Search(c *gin.Context) {
	// 获取年份、大学名称、学科名称、分页参数
	year := c.DefaultQuery("year", "null")
	schName := c.DefaultQuery("schName", "")
	subName := c.DefaultQuery("subName", "")
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "7"))
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
	var subRank []model.SubRank
	// 学科总数
	var count int
	// 查询学科
	switch len(args) {
	case 0:
		sr.DB.Table("sub_ranks").Select("*").Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&subRank)
		sr.DB.Model(model.SubRank{}).Count(&count)
	case 1:
		sr.DB.Table("sub_ranks").Select("*").Where(querystr, args[0]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&subRank)
		sr.DB.Model(model.SubRank{}).Where(querystr, args[0]).Count(&count)
	case 2:
		sr.DB.Table("sub_ranks").Select("*").Where(querystr, args[0], args[1]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&subRank)
		sr.DB.Model(model.SubRank{}).Where(querystr, args[0], args[1]).Count(&count)
	case 3:
		sr.DB.Table("sub_ranks").Select("*").Where(querystr, args[0], args[1], args[2]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&subRank)
		sr.DB.Model(model.SubRank{}).Where(querystr, args[0], args[1], args[2]).Count(&count)
	}
	// 展示学科排名列表
	response.Success(c, gin.H{"subRank": subRank, "count": count}, "查找成功")
}

func NewSubRankController() ISubRankController {
	db := common.GetDB()
	db.AutoMigrate(model.SubRank{})
	return SubRankController{DB: db}
}
