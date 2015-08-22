package main

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mholt/binding"
)

// DBActions struct for database actions - create, drop, migrate
type DBActions struct {
	db *gorm.DB
}

// Model provides a default model struct, you could embed it in your struct
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Stubo is a struct for keeping information about single stubo instance
type Stubo struct {
	gorm.Model
	Name     string
	Version  string
	Hostname string
	Port     string
	Protocol string
	// Clusters []Cluster `gorm:"many2many:stubo_clusters;"` // Many-To-Many relationship, 'stubo_clusters' is join table
}

// StuboForm form values we need for updating/creating a stubo.
type StuboForm struct {
	Name     string
	Version  string
	Hostname string
	Port     string
	Protocol string
}

// FieldMap - to do some validation on our input fields. File is done separately.
func (sf *StuboForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&sf.Name: binding.Field{
			Form:     "name",
			Required: true,
		},
		&sf.Version: binding.Field{
			Form:     "version",
			Required: false,
		},
		&sf.Hostname: binding.Field{
			Form:     "hostname",
			Required: true,
		},
		&sf.Port: binding.Field{
			Form:     "port",
			Required: true,
		},
		&sf.Protocol: binding.Field{
			Form:     "protocol",
			Required: true,
		},
	}
}

// Cluster lets users to group stubo instances
type Cluster struct {
	gorm.Model
	Name string `sql:"index:idx_name_code"`
	Code string
}

func (d DBActions) createTables() {
	// creating Stubo table
	d.db.CreateTable(&Stubo{})
	d.db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Stubo{})
	// creating Cluster table
	d.db.CreateTable(&Cluster{})
	d.db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Cluster{})
}

func (d DBActions) dropTables() {
	d.db.DropTable(&Stubo{})
	d.db.DropTable(&Cluster{})
}
