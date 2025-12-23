package handlers

import (
	"hla_finder/internal/db"
	"hla_finder/internal/middleware"
	"hla_finder/internal/models"
	"net/http"
	//"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context){
	var register_request models.RegisterRequset
	err:=c.BindJSON(&register_request)
	if err!=nil{
		c.JSON(http.StatusBadRequest,
		gin.H{
			"message":err.Error(),
		})
		return
	}

	//hash the password
	hashed_b,err:=bcrypt.GenerateFromPassword([]byte(register_request.Password),10)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":"password hashed failed",
		})
	}
	hashed:=string(hashed_b)


	//Save it to database
	user:=models.User{
		Name:register_request.Name,
		Email: register_request.Email,
		Password: hashed,
		Admin: false,
		City:register_request.City,
	}
	result:=db.DB.Create(&user)
	if result.Error!=nil{
		c.JSON(http.StatusBadRequest,
		gin.H{
			"message":result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated,
	gin.H{
		"message":"User Created Successfully",
	})

}

func Login(c *gin.Context){
	var login_request models.LoginRequset
	
	err:=c.BindJSON(&login_request)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	
	//get the user from database with the requested email
	var user models.User
	db.DB.First(&user,"Email=?",login_request.Email)
	if user.ID==0{
		c.JSON(http.StatusNotFound,gin.H{
			"message":"User not found with this email",
		})
		return
	}
	
	//hases the input pass and compare it with db saved one
	err=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(login_request.Password))
	//not matched
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"wrong password",
		})
		return
	}
	//matched

	//create token struct containing signing meteod sub and expire time
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"sub":user.ID,
		"exp":time.Now().Add(time.Hour*24*30).Unix(),

	})
	SECRET_KEY:="helloworld"
	//generate token string (header.payload.signature) from token
	tokenString,err:=token.SignedString([]byte(SECRET_KEY))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization",tokenString,3600*24*7,"","",false,true)

	c.JSON(http.StatusOK,gin.H{
		"message":"login successful",
	})
	

}

func UserInfo(c *gin.Context){
	user,err:=middleware.UserInfo(c)
	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"message":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"user":user,
	})
}

func UpdateUserInfo(c *gin.Context){
	user,err:=middleware.UserInfo(c)
	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"message":err.Error(),
		})
		return
	}

	type UpdateRequest struct {
		Name     string `json:"name"`
		City 	string `json:"city"`
	}

	var update_request UpdateRequest

	err=c.BindJSON(&update_request)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	if update_request.Name!=""{
		user.Name=update_request.Name
	}
	if update_request.City!=""{
		user.City=update_request.City
	}
	
	result:=db.DB.Save(user)
	if result.Error!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"user info updated",
		"user":user,
	})

}
