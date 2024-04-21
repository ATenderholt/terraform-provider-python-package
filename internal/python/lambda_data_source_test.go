package python_test

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"os"
	"testing"
)

func TestAccAwsLambda_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest:                false,
		PreCheck:                  nil,
		ProtoV6ProviderFactories:  protoV6ProviderFactories(),
		PreventPostDestroyRefresh: false,
		CheckDestroy:              nil,
		ErrorCheck:                nil,
		Steps: []resource.TestStep{
			{
				Config: config("example_without_deps"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testFileExists("output/example_without_deps.zip"),
					resource.TestCheckResourceAttr("data.python_aws_lambda.test", "archive_base64sha256", "5194175f5f67b6492d02d22a339baaaca0ce0d39aa38edf0c07713bd277e7b2f"),
				),
			},
		},
		WorkingDir: "",
	})
}

func testFileExists(path string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		_, err := os.Stat(path)
		if err != nil {
			return err
		}
		return nil
	}
}

const configTemplate = `
provider "python" {
  pip_command = "pip3.10"
}

data "python_aws_lambda" "test" {
  source_dir  = "test-fixtures/%s"
  archive_path = "output/%s.zip"
}
`

func config(name string) string {
	return fmt.Sprintf(
		configTemplate,
		name,
		name,
	)
}
