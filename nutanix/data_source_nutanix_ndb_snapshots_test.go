package nutanix

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const dataSourceNDBSnapshotsName = "data.nutanix_ndb_snapshots.test"

func TestAccEraSnapshotsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccEraPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEraSnapshotsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceNDBSnapshotsName, "snapshots.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceNDBSnapshotsName, "snapshots.0.owner_id"),
					resource.TestCheckResourceAttrSet(dataSourceNDBSnapshotsName, "snapshots.0.properties.#"),
					resource.TestCheckResourceAttr(dataSourceNDBSnapshotsName, "snapshots.0.metadata.#", "1"),
					resource.TestCheckResourceAttrSet(dataSourceNDBSnapshotsName, "snapshots.0.snapshot_uuid"),
					resource.TestCheckResourceAttr(dataSourceNDBSnapshotsName, "snapshots.0.status", "ACTIVE"),
				),
			},
		},
	})
}

func TestAccEraSnapshotsDataSource_WithFilters(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccEraPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccEraSnapshotsDataSourceConfigWithFilters(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceNDBSnapshotsName, "snapshots.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceNDBSnapshotsName, "snapshots.0.owner_id"),
					resource.TestCheckResourceAttrSet(dataSourceNDBSnapshotsName, "snapshots.0.properties.#"),
					resource.TestCheckResourceAttr(dataSourceNDBSnapshotsName, "snapshots.0.metadata.#", "1"),
					resource.TestCheckResourceAttrSet(dataSourceNDBSnapshotsName, "snapshots.0.snapshot_uuid"),
					resource.TestCheckResourceAttr(dataSourceNDBSnapshotsName, "snapshots.0.status", "ACTIVE"),
				),
			},
		},
	})
}

func testAccEraSnapshotsDataSourceConfig() string {
	return `
		data "nutanix_ndb_snapshots" "test" {}
	`
}

func testAccEraSnapshotsDataSourceConfigWithFilters() string {
	return `
		data "nutanix_ndb_time_machines" "test1" {}

		data "nutanix_ndb_snapshots" "test" {
			filters{
				time_machine_id = data.nutanix_ndb_time_machines.test1.time_machines.0.id
			}
		}
	`
}
