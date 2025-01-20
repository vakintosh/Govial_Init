package terraform

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetTerraformVersion checks the Terraform version and returns true if it is 0.13 or later
func GetTerraformVersion() (bool, error) {
	cmd := exec.Command("terraform", "version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("failed to get terraform version: %v", err)
	}

	// Check if the version is 0.13 or later by parsing the output
	versionOutput := string(output)
	if strings.Contains(versionOutput, "v0.13") || strings.Contains(versionOutput, "v0.14") || strings.Contains(versionOutput, "v0.15") || strings.Contains(versionOutput, "v1") {
		return true, nil
	}
	return false, nil
}
