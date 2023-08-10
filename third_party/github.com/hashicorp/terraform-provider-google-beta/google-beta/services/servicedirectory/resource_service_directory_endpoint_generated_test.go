// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package servicedirectory_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccServiceDirectoryEndpoint_serviceDirectoryEndpointBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckServiceDirectoryEndpointDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDirectoryEndpoint_serviceDirectoryEndpointBasicExample(context),
			},
			{
				ResourceName:            "google_service_directory_endpoint.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service", "endpoint_id"},
			},
		},
	})
}

func testAccServiceDirectoryEndpoint_serviceDirectoryEndpointBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "tf-test-example-namespace%{random_suffix}"
  location     = "us-central1"
}

resource "google_service_directory_service" "example" {
  provider   = google-beta
  service_id = "tf-test-example-service%{random_suffix}"
  namespace  = google_service_directory_namespace.example.id
}

resource "google_service_directory_endpoint" "example" {
  provider    = google-beta
  endpoint_id = "tf-test-example-endpoint%{random_suffix}"
  service     = google_service_directory_service.example.id

  metadata = {
    stage  = "prod"
    region = "us-central1"
  }

  address = "1.2.3.4"
  port    = 5353
}
`, context)
}

func TestAccServiceDirectoryEndpoint_serviceDirectoryEndpointWithNetworkExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckServiceDirectoryEndpointDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccServiceDirectoryEndpoint_serviceDirectoryEndpointWithNetworkExample(context),
			},
			{
				ResourceName:            "google_service_directory_endpoint.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service", "endpoint_id"},
			},
		},
	})
}

func testAccServiceDirectoryEndpoint_serviceDirectoryEndpointWithNetworkExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
data "google_project" "project" {
  provider  = google-beta
}

resource "google_compute_network" "example" {
  provider  = google-beta
  name      = "tf-test-example-network%{random_suffix}"
}

resource "google_service_directory_namespace" "example" {
  provider     = google-beta
  namespace_id = "tf-test-example-namespace%{random_suffix}"
  location     = "us-central1"
}

resource "google_service_directory_service" "example" {
  provider   = google-beta
  service_id = "tf-test-example-service%{random_suffix}"
  namespace  = google_service_directory_namespace.example.id
}

resource "google_service_directory_endpoint" "example" {
  provider    = google-beta
  endpoint_id = "tf-test-example-endpoint%{random_suffix}"
  service     = google_service_directory_service.example.id

  metadata = {
    stage  = "prod"
    region = "us-central1"
  }

  network = "projects/${data.google_project.project.number}/locations/global/networks/${google_compute_network.example.name}"
  address = "1.2.3.4"
  port    = 5353
}
`, context)
}

func testAccCheckServiceDirectoryEndpointDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_service_directory_endpoint" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ServiceDirectoryBasePath}}{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ServiceDirectoryEndpoint still exists at %s", url)
			}
		}

		return nil
	}
}