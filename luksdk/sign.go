package luksdk

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// Signature 生成签名
func Signature(signSecret string, params any) string {
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
		// 重要修复：正确处理 map[string]interface{}
		for _, key := range value.MapKeys() {
			keyStr := fmt.Sprint(key.Interface())
			// 跳过 sign 字段
			if keyStr == "sign" {
				continue
			}
			
			// 获取值并转换为字符串
			val := value.MapIndex(key)
			valStr := fmt.Sprint(val.Interface())
			
			// 跳过空值和nil
			if valStr == "" || valStr == "<nil>" {
				continue
			}
			
			// 跳过0值（但保留c_id和timestamp即使是0）
			if keyStr != "c_id" && keyStr != "timestamp" && valStr == "0" {
				continue
			}
			
			result[keyStr] = valStr
		}
	} else if value.Kind() == reflect.Struct {
		// 遍历结构体字段
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			fieldType := value.Type().Field(i)
			jsonTag := fieldType.Tag.Get("json")
			tagName := strings.Split(jsonTag, ",")[0]
			
			// 跳过 sign 字段
			if tagName == "sign" {
				continue
			}
			
			// 跳过复杂类型
			switch field.Kind() {
			case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct, 
			     reflect.Func, reflect.Chan, reflect.UnsafePointer, 
			     reflect.Complex64, reflect.Complex128:
				continue
			case reflect.Pointer, reflect.Interface:
				// 处理指针和interface
				if field.IsNil() {
					continue
				}
				elem := field.Elem()
				if !elem.IsValid() {
					continue
				}
				valStr := fmt.Sprint(elem.Interface())
				if valStr != "" && valStr != "0" {
					result[tagName] = valStr
				}
			default:
				// 处理基本类型
				if !field.IsZero() {
					result[tagName] = fmt.Sprint(field.Interface())
				}
			}
		}
	}

	return result
}
