package tools

import (
	"regexp"
)

var (
	//reg_chinese_phone_number = `[^\w]((?:(?:\+|0{0,2})86)?1(?:(?:3[\d])|(?:4[5-79])|(?:5[0-35-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\d])|(?:9[189]))\d{8})[^\w]`
	reg_chinese_phone_number = `(\b((?:\+|0{0,2})86)?1(?:3\d|4[5-79]|5[0-35-9]|6[5-7]|7[0-8]|8\d|9[189])\d{8}\b)`
	//reg_chinese_phone_number = `[^0-9]((\+86|86)?1[3-9]\d{9})[^0-9]`
	reg_chinese_id_card = `\d{6}(18|19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\d|3[01])\d{3}[\dXx]`
	//reg_url                  = `(\b(?![\w]{0,10}?https?://)(([-A-Za-z0-9]{1,20})://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|])`
	reg_url = `^(https?|ftp):\/\/(?:([^:@\/\s]+):([^@\/\s]*)@)?([^:\/?#\s]+)(?::(\d+))?([^?#\s]*)(?:\?([^#\s]*))?(?:#([^\s]*))?$`
	reg_ip  = `[^0-9]((127\.0\.0\.1)|(10\.\d{1,3}\.\d{1,3}\.\d{1,3})|(172\.((1[6-9])|(2\d)|(3[01]))\.\d{1,3}\.\d{1,3})|(192\.168\.\d{1,3}\.\d{1,3}))`
)

// 匹配电话号码
func Reg_phone(text []string) []string {

	var result []string

	for _, i := range text {
		re := regexp.MustCompile(reg_chinese_phone_number)
		matchs := re.FindAllString(i, -1)
		if len(matchs) > 0 {
			for _, value := range matchs {
				result = append(result, value)
			}
		}
	}

	return removeEmptyStrings(result)

}

func Reg_IP(text []string) []string {
	var result []string

	for _, i := range text {
		re := regexp.MustCompile(reg_ip)
		matchs := re.FindAllString(i, -1)
		if len(matchs) > 0 {
			for _, value := range matchs {
				result = append(result, value)
			}
		}
	}
	return removeEmptyStrings(result)
}

func Reg_id_card(text []string) []string {

	var result []string

	for _, i := range text {
		re := regexp.MustCompile(reg_chinese_id_card)
		matchs := re.FindAllString(i, -1)
		if len(matchs) > 0 {
			for _, value := range matchs {
				if ValidateChecksum(value) {
					result = append(result, value)
				}

			}
		}
	}
	return removeEmptyStrings(result)

}

func Reg_url(text []string) []string {

	var result []string

	for _, i := range text {
		re := regexp.MustCompile(reg_url)
		matchs := re.FindAllString(i, -1)
		if len(matchs) > 0 {
			for _, value := range matchs {
				result = append(result, value)
			}
		}
	}
	return removeEmptyStrings(result)

}

// 返回一个map结果，包含电话号码、身份证号码、url、ip
func Reg_all(text []string) map[string][]string {
	var result = make(map[string][]string)
	result["phone"] = Reg_phone(text)
	result["id_card"] = Reg_id_card(text)
	result["url"] = Reg_url(text)
	result["ip"] = Reg_IP(text)
	return result
}
func removeEmptyStrings(slice []string) []string {
	// 创建一个新的切片，用于存储非空字符串
	result := make([]string, 0, len(slice))

	// 遍历原始切片，过滤掉空字符串
	for _, s := range slice {
		if s != "" {
			result = append(result, s)
		}
	}

	return result
}

// ValidateChecksum 验证身份证号码的校验码
func ValidateChecksum(id string) bool {
	if len(id) != 18 {
		return false
	}

	// 权重系数
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 校验码对应表
	codes := []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

	// 计算加权和
	sum := 0
	for i := 0; i < 17; i++ {
		sum += int(id[i]-'0') * weights[i]
	}

	// 计算校验码
	mod := sum % 11
	expectedCode := codes[mod]

	// 比较校验码
	actualCode := id[17]
	if actualCode == 'x' {
		actualCode = 'X'
	}

	return actualCode == expectedCode
}
