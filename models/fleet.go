package models

type Fleet struct {
	ID          int		`json:"id" gorm:"primaryKey"`
	Name		string		`json:"name"`
	Max_Speed	float32		`json:"max_speed"`
}