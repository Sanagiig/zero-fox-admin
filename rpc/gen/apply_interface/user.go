package apply_interface

import "gorm.io/gen"

type Querier interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int) (gen.T, error) // GetByID query data by id and return it as *struct*

	// GetByRoles query data by roles and return it as *slice of pointer*
	//   (The below blank line is required to comment for the generated method)
	//
	// SELECT * FROM @@table WHERE role IN @rolesName
	GetByRoles(rolesName ...string) ([]*gen.T, error)

	// InsertValue insert value
	//
	// INSERT INTO @@table (name, age) VALUES (@name, @age)
	InsertValue(name string, age int) error
}
