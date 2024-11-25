package routes

import (
	"SiAkademik/controllers"
	"SiAkademik/middlewares"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default() // Inisialisasi router

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://siakademik.up.railway.app"}, // Ganti sesuai domain frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//router.Use(middlewares.BasicAuth()) //basicAuth
	router.Use(middlewares.AuditLog())
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
	}

	router.Use(middlewares.JWTAuthMiddleware()) //jwt

	public := router.Group("/")
	{
		public.PUT("user", controllers.UpdateUser)
		public.GET("userprofile", controllers.GetUserProfile)
		public.PUT("userprofile", controllers.UpdateUserProfile)
	}
	admin := router.Group("/admin")
	// Middleware untuk memeriksa role "admin"
	admin.Use(middlewares.CheckRole("admin"))
	{
		admin.POST("/user", controllers.CreateUser)
		admin.POST("/role", controllers.CreateRole)
		admin.DELETE("/role/:id", controllers.DeleteRole)
		admin.POST("/course", controllers.CreateCourse)
		admin.POST("/semesters", controllers.CreateSemester)
	}

	dosen := router.Group("/dosen")
	// Middleware untuk memeriksa role "admin"
	dosen.Use(middlewares.CheckRole("dosen"))
	{
		dosen.GET("/course", controllers.GetCourseByLectureId)
		dosen.POST("/grade", controllers.CreateGrade)
	}

	mahasiswa := router.Group("/mahasiswa")
	// Middleware untuk memeriksa role "admin"
	mahasiswa.Use(middlewares.CheckRole("mahasiswa"))
	{
		mahasiswa.POST("/enrollment", controllers.CreateEnrollment)
		mahasiswa.GET("/gpa", controllers.GetGPA)
		mahasiswa.GET("/course", controllers.GetStudentCourse)
	}

	router.Run(":" + os.Getenv("PORT"))
	return router
}
