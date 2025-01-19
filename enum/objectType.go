package enum

type ObjectType string

const (
	Blob ObjectType = "blob"
	Tree ObjectType = "tree"
)

var objectTypes = map[string]ObjectType{
	"blob": Blob,
	"tree": Tree,
}

func GetObjectType(typo string) (ObjectType, bool) {
	if _, ok := objectTypes[typo]; ok {
		return objectTypes[typo], true
	}
	return "", false
}

func (o ObjectType) GetObjectType() string {
	return string(o)
}
