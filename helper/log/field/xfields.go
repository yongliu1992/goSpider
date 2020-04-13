package field

// 拥有bool类型的值的日志字段。
type boolField struct {
	name      string
	fieldType Type
	value     bool
}

func (field *boolField) Name() string {
	return field.name
}

func (field *boolField) Type() Type {
	return field.fieldType
}
func (field *boolField) Value() interface{} {
	return field.value
}

func Bool(name string, value bool) Field {
	return &boolField{name: name, fieldType: BoolType, value: value}
}

// 拥有int64类型的值的日志字段。
type int64Field struct {
	name      string
	fieldType Type
	value     int64
}

func (field *int64Field) Name() string {
	return field.name
}

func (field *int64Field) Type() Type {
	return field.fieldType
}
func (field *int64Field) Value() interface{} {
	return field.value
}

func Int64(name string, value int64) Field {
	return &int64Field{name: name, fieldType: Int64Type, value: value}
}

// 拥有float64类型的值的日志字段。
type float64Field struct {
	name      string
	fieldType Type
	value     float64
}

func (field *float64Field) Name() string {
	return field.name
}

func (field *float64Field) Type() Type {
	return field.fieldType
}
func (field *float64Field) Value() interface{} {
	return field.value
}

func Float64(name string, value float64) Field {
	return &float64Field{name: name, fieldType: Float64Type, value: value}
}

// 拥有string类型的值的日志字段。
type stringField struct {
	name      string
	fieldType Type
	value     string
}

func (field *stringField) Name() string {
	return field.name
}

func (field *stringField) Type() Type {
	return field.fieldType
}
func (field *stringField) Value() interface{} {
	return field.value
}

func String(name string, value string) Field {
	return &stringField{name: name, fieldType: StringType, value: value}
}

// 拥有interface{}类型的值的日志字段。
type objectField struct {
	name      string
	fieldType Type
	value     interface{}
}

func (field *objectField) Name() string {
	return field.name
}

func (field *objectField) Type() Type {
	return field.fieldType
}
func (field *objectField) Value() interface{} {
	return field.value
}

func Object(name string, value interface{}) Field {
	return &objectField{name: name, fieldType: ObjectType, value: value}
}
