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

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccComputeRoute_routeBasicExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRoute_routeBasicExample(acctest.RandString(10)),
			},
			{
				ResourceName:      "google_compute_route.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRoute_routeBasicExample(val string) string {
	return fmt.Sprintf(`
resource "google_compute_route" "default" {
  name        = "network-route-%s"
  dest_range  = "15.0.0.0/24"
  network     = "${google_compute_network.default.name}"
  next_hop_ip = "10.132.1.5"
  priority    = 100
}

resource "google_compute_network" "default" {
  name = "compute-network-%s"
}
`, val, val,
	)
}

func testAccCheckComputeRouteDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_route" {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(rs, "https://www.googleapis.com/compute/beta/projects/{{project}}/global/routes/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeRoute still exists at %s", url)
		}
	}

	return nil
}