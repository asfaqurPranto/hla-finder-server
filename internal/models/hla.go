package models


type HLA struct{
	ID uint `gorm:"primaryKey" json:"-"`
	A1 string `json:"a1" binding:"required"`
	A2 string `json:"a2" binding:"required"`
	
	B1 string `json:"b1" binding:"required"`
	B2 string `json:"b2" binding:"required"`

	DR1 string `json:"dr1" binding:"required"`
	DR2 string `json:"dr2" binding:"required"`
	UserID uint `gorm:"unique"`
	User User `gorm:"foreignKey:UserID" json:"-"`
}

type HLA_Request struct{

	A1 string `json:"a1" binding:"required"`
	A2 string `json:"a2" binding:"required"`
	
	B1 string `json:"b1" binding:"required"`
	B2 string `json:"b2" binding:"required"`

	DR1 string `json:"dr1" binding:"required"`
	DR2 string `json:"dr2" binding:"required"`
}

type HLA_Match_Report struct{
	User_ID uint `json:"-"` 
	User_Name string `json:"user_name"`
	Matched uint	`json:"matched"`
	Distance int	`json:"distance"`
	
	A1 bool	`json:"a1"`
	A2 bool	`json:"a2"`
	B1 bool	`json:"b1"`
	B2 bool	`json:"b2"`
	DR1 bool `json:"dr1"`
	DR2 bool `json:"dr2"`
	City string `json:"city"`
}