package handlers

import (
	"fmt"
	"hla_finder/internal/db"
	"hla_finder/internal/middleware"
	"sort"
	"sync"

	"hla_finder/internal/models"
	"net/http"
	"strconv"
	"time"

	"hla_finder/internal/services"

	"github.com/gin-gonic/gin"
)
func Input_HLA(c *gin.Context){
	
	patient_id_string:=c.Param("patient_id")
	patient_id_int,err:=strconv.Atoi(patient_id_string)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	patient_id:=uint(patient_id_int)

	var patient models.User
	result:=db.DB.First(&patient,"ID=?",patient_id)
	if result.Error!=nil || patient.ID==0{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"Patient not found",
		})
		return
	}

	var hla_req models.HLA_Request
	err=c.BindJSON(&hla_req)
	if err!=nil{
		///////
	}

	hla:=models.HLA{

		A1:hla_req.A1,
		A2:hla_req.A2,

		B1:hla_req.B1,
		B2:hla_req.B2,

		DR1:hla_req.DR1,
		DR2:hla_req.DR2,
		UserID: patient_id,

	}
	result=db.DB.Create(&hla)
	if result.Error!=nil || hla.ID==0{
		/////
		c.JSON(http.StatusBadRequest,gin.H{
			"message":result.Error.Error(),
		})
		return

	}
	
	c.JSON(http.StatusCreated,gin.H{
		"message":"hla created successfully",
		"hla":hla,
	})


}


func Donation_Date(c *gin.Context){

	patient_id_string:=c.Param("patient_id")
	patient_id_int,err:=strconv.Atoi(patient_id_string)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	patient_id:=uint(patient_id_int)

	var patient models.User
	result:=db.DB.First(&patient,"ID=?",patient_id)
	if result.Error!=nil || patient.ID==0{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"Patient not found",
		})
		return
	}

	type Date_Request struct{
		Date time.Time `json:"date" binding:"required"`
	}
	var date_request Date_Request
	err=c.BindJSON(&date_request)
	if err!=nil{
		return
	}
	// handle error
	patient.LastDontated=&date_request.Date
	result=db.DB.Save(patient)
	if result.Error!=nil{
		return
	}
	//handle error
	c.JSON(http.StatusOK,gin.H{
		"message":"date added",
	})



}


func Delete_Patient(c *gin.Context){

	patient_id_string:=c.Param("patient_id")
	patient_id_int,err:=strconv.Atoi(patient_id_string)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	patient_id:=uint(patient_id_int)
	var patient models.User
	result:=db.DB.First(&patient,"ID=?",patient_id)
	if result.Error!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"User not found",
		})
		return
	}
	if patient.Admin{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"You can not delete admin",
		})
		return
	}

	//delete user hla from hla table
	result=db.DB.Where("user_id=?",patient_id).Delete(&models.HLA{})
	// if result.Error!=nil{
	// 	c.JSON(http.StatusBadRequest,gin.H{
	// 		"message":"user deleted but hla was not found",
	// 	})
	// 	return
	// }
	//delete user from user table
	result=db.DB.Where("id = ?", patient_id).Delete(&models.User{})
	if result.Error!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"User not found",
		})
		return
	}
	
	c.JSON(http.StatusOK,gin.H{
		"message":"User deleted",
	})


	
}

func Show_HLA(c *gin.Context){
	user,err:=middleware.UserInfo(c)
	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"message":err.Error(),
		})
		return 
	}
	// Find hla record hla table using user.ID
	var hla_patient models.HLA
	result:=db.DB.First(&hla_patient,"user_id=?",user.ID)
	if result.Error!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"message":"we can not find your hla record in our databaseplease, contact your doctor/admin",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"hla":hla_patient,
	})
	fmt.Println(hla_patient)
}





func Find_Match(c *gin.Context)(error,[]models.HLA_Match_Report){
	user,err:=middleware.UserInfo(c)
	if err!=nil{
		return err,nil
	}
	
	var patient_hla models.HLA
	result:=db.DB.First(&patient_hla,"user_id=?",user.ID)
	if result.Error!=nil{
		return err,nil
	}
	
	var all_hla []models.HLA
	result=db.DB.Find(&all_hla)
	if result.Error!=nil{
		return result.Error,nil
	}
	

	var users []models.User
	db.DB.Find(&users) //loading all users and saved them in hashed map by their id
	//  jete bar bar query na kora lage time c. reduced to o(n^2)-> o(n)
	userMap:=make(map[uint]models.User)
	for _,u:= range users{
		userMap[u.ID]=u
	}


	var Matched_Report []models.HLA_Match_Report

	var wg sync.WaitGroup
	var mu sync.Mutex
	

	for _,hla:= range all_hla{
		if hla.UserID==patient_hla.UserID{
			continue
		}
		wg.Add(1) // concurrently match korbo,single match valoi time khay(specially distance api call korte jay)
		go func (hla models.HLA){
			defer wg.Done()
			var tem models.HLA_Match_Report
			tem.User_ID=hla.UserID
			if hla.A1==patient_hla.A1{
				tem.Matched++
				tem.A1=true
			}
			if hla.A2==patient_hla.A2{
				tem.Matched++
				tem.A2=true
			}
			if hla.B1==patient_hla.B1{
				tem.Matched++
				tem.B1=true
			}
			if hla.B2==patient_hla.B2{
				tem.Matched++
				tem.B2=true
			}
			if hla.DR1==patient_hla.DR1{
				tem.Matched++
				tem.DR1=true
			}
			if hla.DR2==patient_hla.DR2{
				tem.Matched++
				tem.DR2=true
			}
			//find patient and donor
			var donor models.User
			
			//db.DB.First(&donor,"ID=?",hla.UserID)
			donor=userMap[hla.UserID]
			
			//ei distance call valoi time ney, redis use korte hobe must
			tem.Distance,_=services.City_Distance(user.City,donor.City)

			tem.City=donor.City
			tem.User_Name=donor.Name
			mu.Lock()
			Matched_Report=append(Matched_Report, tem)
			mu.Unlock()
		}(hla)
	}
	wg.Wait()
	return nil,Matched_Report

	
	
}

func Get_Report(c *gin.Context){
	err,Matched_Report:=Find_Match(c)
	if err!=nil || len(Matched_Report)==0{
		c.JSON(http.StatusNotFound,gin.H{
			"message":"not found",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"Report":Matched_Report,
	})
}

func Sort_Match(c *gin.Context){
	
	type SortResponse struct{
		Sortby string `json:"sortby" binding:"required"`
	}
	var response SortResponse
	err:=c.BindJSON(&response)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	err,Matched_Report:=Find_Match(c)
	if err!=nil || len(Matched_Report)==0{
		c.JSON(http.StatusNotFound,gin.H{
			"message":"not found",
		})
		return
	}

	if response.Sortby=="distance"{
		sort.Slice(Matched_Report,func(i,j int) bool{
			return Matched_Report[i].Distance<Matched_Report[j].Distance
		})

	}else if response.Sortby=="best match"{
		sort.Slice(Matched_Report,func(i,j int) bool{
			return Matched_Report[i].Matched>Matched_Report[j].Matched
		})
	}
	
	c.JSON(http.StatusOK,gin.H{
		"Report":Matched_Report,
	})
}
