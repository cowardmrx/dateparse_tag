package dateparse_tag

import (
	"fmt"
	"github.com/araddon/dateparse"
)

//	@method validateDateFormat
//	@description: validate date layout
//	@param layout string
//	@return string
//	@return error
func validateDateFormat(layout string) (string, error) {
	tFormat, err := dateparse.ParseFormat(layout)

	if err != nil {
		return "", err
	}

	return tFormat, nil
}

//	@method parseTime
//	@description: parse time
//	@param layout string
//	@param times string
//	@return string
func parseTime(layout, times string) string {
	t1, err := dateparse.ParseLocal(times)

	if err != nil {
		panic(fmt.Sprintf("time parse failed: %v : %v", times, err.Error()))
	}

	return t1.Format(layout)
}
