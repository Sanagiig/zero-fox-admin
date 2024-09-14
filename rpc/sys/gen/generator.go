// configuration.go
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./rpc/sys/gen/query", // output directory, default value is ./query
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	//var dsn = "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	var dsn = "root:simple-admin.@tcp(101.43.41.211:3306)/zero_fox_admin?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	// Initialize a *gorm.DB instance
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	g.ApplyBasic(
		// g.GenerateModel("sys_dept"),
		// g.GenerateModel("sys_dict"),
		// g.GenerateModel("sys_job"),
		// g.GenerateModel("sys_log"),
		// g.GenerateModel("sys_login_log"),
		// g.GenerateModel("sys_menu"),
		// g.GenerateModel("sys_role"),
		// g.GenerateModel("sys_role_dept"),
		// g.GenerateModel("sys_role_menu"),
		g.GenerateModel("sys_user"),
		// g.GenerateModel("sys_user_role"),
	)

	// Execute the generator
	g.Execute()
}
