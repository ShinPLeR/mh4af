package model

import (
	"fmt"
	"sort"
	"strings"
)

/** yaml models */

type Option struct {
	AutoIncrement bool   `yaml:"auto_increment" validate:"boolean"`
	OnUpdate      string `yaml:"on_update" validate:"ascii"`
}

type Column struct {
	Pk        bool     `yaml:"pk" validate:"boolean"`
	Fk        bool     `yaml:"fk" validate:"boolean"`
	Unique    bool     `yaml:"unique" validate:"boolean"`
	Type      string   `yaml:"type" validate:"required"`
	Default   string   `yaml:"default" validate:""`
	Nullable  bool     `yaml:"nullable" validate:"boolean"`
	Reference string   `yaml:"reference" validate:"required_if fk,excluded_if pk"`
	On        string   `yaml:"on" validate:"required_if fk,excluded_if pk"`
	Cascade   []string `yaml:"cascade" validate:"dive,alpha"`
	Options   Option   `yaml:"options" validate:"dive"`
	Comment   string   `yaml:"comment" validate:""`
}

type Table struct {
	Columns     []map[string]Column `yaml:"columns" validate:"dive,required"`
	Description string              `yaml:"description"`
}

type Schema struct {
	DbName string           `yaml:"name" validate:"ascii"`
	Tables map[string]Table `yaml:"tables" validate:"dive,required"`
}

type Definition struct {
	Version string `yaml:"version" validate:"alphanum"`
	Title   string `yaml:"title" validate:"required"`
	Author  string `yaml:"author" validated:"unique"`
	Schema  Schema `yaml:"schema" validate:"required,dive,required,unique"`
}

type RearrangedOption struct {
	Option
	AutoIncrement bool
	OnUpdate      string
}

type RearrangedColumn struct {
	Column
	Name        string
	Pk          bool
	Fk          bool
	Type        string
	Nullable    bool
	Default     string
	Option      RearrangedOption
	Constraints []string
	Comment     string
}

type RearrangedTable struct {
	Table
	Name    string
	Columns []RearrangedColumn
}

type RearrangedSchema struct {
	Schema
	Tables []RearrangedTable
}

type RearrangedDefinition struct {
	Definition
	Version string
	Author  string
	Title   string
	Schema  RearrangedSchema
}

/** mutate schemas */

func (o Option) Rearrange() RearrangedOption {
	model := RearrangedOption{
		AutoIncrement: o.AutoIncrement,
		OnUpdate:      o.OnUpdate,
	}

	return model
}

func (c Column) FormatConstraints() []string {
	var result []string
	// foreign constraint
	if c.Fk {
		if c.Reference == "" {
			panic("Fk requires 'reference' parameter")
		}
		if c.On == "" {
			panic("Fk requires 'on' parameter")
		}
		result = append(result, fmt.Sprintf("FK (%s.%s)", c.On, c.Reference))
	}

	// unique constraint
	if c.Unique {
		result = append(result, "UNIQUE")
	}

	return result
}

func (s Schema) Rearrange() RearrangedSchema {
	model := RearrangedSchema{
		Tables: getRearrangedTables(s.Tables),
	}

	return model
}

func (d Definition) Rearrange() (RearrangedDefinition, error) {
	model := RearrangedDefinition{
		Version: d.Version,
		Author:  d.Author,
		Title:   d.Title,
		Schema:  d.Schema.Rearrange(),
	}

	return model, nil
}

func getRearrangedColumns(cols []map[string]Column) []RearrangedColumn {
	var reCols []RearrangedColumn
	for _, val := range cols {
		for key, va := range val {
			reCols = append(reCols, RearrangedColumn{
				Name:        key,
				Pk:          va.Pk,
				Fk:          va.Fk,
				Type:        strings.ToLower(va.Type),
				Nullable:    va.Nullable,
				Default:     va.Default,
				Constraints: va.FormatConstraints(),
				Option:      va.Options.Rearrange(),
				Comment:     va.Comment,
			})
		}

	}
	return reCols
}

func getRearrangedTables(tables map[string]Table) []RearrangedTable {
	var keys []string
	for key := range tables {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var reTables []RearrangedTable
	for _, key := range keys {
		reTables = append(reTables, RearrangedTable{
			Name:    key,
			Columns: getRearrangedColumns(tables[key].Columns),
		})
	}

	return reTables
}
