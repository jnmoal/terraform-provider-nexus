package nexus

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceSecurityLDAP(t *testing.T) {
	resName := "data.nexus_security_ldap.acceptance"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSecurityLDAP(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "ldap.#"),
					resource.TestCheckResourceAttr(resName, "ldap.#", "0"),
				),
			},
		},
	})
}

func testAccDataSourceSecurityLDAP() string {
	return fmt.Sprintf(`data "nexus_security_ldap" "acceptance" {}`)
}
