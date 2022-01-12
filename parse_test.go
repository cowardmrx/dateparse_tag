package dateparse_tag

import (
	"testing"
	"time"
)

type User struct {
	Name     string `json:"name"`
	BirthDay string `json:"birth_day" format_date:"default"`
}

func TestNewDateParseTag(t *testing.T) {
	u := new(User)

	u.Name = "张三"
	u.BirthDay = time.Now().String()

	t.Logf("old user : %v", u)

	dp := NewDateParseTag(WithTagName("format_date"))

	dp.Parse(u, u)

	t.Logf("new user  %v", u)
}
