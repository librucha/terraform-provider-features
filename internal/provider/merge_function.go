// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = MergeFunction{}
)

func NewMergeFunction() function.Function {
	return MergeFunction{}
}

type MergeFunction struct{}

func (r MergeFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "merge"
}

func (r MergeFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Merge maps or objects by features rules",
		Parameters: []function.Parameter{
			function.MapParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The default features setting",
				Name:               "default_map",
				ElementType:        types.BoolType,
			}, function.MapParameter{
				AllowNullValue:     true,
				AllowUnknownValues: false,
				Description:        "The features map to merge into default features map",
				Name:               "map",
				ElementType:        types.BoolType,
			},
		},
		Return: function.MapReturn{
			ElementType: types.BoolType,
		},
	}
}

func (r MergeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var defaultFeatures *map[string]*bool
	var features *map[string]*bool

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &defaultFeatures, &features))

	if resp.Error != nil {
		return
	}
	for k, _ := range *defaultFeatures {
		val, ok := (*features)[k]
		if ok {
			(*defaultFeatures)[k] = val
		}
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, defaultFeatures))
}
