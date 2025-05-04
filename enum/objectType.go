package enum

type ObjectType string

const (
	Blob ObjectType = "blob"
	Tree ObjectType = "tree"
	Commit ObjectType = "commit"
)

var objectTypes = map[string]ObjectType{
	"blob": Blob,
	"tree": Tree,
	"commit": Commit,
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
