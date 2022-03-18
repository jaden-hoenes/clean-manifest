package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Sizing struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Manifest struct {
	Name        string        `json:"name"`
	Id          string        `json:"id,omitempty"`
	Version     string        `json:"version"`
	Fullpage    bool          `json:"fullpage,omitempty"`
	Size        Sizing        `json:"size"`
	Mapping     []DataMapping `json:"mapping"`
	ProxyId     string        `json:"proxyId,omitempty"`
	Ignore      []string      `json:"ignore,omitempty"`
	Collections []Collection  `json:"collections,omitempty"`
	Path        string        `json:"-"`
}

type DataField struct {
	Alias      string `json:"alias"`
	ColumnName string `json:"columnName"`
}

type DataMapping struct {
	Alias     string      `json:"alias"`
	DataSetId string      `json:"dataSetId,omitempty"`
	Fields    []DataField `json:"fields"`
}

type SchemaColumn struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type Schema struct {
	Columns []SchemaColumn `json:"columns,omitempty"`
}

type Authorities struct {
	CreateContent []string `json:"CREATE_CONTENT,omitempty"`
	ReadContent   []string `json:"READ_CONTENT,omitempty"`
	UpdateContent []string `json:"UPDATE_CONTENT,omitempty"`
	DeleteContent []string `json:"DELETE_CONTENT,omitempty"`
}

type Collection struct {
	Name                string       `json:"name"`
	Id                  string       `json:"id,omitempty"`
	Schema              *Schema      `json:"schema,omitempty"`
	SyncEnabled         bool         `json:"syncEnabled,omitempty"`
	DefaultPermission   []string     `json:"defaultPermission,omitempty"`
	RequiredAuthorities *Authorities `json:"requiredAuthorities,omitempty"`
}

func main() {
	var manifest Manifest

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &manifest)
	if err != nil {
		log.Fatal(err)
	}

	manifest.Id = ""
	manifest.ProxyId = ""

	for index := range manifest.Mapping {
		manifest.Mapping[index].DataSetId = ""
	}
	for index := range manifest.Collections {
		manifest.Collections[index].Id = ""
	}

	file, _ := json.MarshalIndent(manifest, "", "  ")
	ioutil.WriteFile(os.Args[1], file, 0644)

	fmt.Println("\033[32m✓\033[0m Successfully cleaned manifest")
}
