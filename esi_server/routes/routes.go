package routes

import (
	"blog_server/controller"
	"blog_server/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	// 允许跨域访问
	r.Use(middleware.CORSMiddleware())
	// 登录
	r.POST("/login", controller.Login)
	// 登录获取信息
	r.GET("/getInfo", middleware.UserMiddleware(), controller.GetInfo)
	// 用户信息的增删查改
	UserRoutes := r.Group("/user")
	UserRoutes.Use(middleware.UserMiddleware())
	UserRoutes.POST("", controller.CreateUser)
	UserRoutes.PUT(":id", controller.UpdateUser)
	UserRoutes.DELETE(":id", controller.DeleteUser)
	UserRoutes.POST("search", controller.SearchUser)
	//大学信息的增删查改
	SchoolRoutes := r.Group("/school")
	SchoolRoutes.Use(middleware.UserMiddleware())
	SchoolController := controller.NewSchoolController()
	SchoolRoutes.POST("", middleware.AdminMiddleware(), SchoolController.Create)
	SchoolRoutes.PUT(":id", middleware.AdminMiddleware(), SchoolController.Update)
	SchoolRoutes.DELETE(":id", middleware.AdminMiddleware(), SchoolController.Delete)
	SchoolRoutes.POST("search", SchoolController.Search)
	SchoolRoutes.GET("", SchoolController.List)
	//大学排名的增删查改
	SchRankRoutes := r.Group("/schRank")
	SchRankRoutes.Use(middleware.UserMiddleware())
	SchRankController := controller.NewSchRankController()
	SchRankRoutes.POST("", middleware.AdminMiddleware(), SchRankController.Create)
	SchRankRoutes.PUT(":id", middleware.AdminMiddleware(), SchRankController.Update)
	SchRankRoutes.DELETE(":id", middleware.AdminMiddleware(), SchRankController.Delete)
	SchRankRoutes.POST("search", SchRankController.Search)
	//学科信息的增删查改
	SubjectRoutes := r.Group("/subject")
	SubjectRoutes.Use(middleware.UserMiddleware())
	SubjectController := controller.NewSubjectController()
	SubjectRoutes.POST("", middleware.AdminMiddleware(), SubjectController.Create)
	SubjectRoutes.PUT(":id", middleware.AdminMiddleware(), SubjectController.Update)
	SubjectRoutes.DELETE(":id", middleware.AdminMiddleware(), SubjectController.Delete)
	SubjectRoutes.POST("search", SubjectController.Search)
	SubjectRoutes.GET("", SubjectController.List)
	//学科排名的增删查改
	SubRankRoutes := r.Group("/subRank")
	SubRankRoutes.Use(middleware.UserMiddleware())
	SubRankController := controller.NewSubRankController()
	SubRankRoutes.POST("", middleware.AdminMiddleware(), SubRankController.Create)
	SubRankRoutes.PUT(":id", middleware.AdminMiddleware(), SubRankController.Update)
	SubRankRoutes.DELETE(":id", middleware.AdminMiddleware(), SubRankController.Delete)
	SubRankRoutes.POST("search", SubRankController.Search)
	//学科排名的增删查改
	PaperRoutes := r.Group("/paper")
	PaperRoutes.Use(middleware.UserMiddleware())
	PaperController := controller.NewPaperController()
	PaperRoutes.POST("", middleware.AdminMiddleware(), PaperController.Create)
	PaperRoutes.PUT(":id", middleware.AdminMiddleware(), PaperController.Update)
	PaperRoutes.DELETE(":id", middleware.AdminMiddleware(), PaperController.Delete)
	PaperRoutes.POST("search", PaperController.Search)
	// 返回excel文件
	r.GET("/exportSchool", controller.ExportSchool)
	r.GET("/exportSchRank", controller.ExportSchRank)
	r.GET("/exportSubject", controller.ExportSubject)
	r.GET("/exportSubRank", controller.ExportSubRank)
	r.GET("/exportPaper", controller.ExportPaper)
	r.GET("/exportUser", controller.ExportUser)
	return r
}
