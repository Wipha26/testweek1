// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// GendersColumns holds the columns for the "genders" table.
	GendersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gender", Type: field.TypeString},
	}
	// GendersTable holds the schema information for the "genders" table.
	GendersTable = &schema.Table{
		Name:        "genders",
		Columns:     GendersColumns,
		PrimaryKey:  []*schema.Column{GendersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "idemployee", Type: field.TypeString, Size: 6},
		{Name: "gender_users", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_genders_users",
				Columns: []*schema.Column{UsersColumns[3]},

				RefColumns: []*schema.Column{GendersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GendersTable,
		UsersTable,
	}
)

func init() {
	UsersTable.ForeignKeys[0].RefTable = GendersTable
}
