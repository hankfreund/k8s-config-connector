---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file in
#     .github/CONTRIBUTING.md.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud TPU v2"
description: |-
  A Cloud TPU VM instance.
---

# google\_tpu\_v2\_vm

A Cloud TPU VM instance.

~> **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
See [Provider Versions](https://terraform.io/docs/providers/google/guides/provider_versions.html) for more details on beta resources.

To get more information about Vm, see:

* [API documentation](https://cloud.google.com/tpu/docs/reference/rest/v2/projects.locations.nodes)
* How-to Guides
    * [Official Documentation](https://cloud.google.com/tpu/docs/)

<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=tpu_v2_vm_basic&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Tpu V2 Vm Basic


```hcl
data "google_tpu_v2_runtime_versions" "available" {
  provider = google-beta
}

resource "google_tpu_v2_vm" "tpu" {
  provider = google-beta

  name = "test-tpu"
  zone = "us-central1-c"

  runtime_version = "tpu-vm-tf-2.13.0"
}
```
<div class = "oics-button" style="float: right; margin: 0 0 -15px">
  <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=tpu_v2_vm_full&cloudshell_image=gcr.io%2Fcloudshell-images%2Fcloudshell%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
    <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
  </a>
</div>
## Example Usage - Tpu V2 Vm Full


```hcl
data "google_tpu_v2_runtime_versions" "available" {
  provider = google-beta
}

data "google_tpu_v2_accelerator_types" "available" {
  provider = google-beta
}

resource "google_tpu_v2_vm" "tpu" {
  provider = google-beta

  name = "test-tpu"
  zone = "us-central1-c"
  description = "Text description of the TPU."

  runtime_version  = "tpu-vm-tf-2.13.0"
  accelerator_type = "v2-8"
}
```

## Argument Reference

The following arguments are supported:


* `name` -
  (Required)
  The immutable name of the TPU.

* `runtime_version` -
  (Required)
  Runtime version for the TPU.


- - -


* `accelerator_type` -
  (Optional)
  TPU accelerator type for the TPU. If not specified, this defaults to 'v2-8'.

* `description` -
  (Optional)
  Text description of the TPU.

* `zone` -
  (Optional)
  The GCP location for the TPU. If it is not provided, the provider zone is used.

* `project` - (Optional) The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `projects/{{project}}/locations/{{zone}}/nodes/{{name}}`


## Timeouts

This resource provides the following
[Timeouts](https://developer.hashicorp.com/terraform/plugin/sdkv2/resources/retries-and-customizable-timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


Vm can be imported using any of these accepted formats:

```
$ terraform import google_tpu_v2_vm.default projects/{{project}}/locations/{{zone}}/nodes/{{name}}
$ terraform import google_tpu_v2_vm.default {{project}}/{{zone}}/{{name}}
$ terraform import google_tpu_v2_vm.default {{zone}}/{{name}}
$ terraform import google_tpu_v2_vm.default {{name}}
```

## User Project Overrides

This resource supports [User Project Overrides](https://registry.terraform.io/providers/hashicorp/google/latest/docs/guides/provider_reference#user_project_override).
