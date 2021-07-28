// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIdemployee holds the string denoting the idemployee field in the database.
	FieldIdemployee = "idemployee"

	// EdgeGender holds the string denoting the gender edge name in mutations.
	EdgeGender = "gender"

	// Table holds the table name of the user in the database.
	Table = "users"
	// GenderTable is the table the holds the gender relation/edge.
	GenderTable = "users"
	// GenderInverseTable is the table name for the Gender entity.
	// It exists in this package in order to avoid circular dependency with the "gender" package.
	GenderInverseTable = "genders"
	// GenderColumn is the table column denoting the gender relation/edge.
	GenderColumn = "gender_users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldIdemployee,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the User type.
var ForeignKeys = []string{
	"gender_users",
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// IdemployeeValidator is a validator for the "idemployee" field. It is called by the builders before save.
	IdemployeeValidator func(string) error
)
