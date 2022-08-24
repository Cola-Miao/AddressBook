package models

const (
	AppFolder      = "AddressBook"
	DataFolder     = "Contacts"
	LogsFolder     = "Logs"
	DocumentSuffix = "ctt"
	TextSuffix     = "txt"
	TemplatePath   = "models/template.tmpl"
)

type Contact struct {
	Name   string
	Number string
}
