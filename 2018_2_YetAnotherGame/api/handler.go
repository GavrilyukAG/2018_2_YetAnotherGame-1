package api
import (
	"2018_2_YetAnotherGame/domain/models"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"log"
	"golang.org/x/net/context"
	//"net/http"

)

func dbSettings() string {
	conf := &DbConfig{}
	_, err:=toml.DecodeFile("./config/DBsettings.toml", conf)
	fmt.Println(err)
	fmt.Printf("%s", conf.String())
	return conf.String()
}
type DbConfig struct {
	Host    string `toml:"host"`
	Port    string `toml:"port"`
	Sslmode string `toml:"sslmode"`
	Dbname  string `toml:"dbname"`
	User    string `toml:"user"`
	Pass    string `toml:"pass"`
}
func (db DbConfig) String() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s "+
		"sslmode=%s user=%s password=%s ",
		db.Host, db.Port, db.Dbname, db.Sslmode, db.User, db.Pass,
	)
}
// Server represents the gRPC server
type Server struct {
}
// CheckSession generates response to a Ping request
func (s *Server) CheckSession(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s", in.Message)



id:=in.Message
	db, err := gorm.Open("postgres", dbSettings())
	if err!=nil{
		fmt.Println(err)
		logrus.Error(err)
	}
	defer db.Close()

	tmp := models.Session{}
	db.Table("sessions").Select("id, email").Where("id = ?", id).Scan(&tmp)
	if tmp.Email == "" {

		return &PingMessage{Message: "Unauthorized"}, nil
	}



	return &PingMessage{Message: "OK"}, nil
}
