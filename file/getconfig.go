var Configobj map[string]map[string]interface{}

func init(){
	Configobj = make(map[string]map[string]interface{})
	f, err1 := os.OpenFile(path + "/config.json", os.O_RDONLY, 0666)

	if err1 != nil {
		utils.Log(err1)
	}

	err := json.NewDecoder(f).Decode(&Configobj)
	if err != nil {
		fmt.Println(err)
	}
}

//获取对象值
func GetConfig(db,key string) interface{} {

	if obj,ok := Configobj[db];ok{
		if k,o := obj[key];o{
			return k
		}else {
			panic("不存在")
		}
	}else{
		panic("不存在")
	}

	return ""
}
