package utils

import (
	"reflect"
	"strings"
)

type Tags struct {
	FieldName []string
	tags      []string
}

func NewTags(t interface{}, nameTag string) *Tags {
	tags, fieldNames := getArrayTags(t, nameTag)
	return &Tags{
		tags:      tags,
		FieldName: fieldNames,
	}
}

// GetTagsClean GetArrayTags return array of tags without empty values or lastname attributes
func (t *Tags) GetTagsClean() []string {
	var tagsClean []string
	for i := range t.tags {
		if t.tags[i] != "-" {
			tagsClean = append(tagsClean, strings.Split(t.tags[i], ",")[0])
		}
	}
	return tagsClean
}

// GetMapFieldsAndTag return each field with its tag
func (t *Tags) GetMapFieldsAndTag() map[string]string {
	var mapFieldsAndTag = make(map[string]string)
	for i := range t.tags {
		mapFieldsAndTag[t.FieldName[i]] = t.tags[i]
	}
	return mapFieldsAndTag
}

// GetMapFieldsAndTagCleans return each field with its tag
func (t *Tags) GetMapFieldsAndTagCleans() map[string]string {
	var mapFieldsAndTag = make(map[string]string)
	for i, v := range t.GetTagsClean() {
		mapFieldsAndTag[t.FieldName[i]] = v
	}
	return mapFieldsAndTag
}

func (t *Tags) GetTagByField(s string) string {
	for i, v := range t.FieldName {
		if v == s {
			return t.tags[i]
		}
	}

	panic("tags: no se encontro el campo " + s)
}

func getArrayTags(t interface{}, nameTag string) ([]string, []string) {

	field := reflect.TypeOf(t).NumField()

	var tags []string
	var fieldNames []string
	for i := 0; i < field; i++ {
		fieldName := reflect.TypeOf(t).Field(i).Name
		tag := reflect.TypeOf(t).Field(i).Tag.Get(nameTag)
		if tag != "" {
			tags = append(tags, tag)
		}
		if fieldName != "" {
			fieldNames = append(fieldNames, fieldName)
		}
	}

	return tags, fieldNames
}
