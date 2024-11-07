// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = EnabledFunction{}
)

func NewEnabledFunction() function.Function {
	return EnabledFunction{}
}

type EnabledFunction struct{}

func (r EnabledFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "enabled"
}

func (r EnabledFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Check feature is enabled",
		Parameters: []function.Parameter{
			function.SetParameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				Description:        "The feature names to check",
				Name:               "names",
				ElementType:        types.StringType,
			}, function.MapParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The features map to check",
				Name:               "map",
				ElementType:        types.BoolType,
			},
		},
		Return: function.BoolReturn{},
	}
}

func (r EnabledFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var names []string
	var features *map[string]*bool

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &names, &features))

	if resp.Error != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, hasAllKeys(names, features)))
}

func hasAllKeys(names []string, features *map[string]*bool) bool {
	if features == nil || names == nil {
		return false
	}

	for _, name := range names {
		if enabled, ok := (*features)[name]; !ok || enabled == nil || !*enabled {
			return false
		}
	}
	return true
}
