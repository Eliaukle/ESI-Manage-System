package controller

import (
	"blog_server/common"
	"blog_server/model"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
)

func ExportSchool(c *gin.Context) {
	db := common.GetDB()
	// 获取大学名称、学校类别、办学性质
	schName := c.DefaultQuery("schName", "")
	schType := c.DefaultQuery("schType", "")
	schManage := c.DefaultQuery("schManage", "")
	var query []string
	var args []string
	// 若大学名称存在
	if schName != "" {
		query = append(query, "sch_name LIKE ?")
		args = append(args, "%"+schName+"%")
	}
	// 若学校类别存在
	if schType != "null" {
		query = append(query, "sch_type = ?")
		args = append(args, schType)
	}
	// 若办学性质存在
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
	// 查询大学信息
	switch len(args) {
	case 0:
		db.Table("schools").Select("*").Order("id asc").Find(&school)
	case 1:
		db.Table("schools").Select("*").Where(querystr, args[0]).Order("id asc").Find(&school)
	case 2:
		db.Table("schools").Select("*").Where(querystr, args[0], args[1]).Order("id asc").Find(&school)
	case 3:
		db.Table("schools").Select("*").Where(querystr, args[0], args[1], args[2]).Order("id asc").Find(&school)
	}
	// 创建excel表
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "大学编号")
	xlsx.SetCellValue("Sheet1", "B1", "大学名称")
	xlsx.SetCellValue("Sheet1", "C1", "学校类别")
	xlsx.SetCellValue("Sheet1", "D1", "办学性质")
	for i, v := range school {
		curIndex := i + 2
		strIndex := strconv.Itoa(curIndex)
		xlsx.SetCellValue("Sheet1", "A"+strIndex, v.ID)
		xlsx.SetCellValue("Sheet1", "B"+strIndex, v.SchName)
		xlsx.SetCellValue("Sheet1", "C"+strIndex, v.SchType)
		xlsx.SetCellValue("Sheet1", "D"+strIndex, v.SchManage)
	}
	// 返回文件流
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"result.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	_ = xlsx.Write(c.Writer)
}

func ExportSchRank(c *gin.Context) {
	db := common.GetDB()
	// 获取年份、大学名称
	year := c.DefaultQuery("year", "null")
	schName := c.DefaultQuery("schName", "")
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
	// 查询大学排名
	switch len(args) {
	case 0:
		db.Table("sch_ranks").Select("*").Order("id asc").Find(&schRank)
	case 1:
		db.Table("sch_ranks").Select("*").Where(querystr, args[0]).Order("id asc").Find(&schRank)
	case 2:
		db.Table("sch_ranks").Select("*").Where(querystr, args[0], args[1]).Order("id asc").Find(&schRank)
	}
	// 创建excel表
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "编号")
	xlsx.SetCellValue("Sheet1", "B1", "年份")
	xlsx.SetCellValue("Sheet1", "C1", "大学名称")
	xlsx.SetCellValue("Sheet1", "D1", "大学全国排名")
	xlsx.SetCellValue("Sheet1", "E1", "大学全球排名")
	for i, v := range schRank {
		curIndex := i + 2
		strIndex := strconv.Itoa(curIndex)
		xlsx.SetCellValue("Sheet1", "A"+strIndex, v.ID)
		xlsx.SetCellValue("Sheet1", "B"+strIndex, v.Year)
		xlsx.SetCellValue("Sheet1", "C"+strIndex, v.SchName)
		xlsx.SetCellValue("Sheet1", "D"+strIndex, v.SchCountRank)
		xlsx.SetCellValue("Sheet1", "E"+strIndex, v.SchWorldRank)
	}
	// 返回文件流
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"result.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	_ = xlsx.Write(c.Writer)
}

func ExportSubject(c *gin.Context) {
	db := common.GetDB()
	// 获取学科名称、学科类别
	subName := c.DefaultQuery("subName", "")
	subType := c.DefaultQuery("subType", "")
	var query []string
	var args []string
	// 若学科名称存在
	if subName != "" {
		query = append(query, "sub_name LIKE ?")
		args = append(args, "%"+subName+"%")
	}
	// 若学科类别存在
	if subType != "null" {
		query = append(query, "sub_type = ?")
		args = append(args, subType)
	}
	// 拼接字符串
	var querystr string
	if len(query) > 0 {
		querystr = strings.Join(query, " AND ")
	}
	var subject []model.Subject
	// 查询学科
	switch len(args) {
	case 0:
		db.Table("subjects").Select("*").Order("id asc").Find(&subject)
	case 1:
		db.Table("subjects").Select("*").Where(querystr, args[0]).Order("id asc").Find(&subject)
	case 2:
		db.Table("subjects").Select("*").Where(querystr, args[0], args[1]).Order("id asc").Find(&subject)
	case 3:
		db.Table("subjects").Select("*").Where(querystr, args[0], args[1], args[2]).Order("id asc").Find(&subject)
	}
	// 创建excel表
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "学科编号")
	xlsx.SetCellValue("Sheet1", "B1", "学科名称")
	xlsx.SetCellValue("Sheet1", "C1", "学科类别")
	for i, v := range subject {
		curIndex := i + 2
		strIndex := strconv.Itoa(curIndex)
		xlsx.SetCellValue("Sheet1", "A"+strIndex, v.ID)
		xlsx.SetCellValue("Sheet1", "B"+strIndex, v.SubName)
		xlsx.SetCellValue("Sheet1", "C"+strIndex, v.SubType)
	}
	// 返回文件流
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"result.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	_ = xlsx.Write(c.Writer)
}

