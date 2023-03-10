package _const

const (
	INTEGER = "integer"
	FLOAT   = "float"
	DECIMAL = "decimal"
	BIT     = "bit"
	DATE    = "date"
	STRING  = "string"
	BIN     = "bin"
	BLOB    = "blob"
	TEXT    = "text"
)

type DataTypes struct {
	Types map[string][]string `yaml:"types"`
}
