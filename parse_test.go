package dateparse_tag

import (
	"testing"
	"time"
)

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

func TestDateParseTag_Format(t *testing.T) {
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
