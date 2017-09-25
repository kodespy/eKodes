package controllers

import (
	//"golang.org/x/crypto/bcrypt"
	"github.com/revel/revel"
	"github.com/kodespy/ekodes/app/routes"
	m "github.com/kodespy/ekodes/app/models"
)

type App struct {
	*revel.Controller
}
type Ciudades struct {
	*revel.Controller
}

func (c Application) connected() *m.Usuario {
	if c.ViewArgs["user"] != nil {
		return c.ViewArgs["user"].(*m.Usuario)
	}
	//if usuario, ok := c.Session["user"]; ok {
	//	return c.getUser(username)
	//}
	return nil
}


func (c App) Login(email, clave string) revel.Result {
	usuario, err := getUsuario(email, clave)
	revel.INFO.Println("--- Verificando usuarios")
	revel.INFO.Println(usuario)
	revel.INFO.Println(err)
	revel.INFO.Println("Datos")
	revel.INFO.Println(email)
	revel.INFO.Println(clave)
	revel.INFO.Println("------------------")
	if  err == nil {
		//err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		//if err == nil {
			c.Session["user"] = email
			//if remember {
			//	c.Session.SetDefaultExpiration()
			//} else {
			//	c.Session.SetNoExpiration()
			//}
			c.Flash.Success("Welcome, " + email)
			revel.INFO.Println("--- Aqui")
			return c.Redirect(routes.App.Index())
		//}
	}

	revel.INFO.Println("--- Por Aqui")
	c.Flash.Out["email"] = email
	c.Flash.Error("No se pudo autenticar el usuario...")
	return c.Render()
}
func (c App) Index() revel.Result {

	estado, err := c.Session["user"]
	revel.INFO.Println("--------------------")
	revel.INFO.Println(estado)
	revel.INFO.Println("--------------------")
	if err != true {
	    return c.Redirect(routes.App.Login("",""))
	}
	//c.Flash.Error("Please log in first")	
	return c.Render()
	//return c.Render()
}

func (c Ciudades) Lista() revel.Result {
	ciudades, err := allCiudad()
	if err != nil {
		panic(err)
	}
	//return c.Redirect(routes.Ciudades.Lista(), ciudades)
	return c.Render(ciudades)
	//return c.Render()
}

func (c App) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	revel.INFO.Println("--- Logout")
	return c.Redirect(routes.App.Index())
}
