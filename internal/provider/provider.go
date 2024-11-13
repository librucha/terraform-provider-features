// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure FeaturesProvider satisfies various provider interfaces.
var _ provider.Provider = &FeaturesProvider{}
var _ provider.ProviderWithFunctions = &FeaturesProvider{}

// FeaturesProvider defines the provider implementation.
type FeaturesProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func (p *FeaturesProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "features"
	resp.Version = p.version
}

func (p *FeaturesProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{},
	}
}

func (p *FeaturesProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *FeaturesProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *FeaturesProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *FeaturesProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewEnabledFunction,
		NewCountFunction,
		NewMergeFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &FeaturesProvider{
			version: version,
		}
	}
}