func ExportSubRank(c *gin.Context) {
	db := common.GetDB()
	// 获取年份、大学名称、学科名称
	year := c.DefaultQuery("year", "null")
	schName := c.DefaultQuery("schName", "")
	subName := c.DefaultQuery("subName", "")
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
	// 页面内容
	var subRank []model.SubRank
	// 查询学科
	switch len(args) {
	case 0:
		db.Table("sub_ranks").Select("*").Order("id asc").Find(&subRank)
	case 1:
		db.Table("sub_ranks").Select("*").Where(querystr, args[0]).Order("id asc").Find(&subRank)
	case 2:
		db.Table("sub_ranks").Select("*").Where(querystr, args[0], args[1]).Order("id asc").Find(&subRank)
	case 3:
		db.Table("sub_ranks").Select("*").Where(querystr, args[0], args[1], args[2]).Order("id asc").Find(&subRank)
	}
	// 创建excel表
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "编号")
	xlsx.SetCellValue("Sheet1", "B1", "年份")
	xlsx.SetCellValue("Sheet1", "C1", "大学名称")
	xlsx.SetCellValue("Sheet1", "D1", "学科名称")
	xlsx.SetCellValue("Sheet1", "E1", "学科全国排名")
	xlsx.SetCellValue("Sheet1", "F1", "学科全球排名")
	for i, v := range subRank {
		curIndex := i + 2
		strIndex := strconv.Itoa(curIndex)
		xlsx.SetCellValue("Sheet1", "A"+strIndex, v.ID)
		xlsx.SetCellValue("Sheet1", "B"+strIndex, v.Year)
		xlsx.SetCellValue("Sheet1", "C"+strIndex, v.SchName)
		xlsx.SetCellValue("Sheet1", "D"+strIndex, v.SubName)
		xlsx.SetCellValue("Sheet1", "E"+strIndex, v.SubCountRank)
		xlsx.SetCellValue("Sheet1", "F"+strIndex, v.SubWorldRank)
	}
	// 返回文件流
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"result.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	_ = xlsx.Write(c.Writer)
}

func ExportPaper(c *gin.Context) {
	db := common.GetDB()
	// 获取年份、大学名称、学科名称
	year := c.DefaultQuery("year", "null")
	schName := c.DefaultQuery("schName", "")
	subName := c.DefaultQuery("subName", "")
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
	// 页面内容
	var paper []model.Paper
	// 查询论文数据
	switch len(args) {
	case 0:
		db.Table("papers").Select("*").Order("id asc").Find(&paper)
	case 1:
		db.Table("papers").Select("*").Where(querystr, args[0]).Order("id asc").Find(&paper)
	case 2:
		db.Table("papers").Select("*").Where(querystr, args[0], args[1]).Order("id asc").Find(&paper)
	case 3:
		db.Table("papers").Select("*").Where(querystr, args[0], args[1], args[2]).Order("id asc").Find(&paper)
	}
	// 创建excel表
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "编号")
	xlsx.SetCellValue("Sheet1", "B1", "年份")
	xlsx.SetCellValue("Sheet1", "C1", "大学名称")
	xlsx.SetCellValue("Sheet1", "D1", "学科名称")
	xlsx.SetCellValue("Sheet1", "E1", "论文数")
	xlsx.SetCellValue("Sheet1", "F1", "他引数")
	for i, v := range paper {
		curIndex := i + 2
		strIndex := strconv.Itoa(curIndex)
		xlsx.SetCellValue("Sheet1", "A"+strIndex, v.ID)
		xlsx.SetCellValue("Sheet1", "B"+strIndex, v.Year)
		xlsx.SetCellValue("Sheet1", "C"+strIndex, v.SchName)
		xlsx.SetCellValue("Sheet1", "D"+strIndex, v.SubName)
		xlsx.SetCellValue("Sheet1", "E"+strIndex, v.PaperNum)
		xlsx.SetCellValue("Sheet1", "F"+strIndex, v.UsedNum)
	}
	// 返回文件流
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"result.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	_ = xlsx.Write(c.Writer)
}

func ExportUser(c *gin.Context) {
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
	// 创建excel表
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "用户编号")
	xlsx.SetCellValue("Sheet1", "B1", "用户名")
	xlsx.SetCellValue("Sheet1", "C1", "用户身份")
	for i, v := range user {
		curIndex := i + 2
		strIndex := strconv.Itoa(curIndex)
		xlsx.SetCellValue("Sheet1", "A"+strIndex, v.ID)
		xlsx.SetCellValue("Sheet1", "B"+strIndex, v.UserName)
		xlsx.SetCellValue("Sheet1", "C"+strIndex, v.Identify)
	}
	// 返回文件流
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"result.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	_ = xlsx.Write(c.Writer)
}
