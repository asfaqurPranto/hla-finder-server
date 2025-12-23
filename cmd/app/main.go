package main

import (
	"hla_finder/internal/db"
	"hla_finder/internal/handlers"
	"hla_finder/internal/middleware"
	"hla_finder/internal/services"

	"github.com/gin-gonic/gin"
	
	
)

func init(){
	db.Connect_MySql_Server()
	db.Create_Schema()
	services.InitRedis()
	
}


func main(){
	router:=gin.Default()
	
	public_route:=router.Group("")
	{
		public_route.POST("/register",handlers.Register)
		public_route.POST("/login",handlers.Login)
	}
	
	user_route:=router.Group("/user")
	{
		
		user_route.GET("/info",middleware.Login_Required,handlers.UserInfo)
		user_route.PUT("/update",middleware.Login_Required,handlers.UpdateUserInfo)
		//can update name now,will add city later
		
	}
	admin_route:=router.Group("/admin")
	{ 
		admin_route.DELETE("delete/:patient_id",middleware.Login_Required,middleware.Admin_Required,handlers.Delete_Patient)
		admin_route.POST("/donation_date/:patient_id",middleware.Login_Required,middleware.Admin_Required,handlers.Donation_Date)
		admin_route.POST("/input_hla/:patient_id",middleware.Login_Required,middleware.Admin_Required,handlers.Input_HLA)
	}
	hla_route:=router.Group("/hla")
	{
		hla_route.GET("/",middleware.Login_Required,handlers.Show_HLA)
		hla_route.GET("/match",middleware.Login_Required,handlers.Get_Report)
		hla_route.GET("/match/sort",middleware.Login_Required,handlers.Sort_Match)
	}




	router.Run()
}