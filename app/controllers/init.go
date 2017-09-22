package controllers

import (
	r "github.com/revel/revel"
	//"github.com/go-pg/pg"
	"database/sql"
 //   m "github.com/kodespy/ekodes/app/models"
    "fmt"
    
)

//var db *pg.DB
var db *sql.DB

func InitDB() {
	// r.INFO.Println("DB Init///")
	// db := pg.Connect(&pg.Options{
 //      User: "postgres",
 //      Password: "sqlcmd",
 //      Database: "kodesdb",
 //   })
	// if db != nil {
	// 	r.INFO.Println()
	// 	r.INFO.Println(db)	
	// }
	connstring := fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=disable", "postgres", "sqlcmd", "kodesdb")

	var err error
	db, err = sql.Open("postgres", connstring)
	if err != nil {
		r.INFO.Println("DB Error", err)
		panic(err)
	}
	r.INFO.Println("DB Connected")

}
func init() {
	r.INFO.Println("--- Data Base Init")
	r.OnAppStart(InitDB)
		//revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
		//revel.InterceptMethod(Application.AddUser, revel.BEFORE)
		//revel.InterceptMethod(Hotels.checkUser, revel.BEFORE)
		//revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
		//revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
		//r.InterceptFunc(checkUser, r.BEFORE, &App{})
	    //r.InterceptFunc(doNothing, r.AFTER, &App{})
	    //r.InterceptFunc(checkUser, r.BEFORE, &AnotherController{})
}