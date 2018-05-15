package storage


var (
	//TODO describe errors
)


type Giph struct {
	ID   string `gorm:"type:varchar(100);primary_key" json:"id, omitempty"`
	Url  string `gorm:"type:varchar(100)" json:"url, omitempty"`
	Name string `gorm:"type:varchar(100)" json:"title, omitempty"`
}

type Giphs [] Giph