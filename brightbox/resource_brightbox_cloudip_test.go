package brightbox

import (
	"fmt"
	"testing"

	"github.com/brightbox/gobrightbox"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	resourceName = "brightbox_cloudip.foobar"
)

func TestAccBrightboxCloudip_Basic(t *testing.T) {
	resourceName := resourceName
	var cloudip brightbox.CloudIP
	rInt := acctest.RandInt()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBrightboxCloudipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckBrightboxCloudipConfig_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBrightboxCloudipExists(resourceName, &cloudip),
					resource.TestCheckResourceAttr(
						resourceName, "name", fmt.Sprintf("foo-%d", rInt)),
					resource.TestCheckNoResourceAttr(
						resourceName, "target"),
				),
			},
		},
	})
}

func TestAccBrightboxCloudip_clear_name(t *testing.T) {
	var cloudip brightbox.CloudIP
	rInt := acctest.RandInt()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBrightboxCloudipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckBrightboxCloudipConfig_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBrightboxCloudipExists(resourceName, &cloudip),
					resource.TestCheckResourceAttr(
						resourceName, "name", fmt.Sprintf("foo-%d", rInt)),
					resource.TestCheckNoResourceAttr(
						resourceName, "target"),
				),
			},
			{
				Config: testAccCheckBrightboxCloudipConfig_empty_name,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBrightboxCloudipExists(resourceName, &cloudip),
					resource.TestCheckResourceAttr(
						resourceName, "name", ""),
					resource.TestCheckNoResourceAttr(
						resourceName, "target"),
				),
			},
		},
	})
}

func TestAccBrightboxCloudip_Mapped(t *testing.T) {
	var cloudip brightbox.CloudIP
	rInt := acctest.RandInt()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBrightboxCloudipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckBrightboxCloudipConfig_mapped(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBrightboxCloudipExists(resourceName, &cloudip),
					resource.TestCheckResourceAttr(
						resourceName, "name", fmt.Sprintf("bar-%d", rInt)),
				),
			},
		},
	})
}

func TestAccBrightboxCloudip_Remapped(t *testing.T) {
	var cloudip brightbox.CloudIP
	rInt := acctest.RandInt()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBrightboxCloudipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckBrightboxCloudipConfig_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBrightboxCloudipExists(resourceName, &cloudip),
					resource.TestCheckResourceAttr(
						resourceName, "name", fmt.Sprintf("foo-%d", rInt)),
				),
			},
			{
				Config: testAccCheckBrightboxCloudipConfig_mapped(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBrightboxCloudipExists(resourceName, &cloudip),
					resource.TestCheckResourceAttr(
						resourceName, "name", fmt.Sprintf("bar-%d", rInt)),
				),
			},
			{
				Config: testAccCheckBrightboxCloudipConfig_remapped(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBrightboxCloudipExists(resourceName, &cloudip),
					resource.TestCheckResourceAttr(
						resourceName, "name", fmt.Sprintf("baz-%d", rInt)),
				),
			},
			{
				Config: testAccCheckBrightboxCloudipConfig_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBrightboxCloudipExists(resourceName, &cloudip),
					resource.TestCheckResourceAttr(
						resourceName, "name", fmt.Sprintf("foo-%d", rInt)),
				),
			},
		},
	})
}

func testAccCheckBrightboxCloudipDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*CompositeClient).ApiClient

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "brightbox_cloudip" {
			continue
		}

		// Try to find the CloudIP
		_, err := client.CloudIP(rs.Primary.ID)

		// Wait

		if err != nil {
			apierror := err.(brightbox.ApiError)
			if apierror.StatusCode != 404 {
				return fmt.Errorf(
					"Error waiting for cloudip %s to be destroyed: %s",
					rs.Primary.ID, err)
			}
		}
	}

	return nil
}

func testAccCheckBrightboxCloudipExists(n string, cloudip *brightbox.CloudIP) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No CloudIP ID is set")
		}

		client := testAccProvider.Meta().(*CompositeClient).ApiClient

		// Try to find the CloudIP
		retrieveCloudip, err := client.CloudIP(rs.Primary.ID)

		if err != nil {
			return err
		}

		if retrieveCloudip.Id != rs.Primary.ID {
			return fmt.Errorf("CloudIP not found")
		}

		*cloudip = *retrieveCloudip

		return nil
	}
}

func testAccCheckBrightboxCloudipConfig_basic(rInt int) string {
	return fmt.Sprintf(`
resource "brightbox_cloudip" "foobar" {
	name = "foo-%d"
}
`, rInt)
}

const testAccCheckBrightboxCloudipConfig_empty_name = `

resource "brightbox_cloudip" "foobar" {
	name = ""
}
`

func testAccCheckBrightboxCloudipConfig_mapped(rInt int) string {
	return fmt.Sprintf(`

resource "brightbox_cloudip" "foobar" {
	name = "bar-%d"
	target = "${brightbox_server.boofar.interface}"
}

resource "brightbox_server" "boofar" {
	image = "${data.brightbox_image.foobar.id}"
	name = "bar-%d"
	server_groups = ["${data.brightbox_server_group.default.id}"]
}
%s%s`, rInt, rInt, TestAccBrightboxImageDataSourceConfig_blank_disk,
		TestAccBrightboxDataServerGroupConfig_default)
}

func testAccCheckBrightboxCloudipConfig_remapped(rInt int) string {
	return fmt.Sprintf(`
resource "brightbox_cloudip" "foobar" {
	name = "baz-%d"
	target = "${brightbox_server.fred.interface}"
}

resource "brightbox_server" "boofar" {
	image = "${data.brightbox_image.foobar.id}"
	name = "bar-%d"
	server_groups = ["${data.brightbox_server_group.default.id}"]
}

resource "brightbox_server" "fred" {
	image = "${data.brightbox_image.foobar.id}"
	name = "baz-%d"
	server_groups = ["${data.brightbox_server_group.default.id}"]
}
%s%s`, rInt, rInt, rInt, TestAccBrightboxImageDataSourceConfig_blank_disk,
		TestAccBrightboxDataServerGroupConfig_default)
}
