package dateparse_tag

import (
	"fmt"
	"reflect"
)

const (
	DefaultFormatLayout = "2006-01-02 15:04:05"
	DefaultTagName      = "dateFormat"
	DefaultTagValue     = "default"
)

type dateParseTag struct {
	TagName             string // Define the tag you want , default tag is `dateFormat`
	DefaultTagValue     string // define the tag default value you want , default tag value `default`
	DefaultFormatLayout string // date format default layout
	EmptyValue          string // Return when the value is null
}

type Options func(dp *dateParseTag)

func WithTagName(tagName string) Options {
	return func(dp *dateParseTag) {

		if tagName == "" || len(tagName) <= 0 {
			dp.TagName = "dateFormat"
		} else {
			dp.TagName = tagName
		}

	}
}

func WithDefaultTagValue(tagVal string) Options {
	return func(dp *dateParseTag) {
		if tagVal == "" || len(tagVal) <= 0 {
			dp.DefaultTagValue = "default"
		} else {
			dp.DefaultTagValue = tagVal
		}
	}
}

func WithDefaultFormat(defaultFormatLayout string) Options {
	return func(dp *dateParseTag) {
		if defaultFormatLayout == "" || len(defaultFormatLayout) <= 0 {
			dp.DefaultFormatLayout = DefaultFormatLayout
		} else {
			dp.DefaultFormatLayout = defaultFormatLayout
		}
	}
}

func WithEmptyValue(emptyValue string) Options {
	return func(dp *dateParseTag) {
		dp.EmptyValue = emptyValue
	}
}

type DateParseTag interface {
	//Deprecated
	Parse(in, out interface{})
	// Format 格式化
	Format(in interface{})

	// ForRange for range data
	forRange(field reflect.StructField, val reflect.Value)
}

func NewDateParseTag(opts ...Options) DateParseTag {
	dp := new(dateParseTag)

	for _, v := range opts {
		v(dp)
	}

	dp = dp.check()

	return dp
}

func (dp *dateParseTag) check() *dateParseTag {
	if dp.TagName == "" || len(dp.TagName) <= 0 {
		dp.TagName = DefaultTagName
	}

	if dp.DefaultFormatLayout == "" || len(dp.DefaultFormatLayout) <= 0 {
		dp.DefaultFormatLayout = DefaultFormatLayout
	}

	if dp.EmptyValue == "" || len(dp.EmptyValue) <= 0 {
		dp.EmptyValue = ""
	}

	return dp
}

//	@method Parse Deprecated
//	@description: parse time
//	@receiver dp
//	@param in interface{}
//	@param out interface{}
func (dp *dateParseTag) Parse(in, out interface{}) {
	tType := reflect.TypeOf(in).Elem()

	tValue := reflect.ValueOf(in).Elem()

	for i := 0; i < tType.NumField(); i++ {
		field := tType.Field(i)

		tag, ok := field.Tag.Lookup(dp.TagName)

		if !ok {
			continue
		}

		fieldVal := tValue.FieldByName(field.Name)

		if fieldVal.String() == "" || len(fieldVal.String()) <= 0 {

			// 校验是否为该空值返回是否为时间类型格式
			tFormat, err := validateDateFormat(dp.EmptyValue)

			// 如果不是时间类型格式那么直接返回emptyValue的值 反之按照layout格式化返回
			if err != nil {
				fieldVal.SetString(dp.EmptyValue)
			} else {
				fieldVal.SetString(parseTime(dp.dateFormatLayout(tFormat), dp.EmptyValue))
			}

		} else {

			fieldVal.SetString(parseTime(dp.dateFormatLayout(tag), fieldVal.String()))
		}

	}

	out = in
	return
}

//	@method Format
//	@description: parse time
//	@receiver dp
//	@param in interface{}
func (dp *dateParseTag) Format(in interface{}) {
	rt := reflect.TypeOf(in).Elem()

	rv := reflect.ValueOf(in).Elem()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)

		val := rv.Field(i)

		dp.forRange(field, val)
	}
}

//	@method forRange
//	@description: for range data
//	@param field reflect.StructField
//	@param val reflect.Value
func (dp *dateParseTag) forRange(field reflect.StructField, val reflect.Value) {
	switch field.Type.Kind() {
	case reflect.Ptr:
		rt := field.Type.Elem()
		rv := val.Elem()
		for k := 0; k < rt.NumField(); k++ {

			field := rt.Field(k)

			val := rv.Field(k)
			dp.forRange(field, val)

		}
	case reflect.Struct:
		rt := field.Type
		rv := val
		for k := 0; k < rt.NumField(); k++ {

			field := rt.Field(k)

			val := rv.Field(k)
			dp.forRange(field, val)

		}
	case reflect.String:
		tag, ok := field.Tag.Lookup(dp.TagName)

		if !ok {
			return
		}

		fieldVal := val

		if fieldVal.String() == "" || len(fieldVal.String()) <= 0 {

			// 校验是否为该空值返回是否为时间类型格式
			tFormat, err := validateDateFormat(dp.EmptyValue)

			// 如果不是时间类型格式那么直接返回emptyValue的值 反之按照layout格式化返回
			if err != nil {
				fieldVal.SetString(dp.EmptyValue)
			} else {
				fieldVal.SetString(parseTime(dp.dateFormatLayout(tFormat), dp.EmptyValue))
			}

		} else {

			fieldVal.SetString(parseTime(dp.dateFormatLayout(tag), fieldVal.String()))
		}
	}
}

//	@method dateFormatLayout
//	@description: get date format layout
//	@receiver dp
//	@param tag string
//	@return string
func (dp *dateParseTag) dateFormatLayout(tag string) string {
	if tag == dp.DefaultTagValue {
		tag = dp.DefaultFormatLayout
	} else {

		if tag == DefaultTagValue {
			tag = DefaultFormatLayout
		}

		tFormat, err := validateDateFormat(tag)

		if err != nil {
			panic(fmt.Sprintf("tag value can't parse: %v", tag))
		}

		tag = tFormat
	}

	return tag
}
