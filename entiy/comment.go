package entiy

func (Comment) TableName() string {
	return "comment"
}

type Comment struct {
	CId     string `json:"c_id" gorm:"column:c_id"`
	Content string `json:"content" gorm:"column:content"`
}
