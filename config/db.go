package config
import(
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
func SetupDB() *gorm.DB{
	user:="root"
	pass:=""
	port:="3306"
	host:="localhost"
	dbname:="blog"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",user,pass,host,port,dbname)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
		}
	return db
}
