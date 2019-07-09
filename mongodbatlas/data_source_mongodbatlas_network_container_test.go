package mongodbatlas

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	matlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func TestAccDataSourceMongoDBAtlasNetworkContainer_basic(t *testing.T) {
	var container matlas.Container

	randInt := acctest.RandIntRange(0, 255)

	resourceName := "mongodbatlas_network_container.test"
	projectID := "5cf5a45a9ccf6400e60981b6" // Modify until project data source is created.
	cidrBlock := fmt.Sprintf("10.8.%d.0/24", randInt)
	dataSourceName := "data.mongodbatlas_network_container.test"

	providerName := "AWS"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMongoDBAtlasNetworkContainerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMongoDBAtlasNetworkContainerDSConfig(projectID, cidrBlock),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMongoDBAtlasNetworkContainerExists(resourceName, &container),
					testAccCheckMongoDBAtlasNetworkContainerAttributes(&container, providerName),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttr(resourceName, "provider_name", providerName),
					resource.TestCheckResourceAttrSet(resourceName, "provisioned"),
					resource.TestCheckResourceAttrSet(dataSourceName, "project_id"),
					resource.TestCheckResourceAttr(dataSourceName, "provider_name", providerName),
					resource.TestCheckResourceAttrSet(dataSourceName, "provisioned"),
				),
			},
		},
	})

}

func testAccMongoDBAtlasNetworkContainerDSConfig(projectID, cidrBlock string) string {
	return fmt.Sprintf(`
resource "mongodbatlas_network_container" "test" {
	project_id   		= "%s"
	atlas_cidr_block    = "%s"
	provider_name		= "AWS"
	region_name			= "EU_WEST_1"
}

data "mongodbatlas_network_container" "test" {
	project_id   		= mongodbatlas_network_container.test.project_id
	container_id		= mongodbatlas_network_container.test.id
}
`, projectID, cidrBlock)
}