package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestController struct {
	beego.Controller
}

func (j *RequestController) Post() {
	p := make([]byte, j.Ctx.Request.ContentLength)
	_, err := j.Ctx.Request.Body.Read(p)
	fmt.Println("err", err)
	if err == nil {
		var newUser User
		err1 := json.Unmarshal(p, &newUser)
		if err1 == nil {
			fmt.Println("user name:", newUser.Username)

		} else {
			fmt.Println("Unable to unmarshall the JSON request", err1);
		}
	}
	//j.ServeJSON()
	j.TplName = "index.tpl"
}

//var ob models.Object
//json.Unmarshal(j.Ctx.Input.RequestBody, &ob)
//objectid := models.AddOne(ob)
//j.Data["json"] = map[string]interface{}{"ObjectId": objectid }
//j.ServeJSON()
//}