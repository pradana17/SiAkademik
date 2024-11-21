package routes

import (
	"SiAkademik/controllers"
	"SiAkademik/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default() // Inisialisasi router

	// Rute publik
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "Welcome to the API!"})
	// })

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
		//dosen.PUT("/grade")
	}

	// Grup Admin Role

	// {
	// 	userRoutes.GET("/", controllers.GetAllUsers)
	// 	userRoutes.GET("/:id", controllers.GetUserByID)
	// 	userRoutes.PUT("/:id", controllers.UpdateUser)
	// 	userRoutes.DELETE("/:id", controllers.DeleteUser)
	// }

	// roleRoutes := router.Group("/roles")
	// roleRoutes.POST("/", controllers.CreateRole)

	// // Grup Auth
	// authRoutes := router.Group("/auth")
	// {
	// 	authRoutes.POST("/login", controllers.LoginUser)
	// 	authRoutes.POST("/register", controllers.RegisterUser)
	// }

	// // Grup dengan middleware autentikasi
	// securedRoutes := router.Group("/secured", middleware.AuthMiddleware())
	// {
	// 	securedRoutes.GET("/profile", controllers.GetProfile)
	// }
	router.Run(":8080")
	return router
}
