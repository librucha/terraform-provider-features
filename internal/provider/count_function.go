// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = CountFunction{}
)

func NewCountFunction() function.Function {
	return CountFunction{}
}

type CountFunction struct{}

func (r CountFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "count"
}

func (r CountFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Check feature will be used in count attribute",
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
		Return: function.Int32Return{},
	}
}

func (r CountFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var names []string
	var features *map[string]*bool

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &names, &features))

	if resp.Error != nil {
		return
	}

	count := 0
	if hasAllKeys(names, features) {
		count = 1
	}
	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, count))
}
