package commercetools

import (
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccProjectCreate_basic(t *testing.T) {
	rName := acctest.RandString(5)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"commercetools_project_settings.acctest_project_settings", "name", "Test this thing",
					),
					resource.TestCheckResourceAttr(
						"commercetools_project_settings.acctest_project_settings", "countries.#", "3",
					),
					resource.TestCheckResourceAttr(
						"commercetools_project_settings.acctest_project_settings", "currencies.#", "2",
					),
					resource.TestCheckResourceAttr(
						"commercetools_project_settings.acctest_project_settings", "languages.#", "4",
					),
					resource.TestCheckResourceAttr(
						"commercetools_project_settings.acctest_project_settings", "messages.enabled", "true",
					),
				),
			},
		},
	})
}

func testAccCheckProjectDestroy(s *terraform.State) error {
	return nil
}

func testAccProjectConfig(name string) string {
	return `
resource "commercetools_project_settings" "acctest_project_settings" {
	name       = "Test this thing"
	countries  = ["NL", "DE", "US"]
	currencies = ["EUR", "USD"]
	languages  = ["nl", "de", "en", "en-US"]
	external_oauth = {
		url = "https://example.com/oauth/token"
		authorization_header = "Bearer secret"
	}

	messages = {
	  enabled = true
	}
}`
}
