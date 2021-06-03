### null
要保证获取的字段必须都是存在的，如果当前的字段内容不存在则需要设置相关的默认值，
 PostgreSQL 的方式 COALESCE(user_data.dnpu,0)
 mysql ifnull(user_data.dnpu,0) 当前的字段

数据库字段为null的处理
```
	switch c.ScanType().Name() {
		case "string":
			scanArgs[i] = &sql.NullString{String: "", Valid: true}
		case "int64":
			scanArgs[i] = &sql.NullInt64{Int64: 0, Valid: true}
		case "int32":
			scanArgs[i] = &sql.NullInt32{Int32: 0, Valid: false}
		case "float64":
			scanArgs[i] = &sql.NullFloat64{Valid: false}
		case "bool":
			scanArgs[i] = &sql.NullBool{Valid: false}
		case "datetime":
			scanArgs[i] = &sql.NullTime{Valid: false}

		default:
			scanArgs[i] = reflect.New(c.ScanType()).Interface()
		}
```
