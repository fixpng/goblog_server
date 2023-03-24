package utils

// InList 判断key是否存在于列表中
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}
