package enum

type ObjectType string

const (
	Blob ObjectType = "blob"
	Tree ObjectType = "tree"
)

func GetObjectType(typo string) ObjectType {
	return map[string]ObjectType{
		"blob": Blob,
		"tree": Tree,
	}[typo]
}

func (o ObjectType) GetObjectType() string {
	return string(o)
}
