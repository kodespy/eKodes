package controllers

import (
//	"fmt"/
//	r "github.com/revel/revel"
//	"github.com/go-pg/pg"
 //   m "github.com/kodes/ekodes/app/models"
)


// var DB *pg.DB

// func InitDB() {
// 	DB := pg.Connect(&pg.Options{
//       User: "postgres",
//       Password: "sqlcmd",
//       Database: "kodesdb",
//    })

//    err := createSchema(DB)
//    if err != nil {
//    	  r.INFO.Println("DB Error", err)
//       //panic(err)
//    }	
//    // user1 := &m.User{
//    //    Name:   "admin",
//    //    Emails: []string{"admin1@admin", "admin2@admin"},
//    // }
//    // err = DB.Insert(user1)
//    // if err != nil {
//    //    panic(err)
//    // }
//    // fmt.Println(user1)
//    // ciudad1 := &m.Ciudad{
//    //    Nombre:    "Curuguaty",
//    //    UsuarioA: user1,
//    // }
//    // err = DB.Insert(ciudad1)
//    // if err != nil {
//    //    panic(err)
//    // }

//    // var users []m.User
//    // err = DB.Model(&users).Select()
//    // if err != nil {
//    //    panic(err)
//    // }

//    // var ciudades []m.Ciudad
//    // err = DB.Model(&ciudades).Select()
//    // if err != nil {
//    //    panic(err)
//    // }

//     fmt.Println("Init Server.....")
//    // fmt.Println(ciudades)

//    // --- 
   
//     // connstring := fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=disable", "postgres", "sqlcmd", "kodesbd")
//     // r.INFO.Println(connstring)

//     // var err error
//     // DB, err = sql.Open("postgres", connstring)
//     // if err != nil {
//     //     r.INFO.Println("DB Error", err)
//     // }
//     // r.INFO.Println("DB Connected")
//     // err := createSchema(DB)
//     // if err != nil {
//     //   panic(err)
//     // }

// }



// func createSchema(db *pg.DB) error {
//    for _, model := range []interface{}{&m.User{},&m.Ciudad{}} {
//       err := db.CreateTable(model, nil)
//       if err != nil {
//          return err
//       }
//    }
//    return nil
// }


// 	func checkUser(c *r.Controller) r.Result {
// 		r.INFO.Println("---Aqui checkUser")
// 	    return c.Redirect(App.Index)
// 	}

// 	func doNothing(c *r.Controller) r.Result { 
// 		r.INFO.Println("---Aqui doNothing")
// 		return nil 
// 	}


// 	func init() {
// 		r.INFO.Println("--- Data Base Init")
// 		r.OnAppStart(InitDB)
// 		//revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
// 		//revel.InterceptMethod(Application.AddUser, revel.BEFORE)
// 		//revel.InterceptMethod(Hotels.checkUser, revel.BEFORE)
// 		//revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
// 		//revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
// 		//r.InterceptFunc(checkUser, r.BEFORE, &App{})
// 	    //r.InterceptFunc(doNothing, r.AFTER, &App{})
// 	    //r.InterceptFunc(checkUser, r.BEFORE, &AnotherController{})
// 	}