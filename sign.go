package luksdk

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// signature 签名
func signature(signSecret string, params any) string {
	paramsMap := castToSignatureParams(params)
	return generateSignature(signSecret, paramsMap)
}

// generateSignature 生成签名
func generateSignature(signSecret string, params map[string]string) string {
	// 提取并排序参数键
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构建签名字符串
	var signatureBuilder strings.Builder
	for _, k := range keys {
		if value := params[k]; value != "" {
			signatureBuilder.WriteString(k)
			signatureBuilder.WriteString("=")
			signatureBuilder.WriteString(value)
			signatureBuilder.WriteString("&")
		}
	}

	// 删除末尾的 '&' 并拼接密钥
	signatureString := strings.TrimSuffix(signatureBuilder.String(), "&") + fmt.Sprintf("&key=%s", signSecret)

	// 生成 MD5 哈希并将结果转换为大写
	hash := md5.Sum([]byte(signatureString))
	signature := strings.ToUpper(hex.EncodeToString(hash[:]))

	return signature
}

// castToSignatureParams 将结构体转换为 map[string]string
func castToSignatureParams(obj any) map[string]string {
	result := make(map[string]string)
	value := reflect.ValueOf(obj)

	// 处理指针类型，获取其指向的元素
	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	if value.Kind() == reflect.Map {
		for _, key := range value.MapKeys() {
			result[key.String()] = fmt.Sprint(value.MapIndex(key).Interface())
		}
	} else {
		// 遍历结构体字段
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			switch field.Kind() {
			case reflect.Map, reflect.Pointer, reflect.Slice, reflect.Chan, reflect.Func, reflect.Interface, reflect.Array, reflect.Struct, reflect.UnsafePointer, reflect.Complex128, reflect.Complex64:
				continue
			default:
				jsonTag := value.Type().Field(i).Tag.Get("json")
				tagName := strings.Split(jsonTag, ",")[0]

				// 跳过 "sign" 字段和空字段
				if tagName != "sign" && !field.IsZero() {
					result[tagName] = fmt.Sprint(field.Interface())
				}
			}
		}
	}

	return result
}
