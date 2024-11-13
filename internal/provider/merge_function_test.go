// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestMergeFunction_Full(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				locals {
				  default_features = {
					"key1" = true
					"key2" = false
					"key3" = null
					"key4" = null
				  }
				  features = {
					"key1" = false
					"key2" = null
					"key4" = true
				  }
				}
				output "test_key1" {
				  value = coalesce(provider::features::merge(local.default_features, local.features).key1, "null")
				}
				output "test_key2" {
				  value = coalesce(provider::features::merge(local.default_features, local.features).key2, "null")
				}
				output "test_key3" {
				  value = coalesce(provider::features::merge(local.default_features, local.features).key3, "null")
				}
				output "test_key4" {
				  value = coalesce(provider::features::merge(local.default_features, local.features).key4, "null")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test_key1", "false"),
					resource.TestCheckOutput("test_key2", "null"),
					resource.TestCheckOutput("test_key3", "null"),
					resource.TestCheckOutput("test_key4", "true"),
				),
			},
		},
	})
}
