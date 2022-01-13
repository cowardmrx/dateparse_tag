# dateparse_tag
dateparse time by struct tag

## intro&简介
WithTagName() // 自定义你想要使用的tag名称，默认为dateFormat  
WithDefaultTagValue() // 定义这个tag的默认值，默认为 default  
WithDefaultFormat() // 定义时间格式化样式，默认为 2006-01-02 15:04:05  
WithEmptyValue() // 定义一个空值返回，当指定结构体的指定字段为空值时，返回你想要的空值，默认为 `""`
## install&安装
```go
    go get github.com/cowardmrx/dateparse_tag
```

### attention&&注意
```go
Parse(in,out interface{}) // 该方法已经弃用，但是仍然可以使用，不建议使用
Format(in interface{}) // 该方法已经替代Parse()方法，支持嵌套结构体，in 参数请使用指针传递
```
## use&使用
```go
type User struct {
    Name     string `json:"name"`
    BirthDay string `json:"birth_day" format_date:"default"`
}

func TestNewDateParseTag(t *testing.T) {
    u := new(User)
    
    u.Name = "张三"
    u.BirthDay = time.Now().String()
    
    t.Logf("user : %v", u)
    
    dp := NewDateParseTag(WithTagName("format_date"))
    
    dp.Parse(u, u)
    
    t.Logf("user new %v", u)
}

// old user : &{张三 2022-01-12 14:10:17.1867047 +0800 CST m=+0.003444301}
// new user  &{张三 2022-01-12 14:10:17}


type User struct {
    Name     string `json:"name"`
    BirthDay string `json:"birth_day" format_date:"default"`
    DateS    DateS  `json:"date_s"`
}

type DateS struct {
    OldData  string    `json:"old_data" format_date:"default"`
    NewDates *NewDates `json:"new_dates"`
}

type NewDates struct {
    NewDates string `json:"new_dates" format_date:"default"`
}

func TestNewDateParseTag(t *testing.T) {
    u := new(User)
    
    u.Name = "张三"
    u.BirthDay = time.Now().String()
    
    t.Logf("user : %v", u)
    
    dp := NewDateParseTag(WithTagName("format_date"))
    
    dp.Parse(u, u)
    
    t.Logf("user new %v", u)
}

// user : &{张三 2022-01-13 16:21:36.2974741 +0800 CST m=+0.003138401 {2022-01-13 16:21:36.3125376 +0800 CST m=+0.018201901 0xc000088550}}
// user new &{张三 2022-01-13 16:21:36 {2022-01-13 16:21:36 0xc000088550}}  {2022-01-13 16:21:36 0xc000088550}  &{2022-01-13 16:21:36}



func TestDateParseTag_Format2(t *testing.T) {
    u := new(User)
    
    u.Name = "张三"
    u.BirthDay = time.Now().String()
    u.DateS = DateS{
    OldData: time.Now().String(),
    NewDates: &NewDates{
    NewDates: time.Now().String(),
    },
    }
    
    t.Logf("user : %v", u)
    
    dp := NewDateParseTag(WithTagName("format_date"))
    
    dp.Format(u)
    
    t.Logf("user new %v  %v  %v", u, u.DateS, u.DateS.NewDates)
}

// user : &{张三 2022-01-13 16:36:50.1955716 +0800 CST m=+0.004237901 {2022-01-13 16:36:50.2103154 +0800 CST m=+0.018981701 0xc000044570}}
// user new &{张三 2022-01-13 16:36:50 {2022-01-13 16:36:50 0xc000044570}}  {2022-01-13 16:36:50 0xc000044570}  &{2022-01-13 16:36:50}
```