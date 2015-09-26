package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:WebsocketController"] = append(beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:WebsocketController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:WebsocketController"] = append(beego.GlobalControllerRouter["github.com/Iverson/pusher-api/controllers:WebsocketController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

}