package controllers

import (
//	"fmt"
//	"time"
//	"github.com/lib/pq"
	m "github.com/kodespy/ekodes/app/models"
	r "github.com/revel/revel"
)


func allCiudad() ([]m.Ciudad, error) {
	//Retrieve
	ciudades := []m.Ciudad{}
	r.INFO.Println("--- Cargando ciudades")
	rows, err := db.Query("SELECT id, nombre FROM ciudades order by nombre")
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var id string
			var nombre string
			err = rows.Scan(&id, &nombre)
			if err == nil {
				_Ciudad := m.Ciudad{Id: id, Nombre: nombre}
				ciudades = append(ciudades, _Ciudad)
			} else {
				return ciudades, err
			}
		}
	} else {
		return ciudades, err
	}
	r.INFO.Println(ciudades)
	return ciudades, err
}
