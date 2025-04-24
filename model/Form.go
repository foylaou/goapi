package model

// Form 表單資料模型
type Form struct {
	ID         string `gorm:"type:uniqueidentifier;primaryKey;default:newsequentialid()" json:"id"`
	Code       string `gorm:"type:nvarchar(255);not null" json:"code"`
	Title      string `gorm:"type:nvarchar(255);not null" json:"title"`
	JsonSchema string `gorm:"type:text;not null" json:"jsonSchema"`
	CreatedAt  string `gorm:"type:datetime2;default:getdate()" json:"createdAt"`
	UpdateAt   string `gorm:"type:datetime2;default:getdate()" json:"updateAt"`
}
