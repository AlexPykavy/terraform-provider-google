// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccBigQueryDataset_bigqueryDatasetBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryDatasetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryDataset_bigqueryDatasetBasicExample(context),
			},
			{
				ResourceName:      "google_bigquery_dataset.dataset",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBigQueryDataset_bigqueryDatasetBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "dataset" {
  dataset_id                  = "tf_test_example_dataset%{random_suffix}"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "EU"
  default_table_expiration_ms = 3600000

  labels = {
    env = "default"
  }

  access {
    role          = "OWNER"
    user_by_email = google_service_account.bqowner.email
  }

  access {
    role   = "READER"
    domain = "hashicorp.com"
  }
}

resource "google_service_account" "bqowner" {
  account_id = "bqowner%{random_suffix}"
}
`, context)
}

func TestAccBigQueryDataset_bigqueryDatasetWithMaxTimeTravelHoursExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryDatasetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryDataset_bigqueryDatasetWithMaxTimeTravelHoursExample(context),
			},
			{
				ResourceName:      "google_bigquery_dataset.dataset",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBigQueryDataset_bigqueryDatasetWithMaxTimeTravelHoursExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "dataset" {
  dataset_id                  = "tf_test_example_dataset%{random_suffix}"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "EU"
  default_table_expiration_ms = 3600000
  max_time_travel_hours       = "72"

  labels = {
    env = "default"
  }

  access {
    role          = "OWNER"
    user_by_email = google_service_account.bqowner.email
  }

  access {
    role   = "READER"
    domain = "hashicorp.com"
  }
}

resource "google_service_account" "bqowner" {
  account_id = "bqowner%{random_suffix}"
}
`, context)
}

func TestAccBigQueryDataset_bigqueryDatasetAuthorizedDatasetExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryDatasetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryDataset_bigqueryDatasetAuthorizedDatasetExample(context),
			},
			{
				ResourceName:      "google_bigquery_dataset.dataset",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBigQueryDataset_bigqueryDatasetAuthorizedDatasetExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "public" {
  dataset_id                  = "public%{random_suffix}"
  friendly_name               = "test"
  description                 = "This dataset is public"
  location                    = "EU"
  default_table_expiration_ms = 3600000

  labels = {
    env = "default"
  }

  access {
    role          = "OWNER"
    user_by_email = google_service_account.bqowner.email
  }

  access {
    role   = "READER"
    domain = "hashicorp.com"
  }
}

resource "google_bigquery_dataset" "dataset" {
  dataset_id                  = "private%{random_suffix}"
  friendly_name               = "test"
  description                 = "This dataset is private"
  location                    = "EU"
  default_table_expiration_ms = 3600000

  labels = {
    env = "default"
  }

  access {
    role          = "OWNER"
    user_by_email = google_service_account.bqowner.email
  }

  access {
    role   = "READER"
    domain = "hashicorp.com"
  }

  access {
    dataset {
      dataset {
        project_id = google_bigquery_dataset.public.project
        dataset_id = google_bigquery_dataset.public.dataset_id
      }
      target_types = ["VIEWS"]
    }
  }
}

resource "google_service_account" "bqowner" {
  account_id = "bqowner%{random_suffix}"
}
`, context)
}

func TestAccBigQueryDataset_bigqueryDatasetAuthorizedRoutineExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"service_account": getTestServiceAccountFromEnv(t),
		"random_suffix":   randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigQueryDatasetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryDataset_bigqueryDatasetAuthorizedRoutineExample(context),
			},
			{
				ResourceName:      "google_bigquery_dataset.private",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBigQueryDataset_bigqueryDatasetAuthorizedRoutineExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "public" {
  dataset_id  = "tf_test_public_dataset%{random_suffix}"
  description = "This dataset is public"
}

resource "google_bigquery_routine" "public" {
  dataset_id      = google_bigquery_dataset.public.dataset_id
  routine_id      = "tf_test_public_routine%{random_suffix}"
  routine_type    = "TABLE_VALUED_FUNCTION"
  language        = "SQL"
  definition_body = <<-EOS
    SELECT 1 + value AS value
  EOS
  arguments {
    name          = "value"
    argument_kind = "FIXED_TYPE"
    data_type     = jsonencode({ "typeKind" = "INT64" })
  }
  return_table_type = jsonencode({ "columns" = [
    { "name" = "value", "type" = { "typeKind" = "INT64" } },
  ] })
}

resource "google_bigquery_dataset" "private" {
  dataset_id  = "tf_test_private_dataset%{random_suffix}"
  description = "This dataset is private"
  access {
    role          = "OWNER"
    user_by_email = "%{service_account}"
  }
  access {
    routine {
      project_id = google_bigquery_routine.public.project
      dataset_id = google_bigquery_routine.public.dataset_id
      routine_id = google_bigquery_routine.public.routine_id
    }
  }
}
`, context)
}

func testAccCheckBigQueryDatasetDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_bigquery_dataset" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("BigQueryDataset still exists at %s", url)
			}
		}

		return nil
	}
}
