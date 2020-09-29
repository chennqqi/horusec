// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package workdir

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	t.Run("should success parse workdir to string", func(t *testing.T) {
		workDir := &WorkDir{}
		assert.NotEmpty(t, workDir.String())
	})
}

func TestSet(t *testing.T) {
	t.Run("should success parse interface to workdir", func(t *testing.T) {
		workDir := &WorkDir{}

		assert.NotPanics(t, func() {
			workDir.ParseInterfaceToStruct(workDir)
		})
	})
}

func TestType(t *testing.T) {
	t.Run("should return type", func(t *testing.T) {
		workDir := &WorkDir{}
		assert.Empty(t, workDir.Type())
	})
}

func TestMap(t *testing.T) {
	t.Run("should success return a new map of workdir", func(t *testing.T) {
		workDir := &WorkDir{}
		assert.NotEmpty(t, workDir.Map())
	})
}
