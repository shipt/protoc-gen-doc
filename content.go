package gendoc

import (
	"strings"
)

type LinkedMessageField struct {
	Self     *MessageField
	Children []*LinkedMessageField
}

func Print(f *LinkedMessageField) {
	println(f.Self.Name)
	if f.Children != nil {
		for _, child := range f.Children {
			Print(child)
		}
	}
}

// GetContent parses file data and returns tree structured list of LinkedMessageField
func GetContent(files []*File, baseMessage string) []*LinkedMessageField {
	var linkedFields []*LinkedMessageField
	for _, file := range files {
		for _, message := range file.Messages {
			if strings.ToLower(message.LongName) == strings.ToLower(baseMessage) {
				for _, field := range message.Fields {
					linkedField := &LinkedMessageField{Self: field}
					if !IsScalarType(field.LongType) {
						getChildField(files, linkedField)
					}
					linkedFields = append(linkedFields, linkedField)
				}
			}
		}
	}
	return linkedFields
}

func getChildField(files []*File, parentLinkedField *LinkedMessageField) {
	for _, file := range files {
		for _, message := range file.Messages {
			if message.FullName == parentLinkedField.Self.FullType {
				for _, field := range message.Fields {
					linkedField := &LinkedMessageField{Self: field}
					if !IsScalarType(field.LongType) {
						getChildField(files, linkedField)
					}
					if parentLinkedField.Children == nil {
						parentLinkedField.Children = []*LinkedMessageField{}
					}
					parentLinkedField.Children = append(parentLinkedField.Children, linkedField)
				}
			}
		}
	}
}

func IsScalarType(fieldType string) bool {
	switch strings.ToLower(fieldType) {
	case
		"double",
		"float",
		"int32",
		"int64",
		"uint32",
		"uint64",
		"sint32",
		"sint64",
		"fixed32",
		"sfixed32",
		"sfixed64",
		"bool",
		"string",
		"bytes":
		return true
	}
	return false
}
