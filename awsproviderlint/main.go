// The awsproviderlint command is a static checker for the Terraform AWS Provider.
package main

import (
	awspasses "github.com/Cazoo-uk/terraform-provider-cazoo-aws/awsproviderlint/passes"
	tfpasses "github.com/bflad/tfproviderlint/passes"
	tfxpasses "github.com/bflad/tfproviderlint/xpasses"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	var analyzers []*analysis.Analyzer
	analyzers = append(analyzers, tfpasses.AllChecks...)
	analyzers = append(analyzers, tfxpasses.AllChecks...)
	analyzers = append(analyzers, awspasses.AllChecks...)
	multichecker.Main(analyzers...)
}
