package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceAsset_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceAsset_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_asset.asset_1", "url", "http://example.com/asset/example.tar.gz"),
				),
			},
		},
	})
}

func TestAccDataSourceCheck_headers(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceAsset_headers,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_asset.asset_1", "headers.header1", "test1"),
					resource.TestCheckResourceAttr(
						"data.sensu_asset.asset_1", "headers.header2", "test2"),
				),
			},
		},
	})
}

var testAccDataSourceAsset_basic = fmt.Sprintf(`
  %s

  data "sensu_asset" "asset_1" {
    name = "${sensu_asset.asset_1.name}"
  }
`, testAccResourceAsset_basic)

var testAccDataSourceAsset_headers = fmt.Sprintf(`
  %s

  data "sensu_asset" "asset_1" {
    name = "${sensu_asset.asset_1.name}"
  }
`, testAccResourceAsset_headers_1)
