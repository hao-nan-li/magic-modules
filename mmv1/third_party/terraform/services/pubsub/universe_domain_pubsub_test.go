package pubsub_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccUniverseDomainPubSub(t *testing.T) {
	// Skip VCR since this test can only run in specific test project.
	t.Skip()

	universeDomain := envvar.GetTestUniverseDomainFromEnv(t)
	topic := fmt.Sprintf("tf-test-topic-%s", acctest.RandString(t, 10))
	subscription := fmt.Sprintf("tf-test-sub-%s", acctest.RandString(t, 10))

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPubsubSubscriptionDestroyProducer(t),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccUniverseDomain_basic_pubsub(universeDomain, topic, subscription),
			},
		},
	})
}

func testAccUniverseDomain_basic_pubsub(universeDomain, topic, subscription string) string {
	return fmt.Sprintf(`
provider "google" {
  universe_domain = "%s"
}
	  
resource "google_pubsub_topic" "foo" {
  name = "%s"
}
  
resource "google_pubsub_subscription" "foo" {
  name  = "%s"
  topic = google_pubsub_topic.foo.id
  
  message_retention_duration = "1200s"
  retain_acked_messages      = true
  ack_deadline_seconds       = 20
  expiration_policy {
	ttl = ""
  }
  enable_message_ordering    = false
}
`, universeDomain, topic, subscription)
}
