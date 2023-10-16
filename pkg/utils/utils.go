package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

func SortFields(request interface{}) string {
	rt := reflect.TypeOf(request)
	rv := reflect.ValueOf(request)
	var fields []string

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i).Interface()
		if reflect.DeepEqual(value, reflect.Zero(field.Type).Interface()) {
			continue
		}
		tagName := field.Tag.Get("json")
		fields = append(fields, fmt.Sprintf("%s%v", tagName, value))
	}
	sort.Strings(fields)
	result := strings.Join(fields, "")

	return result
}

func GenerateHMAC(value, secret string) (string, error) {
	key := []byte(secret)
	message := []byte(value)

	hasher := hmac.New(sha256.New, key)
	hasher.Write(message)
	signature := hex.EncodeToString(hasher.Sum(nil))

	return signature, nil
}

func EncodeForm(request interface{}) string {
	form := url.Values{}

	v := reflect.ValueOf(request)
	t := reflect.TypeOf(request)
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		if field.Type.Kind() == reflect.Map && v.Field(i).Len() == 0 {
			continue
		}
		if reflect.DeepEqual(value, reflect.Zero(field.Type).Interface()) {
			continue
		}
		form.Set(field.Tag.Get("json"), fmt.Sprintf("%v", value))
	}

	return form.Encode()
}
