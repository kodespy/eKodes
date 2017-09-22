package controllers

import (
	m "github.com/kodespy/ekodes/app/models"
	r "github.com/revel/revel"
)


func getUsuario(_email string, _clave string) (m.Usuario, error) {
	//Retrieve
	_usuario := m.Usuario{}
	r.INFO.Println("--- Cargando usuarios")

	var id string
	var nombre string
	var clave string
	var email string
	err := db.QueryRow("SELECT id, nombre, email, clave FROM usuarios where email = $1 and clave = $2", _email, _clave).Scan(&id, &nombre, &email, &clave)
	if err == nil {
		_usuario = m.Usuario{Id: id, Nombre: nombre, Email: email, Clave: clave}
	}
	return _usuario, err
}