package middleware

import (
	"errors"
	"net/http"
	"time"

	"hla_finder/internal/db"
	"hla_finder/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login_Required(c *gin.Context){
	
	tokenString,err:=c.Cookie("Authorization")
	
	//1.token string not found inside browser
	if err!=nil{
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//2.Found now , Decode/validate it 
	SECRET_KEY:="helloworld"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(SECRET_KEY), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		//log.Fatal(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//3 Check exp
		if float64(time.Now().Unix())> claims["exp"].(float64){
			c.AbortWithStatus(http.StatusRequestTimeout)
			
		}
		//4 find the user with token sub(sub===id)
		var user models.User
		db.DB.First(&user,claims["sub"])

		if user.ID==0{
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//Attach to request
		c.Set("user",user)

		//continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}


}

func UserInfo(c *gin.Context) (models.User,error){
	userInterface,exists:=c.Get("user")
	if !exists{
		return models.User{} ,errors.New("invalid user type")
	}
	user:=userInterface.(models.User)
	return user,nil
}

func Admin_Required(c *gin.Context){
	user,err:=UserInfo(c)
	if err!=nil{
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if !user.Admin{
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()

}