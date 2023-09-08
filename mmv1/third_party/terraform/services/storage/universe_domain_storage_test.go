package storage_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccUniverseDomainStorage(t *testing.T) {
	// Skip VCR since this test can only run in specific test project.
	// Location field from `google_storage_bucket` needs to be changed depending on the universe.
	t.Skip()

	universeDomain := envvar.GetTestUniverseDomainFromEnv(t)

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccStorageBucketDestroyProducer(t),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccUniverseDomain_bucket(universeDomain),
			},
		},
	})
}

func testAccUniverseDomain_bucket(universeDomain string) string {
	return fmt.Sprintf(`
provider "google" {
  universe_domain = "%s"
}
	  
resource "google_storage_bucket" "foo" {
  name     = "storage_test_skip"
  location = "US"
}
  
data "google_storage_bucket" "bar" {
  name = google_storage_bucket.foo.name
  depends_on = [
	google_storage_bucket.foo,
  ]
}
`, universeDomain)
}
