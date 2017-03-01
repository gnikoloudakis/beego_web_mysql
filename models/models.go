package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Arduino struct{
	Id 	int			`orm:"auto"`
	Name 	string			`orm:"unique;size(100)"`
	IP 	string			`orm:"size(20)"`
	//Measurements	*Measurements	`orm:"rel(fk)"`
}

type Measurements struct{
	Id 		int		`orm:"auto"`
	Timestamp	time.Time 	`orm:"auto_now_add;type(datetime)"`
	Tempereature	float32		`orm:"size(10)"`
	Arduino 	*Arduino 	`orm:"rel(fk);on_delete(cascade)"`
}

func init(){
	orm.RegisterModel(new(Arduino), new(Measurements))
}