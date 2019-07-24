// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// +build ignore

package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"text/template"
)

var tpl = template.Must(template.New("normalizations").Parse(`
// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by mknormalize_data.go - DO NOT EDIT.

package aucoalesce

import (
	"encoding/base64"
	"fmt"
)

var assets map[string][]byte

func asset(key string) ([]byte, error) {
	if assets == nil {
		assets = map[string][]byte{}

		var value []byte
		{{- range $asset := .Assets }}
		value, _ = base64.StdEncoding.DecodeString("{{ $asset.Data }}")
		assets["{{ $asset.Name }}"] = value{{ end }}
	}

	if value, found := assets[key]; found {
		return value, nil
	}
	return nil, fmt.Errorf("asset not found for key=%v", key)
}
`))

type asset struct {
	Name string
	Data string
}

var assets []asset

func addAsset(key, file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	assets = append(assets, asset{
		Name: key,
		Data: base64.StdEncoding.EncodeToString(data),
	})
	return nil
}

type templateVars struct {
	Assets []asset
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args)%2 != 0 {
		fmt.Fprintln(os.Stderr, "error: expected pairs of arguments (a key and a filename)")
		os.Exit(1)
	}

	for i := 0; i < len(args); i += 2 {
		addAsset(args[i], args[i+1])
	}

	var buf bytes.Buffer
	tpl.Execute(&buf, templateVars{
		Assets: assets,
	})
	bs, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(bs)
}
