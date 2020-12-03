string 使用的8比特字节的卷积核来存储字符，存储的字符是utf-8编码的的格式
例如每个汉字字符的UTF-8编码，占用多少字节
在使用for each 遍历字符串的时候，每次迭代返回utf-8编码的首个字节的下标及字节值，这就意味着下标可能是不连续
func destr (){
   s := "天道"
   for index,value := range s {
      fmt.Printf("index:%d ,value: %c\n ",index,value)
   }
}

分别输出的是 index 0  天
            index 3  道    字符串的长度是字节的长度，而非字符的长度  汉字占3个字节，字符串天道的长度是 6 
            字符串可以为空 但是不能为nil ,可接受新的字符串，赋值，但是不能通过下标的方式，更改字符串的信息   
            
            字符串的 数据类型是 struct { 
                            str unsafe.Pointer
                            len int 
                            }  字符串的存放的地址，和字符串存放的长度
