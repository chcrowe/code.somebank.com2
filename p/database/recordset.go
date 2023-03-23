package database

import (
	"fmt"
	"time"
)

func GetValue(pval *interface{}) string {
	switch v := (*pval).(type) {

	case nil:
		return ""

	case bool:
		if v {
			return "1"
		} else {
			return "0"
		}
	case []byte:
		if len(v) == 16 {
			return fmt.Sprintf("%X-%X-%X-%X-%X", v[0:4], v[4:6], v[6:8], v[8:10], v[10:16])
		} else {
			return string(v)
		}

	case time.Time:
		return v.Format("2006-01-02 15:04:05.999")

	default:
		return fmt.Sprintf("%v", v)
	}
}
