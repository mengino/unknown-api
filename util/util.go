package util

var order = map[string]string{
	"ascend":  "asc",
	"descend": "desc",
}

// BuilderOrder sql构造器排序
func BuilderOrder(field, sort string) string {
	return field + " " + order[sort]
}
