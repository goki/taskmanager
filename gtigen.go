// Code generated by "goki generate"; DO NOT EDIT.

package main

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:      "main.Task",
	ShortName: "main.Task",
	IDName:    "task",
	Doc:       "",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Name", &gti.Field{Name: "Name", Type: "string", LocalType: "string", Doc: "The name of this task", Directives: gti.Directives{}, Tag: ""}},
		{"CPU", &gti.Field{Name: "CPU", Type: "float64", LocalType: "float64", Doc: "The percentage of the CPU time this task uses", Directives: gti.Directives{}, Tag: ""}},
		{"RAM", &gti.Field{Name: "RAM", Type: "float32", LocalType: "float32", Doc: "The percentage of total RAM this task uses", Directives: gti.Directives{}, Tag: ""}},
		{"Threads", &gti.Field{Name: "Threads", Type: "int32", LocalType: "int32", Doc: "The number of threads this task uses", Directives: gti.Directives{}, Tag: ""}},
		{"User", &gti.Field{Name: "User", Type: "string", LocalType: "string", Doc: "The user that started this task", Directives: gti.Directives{}, Tag: ""}},
		{"PID", &gti.Field{Name: "PID", Type: "int32", LocalType: "int32", Doc: "The Process ID (PID) of this task", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Process", &gti.Field{Name: "Process", Type: "*github.com/shirou/gopsutil/v3/process.Process", LocalType: "*process.Process", Doc: "", Directives: gti.Directives{}, Tag: "view:\"-\""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})
