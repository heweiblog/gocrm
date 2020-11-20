package utils

import (
	"encoding/json"
	"fmt"
	"gocrm/models"
)

type BaseMap map[string]map[string]map[string]func(c *models.Content) string

var CheckMethods BaseMap

type ResponseCode struct {
	Responsecode int `json:"responsecode"`
}

func ResponseCodeCheck(c *models.Content) string {
	m := make(map[string]interface{})
	if err := json.Unmarshal(c.Data, &m); err != nil {
		return err.Error()
	}
	if res, ok := m["responsecode"]; ok {
		if fcode, ok := res.(float64); ok {
			code := int(fcode)
			if code < 0 || code > 0xffffffff {
				return "responsecode range error"
			}
			models.DB.Create(c)
			return ""
		}
	}
	return "responsecode not in data"
}

func init() {
	CheckMethods = make(BaseMap)
	fmt.Println(CheckMethods)
	m := make(map[string]func(c *models.Content) string)
	m["responserules"] = ResponseCodeCheck
	n := make(map[string]map[string]func(c *models.Content) string)
	n["selfcheck"] = m
	CheckMethods["handle"] = n
	fmt.Println(CheckMethods)
}