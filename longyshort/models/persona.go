package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Persona struct {
	ID     int    `json:"id" db:"id"`
	Edad   int    `json:"edad" db:"edad"`
	Nombre string `json:"nombre" db:"nombre"`
	Sexo   bool   `json:"sexo" db:"sexo"`
	Genero string `json:"genero" db:"genero"`
}

var db *sqlx.DB

func InitDB() {
	var err error
	db, err = sqlx.Connect("mysql", "root:adrian0710200512#12#@tcp(localhost:3306)/mi_base_de_datos")
	if err != nil {
		log.Fatal(err)
	}
}

// Crear una nueva persona
func CreatePersona(persona *Persona) error {
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	_, err := db.Exec("INSERT INTO persona (edad, nombre, sexo, genero) VALUES (?, ?, ?, ?)", persona.Edad, persona.Nombre, persona.Sexo, persona.Genero)
	return err
}

func GetAllPersonas() ([]Persona, error) {
	var personas []Persona
	err := db.Select(&personas, "SELECT id, edad, nombre, sexo, genero FROM persona")
	if err != nil {
		return nil, fmt.Errorf("Error al obtener las personas: %v", err)
	}
	return personas, nil
}

func GetRecentPersonas(limit int) ([]Persona, error) {
	var personas []Persona
	query := fmt.Sprintf("SELECT id, edad, nombre, sexo, genero FROM persona ORDER BY id DESC LIMIT %d", limit)
	err := db.Select(&personas, query)
	if err != nil {
		return nil, fmt.Errorf("Error al obtener las personas recientes: %v", err)
	}
	return personas, nil
}

func GetGenderCount() (map[string]int, error) {
	count := make(map[string]int)

	// Contar los hombres
	var maleCount int
	err := db.Get(&maleCount, "SELECT COUNT(*) FROM persona WHERE sexo = 1")
	if err != nil {
		return nil, fmt.Errorf("Error al contar hombres: %v", err)
	}
	count["hombres"] = maleCount

	// Contar las mujeres
	var femaleCount int
	err = db.Get(&femaleCount, "SELECT COUNT(*) FROM persona WHERE sexo = 0")
	if err != nil {
		return nil, fmt.Errorf("Error al contar mujeres: %v", err)
	}
	count["mujeres"] = femaleCount

	return count, nil
}
