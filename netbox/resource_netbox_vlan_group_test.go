package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxVlanGroupFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
}
`, testName)
}
func TestAccNetboxVlanGroup_basic(t *testing.T) {
	testSlug := "vlan_group_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVlanGroupFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_vlan_group" "test_basic" {
  name       = "%s"
  slug       = "%s"
  vid_ranges  = [[1, 2], [3, 4]]
  tags       = []
}`, testName, testSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan_group.test_basic", "name", testName),
					resource.TestCheckResourceAttr("netbox_vlan_group.test_basic", "slug", testSlug),
					resource.TestCheckResourceAttr("netbox_vlan_group.test_basic", "description", ""),
					resource.TestCheckResourceAttr("netbox_vlan_group.test_basic", "tags.#", "0"),
					resource.TestCheckResourceAttr("netbox_vlan_group.test_basic", "vid_ranges.0.0", "1"),
					resource.TestCheckResourceAttr("netbox_vlan_group.test_basic", "vid_ranges.0.1", "2"),
					resource.TestCheckResourceAttr("netbox_vlan_group.test_basic", "vid_ranges.1.0", "3"),
					resource.TestCheckResourceAttr("netbox_vlan_group.test_basic", "vid_ranges.1.1", "4"),
				),
			},
			{
				ResourceName:      "netbox_vlan_group.test_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVlanGroup_with_dependencies(t *testing.T) {
	testSlug := "vlan_group_with_dependencies"
	testName := testAccGetTestName(testSlug)
	testDescription := "Test Description"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVlanGroupFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_vlan_group" "test_with_dependencies" {
  name        = "%s"
  slug        = "%s"
  description = "%s"
  scope_type  = "dcim.site"
  scope_id    = netbox_site.test.id
  vid_ranges  = [[1, 4094]]
  tags        = [netbox_tag.test.name]
}`, testName, testSlug, testDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan_group.test_with_dependencies", "name", testName),
					resource.TestCheckResourceAttr("netbox_vlan_group.test_with_dependencies", "slug", testSlug),
					resource.TestCheckResourceAttr("netbox_vlan_group.test_with_dependencies", "description", testDescription),
					resource.TestCheckResourceAttr("netbox_vlan_group.test_with_dependencies", "scope_type", "dcim.site"),
					resource.TestCheckResourceAttrPair("netbox_vlan_group.test_with_dependencies", "scope_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_vlan_group.test_with_dependencies", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_vlan_group.test_with_dependencies", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_vlan_group.test_with_dependencies",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_vlan_group", &resource.Sweeper{
		Name:         "netbox_vlan_group",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := ipam.NewIpamVlanGroupsListParams()
			res, err := api.Ipam.IpamVlanGroupsList(params, nil)
			if err != nil {
				return err
			}
			for _, vlan := range res.GetPayload().Results {
				if strings.HasPrefix(*vlan.Name, testPrefix) {
					deleteParams := ipam.NewIpamVlanGroupsDeleteParams().WithID(vlan.ID)
					_, err := api.Ipam.IpamVlanGroupsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a vlan group")
				}
			}
			return nil
		},
	})
}
