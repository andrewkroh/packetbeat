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

package mage

import (
	"github.com/magefile/mage/mg"

	"github.com/elastic/beats/dev-tools/mage"
	"github.com/elastic/beats/dev-tools/mage/target/build"
	"github.com/elastic/beats/dev-tools/mage/target/common"
	"github.com/elastic/beats/dev-tools/mage/target/dashboards"
	"github.com/elastic/beats/dev-tools/mage/target/integtest"
	"github.com/elastic/beats/dev-tools/mage/target/unittest"
)

func init() {
	common.RegisterCheckDeps(Update.All)

	dashboards.RegisterImportDeps(build.Build, Update.Dashboards)

	unittest.RegisterGoTestDeps(Update.Fields)
	unittest.RegisterPythonTestDeps(Update.Fields)

	integtest.RegisterPythonTestDeps(Update.Fields)
}

var (
	// SelectLogic configures the types of project logic to use (OSS vs X-Pack).
	SelectLogic mage.ProjectType
)

type Update mg.Namespace

// Update is an alias for running fields, dashboards, config, includes.
func (Update) All() {
	mg.Deps(Update.Fields, Update.Dashboards, Update.Config, Update.Includes)
}

func (Update) Config() error {
	return config()
}

// Dashboards collects all the dashboards and generates index patterns.
func (Update) Dashboards() error {
	mg.Deps(fb.FieldsYML)
	return mage.KibanaDashboards(mage.OSSBeatDir("monitors/active"))
}

func (Update) Fields() {
	mg.Deps(fb.All)
}

func (Update) Includes() error {
	if SelectLogic != mage.OSSProject {
		return nil
	}
	return mage.GenerateIncludeListGo([]string{"monitors/active/*"}, nil)
}
