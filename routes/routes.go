package routes

import (
	"SiAkademik/controllers"
	"SiAkademik/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default() // Inisialisasi router

	router.Use(controllers.Handle())
	public := router.Group("/", middlewares.BasicAuth())
	{
		public.PUT("user", controllers.UpdateUser)
		public.GET("userprofile", controllers.GetUserProfile)
		public.PUT("userprofile", controllers.UpdateUserProfile)
	}
	admin := router.Group("/admin", middlewares.BasicAuth())
	// Middleware untuk memeriksa role "admin"
	admin.Use(middlewares.CheckRole("admin"))
	{
		admin.POST("/user", controllers.CreateUser)
		admin.POST("/role", controllers.CreateRole)
		admin.DELETE("/role/:id", controllers.DeleteRole)
		admin.POST("/course", controllers.CreateCourse)
		admin.POST("/semesters", controllers.CreateSemester)
	}

	dosen := router.Group("/dosen", middlewares.BasicAuth())
	// Middleware untuk memeriksa role "admin"
	dosen.Use(middlewares.CheckRole("dosen"))
	{
		dosen.GET("/course", controllers.GetCourseByLectureId)
		dosen.POST("/grade", controllers.CreateGrade)
	}

	mahasiswa := router.Group("/mahasiswa", middlewares.BasicAuth())
	// Middleware untuk memeriksa role "admin"
	mahasiswa.Use(middlewares.CheckRole("mahasiswa"))
	{
		mahasiswa.POST("/enrollment", controllers.CreateEnrollment)
		mahasiswa.GET("/gpa", controllers.GetGPA)
	}

	router.Run(":8080")
	return router
}
