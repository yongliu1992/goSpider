package field

// FieldType 代表日志字段的类型。
type Type int

// 日志字段类型常量。
const (
	UnknownType Type = iota
	BoolType    Type = iota
	Int64Type   Type = iota
	Float64Type Type = iota
	StringType  Type = iota
	ObjectType  Type = iota
)

// Field 代表日志字段的接口
type Field interface {
	Name() string
	Type() Type
	Value() interface{}
}
