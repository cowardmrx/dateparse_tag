# dateparse_tag
dateparse time by struct tag

## intro&简介
WithTagName() // 自定义你想要使用的tag名称，默认为dateFormat  
WithDefaultTagValue() // 定义这个tag的默认值，默认为 default  
WithDefaultFormat() // 定义时间格式化样式，默认为 2006-01-02 15:04:05  
WithEmptyValue() // 定义一个空值返回，当指定结构体的指定字段为空值时，返回你想要的空值，默认为 `""`  

注意： 暂不支持结构体嵌套操作
## install&安装
```go
    go get github.com/cowardmrx/dateparse_tag
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

```