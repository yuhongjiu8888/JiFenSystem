package Map

/*
获取map中对应key的默认value值
*/

func GetMapValue(temp map[string]interface{}, key string, value interface{})(interface{}){

	_ , ok := temp[key]
	if !ok{
		return value
	}

	return temp[key]
}