// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIapAppEngineVersionIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.httpsResourceAccessor",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapAppEngineVersionIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_app_engine_version_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s/versions/%s roles/iap.httpsResourceAccessor", getTestProjectFromEnv(), getTestProjectFromEnv(), "default", context["random_suffix"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccIapAppEngineVersionIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_iap_app_engine_version_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s/versions/%s roles/iap.httpsResourceAccessor", getTestProjectFromEnv(), getTestProjectFromEnv(), "default", context["random_suffix"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapAppEngineVersionIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.httpsResourceAccessor",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccIapAppEngineVersionIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_app_engine_version_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s/versions/%s roles/iap.httpsResourceAccessor user:admin@hashicorptest.com", getTestProjectFromEnv(), getTestProjectFromEnv(), "default", context["random_suffix"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapAppEngineVersionIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"role":          "roles/iap.httpsResourceAccessor",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIapAppEngineVersionIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_app_engine_version_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_web/appengine-%s/services/%s/versions/%s", getTestProjectFromEnv(), getTestProjectFromEnv(), "default", context["random_suffix"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIapAppEngineVersionIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
  name = "appengine-static-content-%{random_suffix}"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "version" {
  version_id      = "%{random_suffix}"
  service         = "default"
  runtime         = "nodejs10"
  noop_on_destroy = false
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/hello-world.zip"
    }
  }
  env_variables = {
    port = "8080"
  }
}

resource "google_iap_app_engine_version_iam_member" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  version_id = "${google_app_engine_standard_app_version.version.version_id}"
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccIapAppEngineVersionIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
  name = "appengine-static-content-%{random_suffix}"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "version" {
  version_id      = "%{random_suffix}"
  service         = "default"
  runtime         = "nodejs10"
  noop_on_destroy = false
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/hello-world.zip"
    }
  }
  env_variables = {
    port = "8080"
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_iap_app_engine_version_iam_policy" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  version_id = "${google_app_engine_standard_app_version.version.version_id}"
  policy_data = "${data.google_iam_policy.foo.policy_data}"
}
`, context)
}

func testAccIapAppEngineVersionIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
  name = "appengine-static-content-%{random_suffix}"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "version" {
  version_id      = "%{random_suffix}"
  service         = "default"
  runtime         = "nodejs10"
  noop_on_destroy = false
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/hello-world.zip"
    }
  }
  env_variables = {
    port = "8080"
  }
}

resource "google_iap_app_engine_version_iam_binding" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  version_id = "${google_app_engine_standard_app_version.version.version_id}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccIapAppEngineVersionIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_storage_bucket" "bucket" {
  name = "appengine-static-content-%{random_suffix}"
}

resource "google_storage_bucket_object" "object" {
  name   = "hello-world.zip"
  bucket = google_storage_bucket.bucket.name
  source = "./test-fixtures/appengine/hello-world.zip"
}

resource "google_app_engine_standard_app_version" "version" {
  version_id      = "%{random_suffix}"
  service         = "default"
  runtime         = "nodejs10"
  noop_on_destroy = false
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/hello-world.zip"
    }
  }
  env_variables = {
    port = "8080"
  }
}

resource "google_iap_app_engine_version_iam_binding" "foo" {
  project = "${google_app_engine_standard_app_version.version.project}"
  app_id = "${google_app_engine_standard_app_version.version.project}"
  service = "${google_app_engine_standard_app_version.version.service}"
  version_id = "${google_app_engine_standard_app_version.version.version_id}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
