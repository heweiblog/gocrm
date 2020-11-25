package utils

import (
	"encoding/json"
	"gocrm/models"
)

type BaseMap map[string]func(c *models.Content) string

var CheckMethods BaseMap

/*
type ResponseCode struct {
	Responsecode int `json:"responsecode"`
}

func ResponseCodeCheck(c *models.Content) string {
	r := ResponseCode{-0x7f7f}
	if err := json.Unmarshal(c.Data, &r); err != nil {
		return err.Error()
	}
	fmt.Println(reflect.TypeOf(r))
	if r.Responsecode == -0x7f7f {
		return "responsecode not in data"
	}
	if r.Responsecode < 0 || r.Responsecode > 0xffffffff {
		return "responsecode range error"
	}
	//models.DB.Create(c)
	return ""
}
*/

func SwitchCheck(c *models.Content) string {
	m := make(map[string]interface{})
	if err := json.Unmarshal(c.Data, &m); err != nil {
		return err.Error()
	}
	if res, ok := m["switch"]; ok {
		if res, ok = res.(string); ok {
			if res == "enable" || res == "disable" {
				return ""
			}
		}
	}
	return "switch format error"
}

func ResponseCodeCheck(c *models.Content) string {
	m := make(map[string]interface{})
	if err := json.Unmarshal(c.Data, &m); err != nil {
		return err.Error()
	}
	if res, ok := m["responsecode"]; ok {
		if rcode, ok := res.(float64); ok {
			code := int(rcode)
			if code > 0 && code <= 0xffffffff {
				return ""
			}
			return "responsecode range error"
			//models.DB.Create(c)
		}
	}
	return "responsecode format error"
}

func init() {
	CheckMethods = make(BaseMap)
	CheckMethods["handle"+"selfcheck"+"responserules"] = ResponseCodeCheck
	CheckMethods["handle"+"selfcheck"+"switch"] = SwitchCheck
}
