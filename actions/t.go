// Copyright 2021 Juan Pablo Tosso
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package actions

import (
	"fmt"

	engine "github.com/jptosso/coraza-waf"
	"github.com/jptosso/coraza-waf/transformations"
)

type T struct{}

func (a *T) Init(r *engine.Rule, transformation string) error {
	if transformation == "none" {
		//remove elements
		r.Transformations = r.Transformations[:0]
		return nil
	}
	transformations := transformations.TransformationsMap()
	tt := transformations[transformation]
	if tt == nil {
		return fmt.Errorf("Unsupported transformation %s", transformation)
	}
	r.Transformations = append(r.Transformations, tt)
	return nil
}

func (a *T) Evaluate(r *engine.Rule, tx *engine.Transaction) {
	// Not evaluated
}

func (a *T) Type() int {
	return engine.ACTION_TYPE_NONDISRUPTIVE
}