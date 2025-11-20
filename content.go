package gendoc

import (
	"fmt"
	"strings"
)

// LinkedMessageField - Tree type object the maps graph between fields of proto messages
type LinkedMessageField struct {
	Self     *MessageField
	Children []*LinkedMessageField
}

func printLinkedMessageField(f *LinkedMessageField) {
	println(f.Self.Name)
	if f.Children != nil {
		for _, child := range f.Children {
			print(child)
		}
	}
}

// GetContent parses file data and returns tree structured list of LinkedMessageField
func GetContent(files []*File, baseMessage, prefix string) []*LinkedMessageField {
	if baseMessage != "" {
		baseMessage = fmt.Sprintf("%s.", baseMessage)
	}
	if prefix != "" {
		prefix = fmt.Sprintf("%s.", prefix)
	}
	var linkedFields []*LinkedMessageField
	for _, file := range files {
		for _, message := range file.Messages {
			if strings.ToLower(message.LongName) == strings.ToLower(baseMessage) {
				for _, field := range message.Fields {
					field.FullPath = fmt.Sprintf("%s%s%s", prefix, baseMessage, field.Name)
					linkedField := &LinkedMessageField{Self: field}
					if !isScalarType(field.LongType) {
						getChildField(files, linkedField, prefix)
					}
					linkedFields = append(linkedFields, linkedField)
				}
			}
		}
	}
	return linkedFields
}

func getChildField(files []*File, parentLinkedField *LinkedMessageField, prefix string) {
	for _, file := range files {
		for _, message := range file.Messages {
			if message.FullName == parentLinkedField.Self.FullType {
				for _, field := range message.Fields {
					field.FullPath = fmt.Sprintf("%s%s.%s", prefix, parentLinkedField.Self.LongType, field.Name)
					linkedField := &LinkedMessageField{Self: field}
					if !isScalarType(field.LongType) {
						getChildField(files, linkedField, prefix)
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

func isScalarType(fieldType string) bool {
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
