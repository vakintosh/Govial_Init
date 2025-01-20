package kubernetes

import (
	"fmt"
	"os"
	"path/filepath"
)

// HelmChart defines the structure of a Helm chart.
type HelmChart struct {
	AppName     string
	Replicas    int
	Image       string
	Port        int
	ServiceType string
}

func GenerateHelmChart(outputDir string, chart HelmChart) error {
	chartDir := filepath.Join(outputDir, chart.AppName)
	if err := os.MkdirAll(filepath.Join(chartDir, "templates"), os.ModePerm); err != nil {
		return fmt.Errorf("error creating chart directory: %v", err)
	}

	chartYaml := fmt.Sprintf(`apiVersion: v2
name: %s
description: A Helm chart for %s
version: 0.1.0
`, chart.AppName, chart.AppName)
	if err := os.WriteFile(filepath.Join(chartDir, "Chart.yaml"), []byte(chartYaml), os.ModePerm); err != nil {
		return fmt.Errorf("error writing Chart.yaml: %v", err)
	}

	valuesYaml := fmt.Sprintf(`replicaCount: %d
image: %s
port: %d
`, chart.Replicas, chart.Image, chart.Port)
	if err := os.WriteFile(filepath.Join(chartDir, "values.yaml"), []byte(valuesYaml), os.ModePerm); err != nil {
		return fmt.Errorf("error writing values.yaml: %v", err)
	}

	templates := []string{"deployment.yaml.tpl", "service.yaml.tpl", "configmap.yaml.tpl"}
	for _, tmpl := range templates {
		templatePath := filepath.Join("pkg/kubernetes/templates", tmpl)
		err := GenerateManifest(filepath.Join(chartDir, "templates"), chart, templatePath)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Helm chart generated at: %s\n", chartDir)
	return nil
}
