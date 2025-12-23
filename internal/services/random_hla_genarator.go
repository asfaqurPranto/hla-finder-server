package services

import (
	"hla_finder/internal/db"
	"hla_finder/internal/models"
	"math/rand"
	"time"
)

func Random_HLA_Genarator() string {
	Sample_A1 := []string{"A*01:01", "A*02:01", "A*03:01", "A*11:01", "A*23:01", "A*24:02", "A*26:01", "A*30:01"}
	Sample_A2 := []string{"A*02:01", "A*03:01", "A*11:01", "A*24:02", "A*26:01", "A*29:01", "A*30:01", "A*33:01"}
	Sample_B1 := []string{"B*07:02", "B*08:01", "B*15:01", "B*27:05", "B*35:01", "B*40:01", "B*44:02", "B*51:01"}
	Sample_B2 := []string{"B*07:02", "B*15:01", "B*27:05", "B*35:01", "B*40:01", "B*44:02", "B*51:01", "B*52:01"}
	Sample_DR1 := []string{"DRB1*01:01", "DRB1*03:01", "DRB1*04:01", "DRB1*07:01", "DRB1*09:01", "DRB1*11:01", "DRB1*13:01", "DRB1*15:01"}
	Sample_DR2 := []string{"DRB1*01:01", "DRB1*03:01", "DRB1*04:01", "DRB1*07:01", "DRB1*09:01", "DRB1*11:01", "DRB1*13:01", "DRB1*15:01"}
	
	db.Connect_MySql_Server()
	rand.Seed(time.Now().UnixNano())
	var i uint
	
	for i=1;i<=25;i++{
		var tem_hla models.HLA
		tem_hla.A1=Sample_A1[rand.Intn(8)]
		tem_hla.A2=Sample_A2[rand.Intn(8)]
		tem_hla.B1=Sample_B1[rand.Intn(8)]
		tem_hla.B2=Sample_B2[rand.Intn(8)]
		tem_hla.DR1=Sample_DR1[rand.Intn(8)]
		tem_hla.DR2=Sample_DR2[rand.Intn(8)]
		tem_hla.UserID=i
		_=db.DB.Create(&tem_hla)
	}
	
	return "auto generation done"

}
