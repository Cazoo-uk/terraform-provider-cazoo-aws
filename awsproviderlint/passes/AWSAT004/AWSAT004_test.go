package AWSAT004_test

import (
	"testing"

	"github.com/Cazoo-uk/terraform-provider-cazoo-aws/awsproviderlint/passes/AWSAT004"
	_ "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAWSAT004(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, AWSAT004.Analyzer, "a")
}
