//查询当前的操作的方式
func mlldat() {

	f, err := os.Create("test.xls")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	w.Write([]string{"编号", "idfa", "登录次数"})

	mdb := GetMongoSession().Copy()
	defer mdb.Clone()

	mongo := GetMongoSession().Copy()
	fmt.Println("23")
	//	var totalMoney int //the cast is fen
	defer mongo.Close()

	type rtype struct {
		Id    string `bson:"_id"`
		Total int    `bson:"total"`
	}
	result := rtype{}
	fmt.Println("asddasasd-----")
	dat := mongo.DB("channel").C("1474_result").Find(nil).Iter()
	var num int = 1
	for dat.Next(&result) {
		oyt, _ := url.QueryUnescape(result.Id)
		fmt.Println(oyt)
		w.Write([]string{strconv.Itoa(num), oyt, strconv.Itoa(result.Total)})
		num += 1
	}
	w.Flush()
}




