// configuration.go
package main

import (
	"strings"
	"zero-fox-admin/rpc/gen/apply_interface"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	config := gen.Config{
		OutPath:       "./query", // output directory, default value is ./query
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	}

	config.WithJSONTagNameStrategy(func(columnName string) (tagContent string) {
		sb := strings.Builder{}
		parts := strings.Split(columnName, "_")
		if len(parts) == 1 {
			sb.WriteString(parts[0])
		} else {
			for i, part := range parts {
				if i == 0 {
					sb.WriteString(part)
				} else {

					sb.WriteString(strings.ToUpper(string(part[0])))
					sb.WriteString(part[1:])
				}
			}
		}

		return sb.String()
	})
	// Initialize the generator with configuration
	g := gen.NewGenerator(config)

	//var dsn = "root:123456@tcp(127.0.0.1:3306)/gozero?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	var dsn = "root:lwj@1993@tcp(localhost:3306)/zero-admin?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(mysql.Open(dsn))

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)
	g.ApplyInterface(func(apply_interface.Querier) {}, g.GenerateModel("sys_user"))
	g.ApplyBasic(
		g.GenerateModel("sys_dept"),
		g.GenerateModel("sys_dict"),
		g.GenerateModel("sys_dict_item"),
		g.GenerateModel("sys_job"),
		g.GenerateModel("sys_operate_log"),
		g.GenerateModel("sys_login_log"),
		g.GenerateModel("sys_menu"),
		g.GenerateModel("sys_role"),
		g.GenerateModel("sys_role_dept"),
		g.GenerateModel("sys_role_menu"),
		g.GenerateModel("sys_user_role"),
	)

	// Execute the generator
	g.Execute()
}
