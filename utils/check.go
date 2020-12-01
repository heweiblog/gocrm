package utils

import (
	"encoding/json"
	"fmt"
	"gocrm/models"
)

const (
	DnsService = "dns"
	ZrpService = "handle"
	MsSource   = "ms"
	CliSource  = "cli"
)

type BaseMap map[string]func(c *models.Content) string

var (
	OpMap        map[string]bool
	CheckMethods BaseMap
	MsRelease    string
)

func init() {
	OpMap = make(map[string]bool)
	OpMap["add"] = true
	OpMap["delete"] = true
	OpMap["update"] = true
	OpMap["clear"] = true
	OpMap["query"] = true
}

func CheckData(release string, contents []models.Content) string {
	//记录 ms release
	MsRelease = release
	for i := 0; i < len(contents); i++ {
		if contents[i].Service != DnsService && contents[i].Service != ZrpService {
			return "service value error"
		}
		if contents[i].Source != MsSource && contents[i].Source != CliSource {
			return "source value error"
		}
		if contents[i].Bt == "" {
			return "bt value error"
		}
		if contents[i].Sbt == "" {
			return "sbt value error"
		}
		if contents[i].Mid < 0 {
			return "id value error"
		}
		if _, ok := OpMap[contents[i].Op]; ok == false {
			return "op value error"
		}
		key := contents[i].Service + contents[i].Bt + contents[i].Sbt
		if f, ok := CheckMethods[key]; ok {
			//校验模块
			if res := f(&contents[i]); res == "" {
				//校验通过 入队(通道) 直接启动协程
				//go testPost(d)
				fmt.Println("数据校验通过")
			} else {
				//校验失败，直接记录oplog，或者发到通道，在另一个协程中收取写入
				fmt.Println("数据校验失败", res)
			}
		} else {
			//service bt sbt或对应的函数有问题，直接记录oplog失败，或者发到通道，在另一个协程中收取写入
			fmt.Println("Unsupported configuration")
		}
	}
	return ""
}

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
