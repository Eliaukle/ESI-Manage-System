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

type PaperController struct {
	DB *gorm.DB
}

type IPaperController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Search(c *gin.Context)
}

func (p PaperController) Create(c *gin.Context) {
	var paperRequest model.Paper
	// 数据验证
	if err := c.ShouldBindJSON(&paperRequest); err != nil {
		response.Fail(c, nil, "数据错误")
		return
	}
	// 判断论文数据是否重复添加
	var existPaper model.Paper
	p.DB.Where("year = ? AND sch_name = ? AND sub_name = ?", paperRequest.Year, paperRequest.SchName, paperRequest.SubName).First(&existPaper)
	if existPaper.ID != 0 {
		response.Fail(c, nil, "论文数据已存在")
		return
	}
	// 添加论文数据
	paper := model.Paper{
		Year:     paperRequest.Year,
		SchName:  paperRequest.SchName,
		SubName:  paperRequest.SubName,
		PaperNum: paperRequest.PaperNum,
		UsedNum:  paperRequest.UsedNum,
	}
	if err := p.DB.Create(&paper).Error; err != nil {
		response.Fail(c, nil, "添加失败")
		return
	}
	response.Success(c, gin.H{"id": paper.ID}, "添加成功")
}

func (p PaperController) Update(c *gin.Context) {
	var paperRequest model.Paper
	// 数据验证
	if err := c.ShouldBindJSON(&paperRequest); err != nil {
		response.Fail(c, nil, "数据错误")
		return
	}
	// 获取path中的id
	paperId := c.Params.ByName("id")
	// 查找论文数据
	var paper model.Paper
	if p.DB.Where("id = ?", paperId).First(&paper).RecordNotFound() {
		response.Fail(c, nil, "论文数据不存在")
		return
	}
	// 修改论文数据
	if err := p.DB.Model(&paper).Update(paperRequest).Error; err != nil {
		response.Fail(c, nil, "修改失败")
		return
	}
	response.Success(c, nil, "修改成功")
}

func (p PaperController) Delete(c *gin.Context) {
	// 获取path中的id
	paperId := c.Params.ByName("id")
	// 查找论文数据
	var paper model.Paper
	if p.DB.Where("id = ?", paperId).First(&paper).RecordNotFound() {
		response.Fail(c, nil, "论文数据不存在")
		return
	}
	// 删除论文数据
	if err := p.DB.Delete(&paper).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}
	response.Success(c, nil, "删除成功")
}

func (p PaperController) Search(c *gin.Context) {
	// 获取年份、大学名称、学科名称、分页参数
	year := c.DefaultQuery("year", "null")
	schName := c.DefaultQuery("schName", "")
	subName := c.DefaultQuery("subName", "")
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
	var paper []model.Paper
	// 论文数据总数
	var count int
	// 查询论文数据
	switch len(args) {
	case 0:
		p.DB.Table("papers").Select("*").Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&paper)
		p.DB.Model(model.Paper{}).Count(&count)
	case 1:
		p.DB.Table("papers").Select("*").Where(querystr, args[0]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&paper)
		p.DB.Model(model.Paper{}).Where(querystr, args[0]).Count(&count)
	case 2:
		p.DB.Table("papers").Select("*").Where(querystr, args[0], args[1]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&paper)
		p.DB.Model(model.Paper{}).Where(querystr, args[0], args[1]).Count(&count)
	case 3:
		p.DB.Table("papers").Select("*").Where(querystr, args[0], args[1], args[2]).Order("id asc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&paper)
		p.DB.Model(model.Paper{}).Where(querystr, args[0], args[1], args[2]).Count(&count)
	}
	// 展示论文数据列表
	response.Success(c, gin.H{"paper": paper, "count": count}, "查找成功")
}

func NewPaperController() IPaperController {
	db := common.GetDB()
	db.AutoMigrate(model.Paper{})
	return PaperController{DB: db}
}
