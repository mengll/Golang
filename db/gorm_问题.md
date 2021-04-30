### null
要保证获取的字段必须都是存在的，如果当前的字段内容不存在则需要设置相关的默认值，
 PostgreSQL 的方式 COALESCE(user_data.dnpu,0)
 mysql ifnull(user_data.dnpu,0) 当前的字段
