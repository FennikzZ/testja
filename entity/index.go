package entity
import(
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Phone     string `valid:"required~Phone is required, stringlength(10|10)"`
}