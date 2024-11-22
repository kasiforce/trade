package model

type School struct {
	SchoolID   int    `gorm:"primaryKey;autoIncrement;column:schoolID"`
	SchoolName string `gorm:"type:varchar(80);not null;column:schoolName"`
	MailSuffix string `gorm:"type:varchar(30);not null;column:mailSuffix"`
}

func (School) TableName() string {
	return "school"
}
