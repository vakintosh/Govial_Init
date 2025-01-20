package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vakintosh/Govial_Init/pkg/k8s"
)

func K8sCommand() *cobra.Command {
	var replicas int
	var image string
	var port int
	var serviceType string
	var useHelm bool

	cmd := &cobra.Command{
		Use:   "k8s [app-name] [directory]",
		Short: "Generate Kubernetes manifests or Helm charts",
		Args:  cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			appName := args[0]
			outputDir := "."
			if len(args) > 1 {
				outputDir = args[1]
			}

			chart := k8s.HelmChart{
				AppName:     appName,
				Replicas:    replicas,
				Image:       image,
				Port:        port,
				ServiceType: serviceType,
			}

			if useHelm {
				err := k8s.GenerateHelmChart(outputDir, chart)
				if err != nil {
					fmt.Printf("Error generating Helm chart: %v\n", err)
				}
			} else {
				err := k8s.GenerateManifest(outputDir, chart, "pkg/kubernetes/templates/k8s.tpl")
				if err != nil {
					fmt.Printf("Error generating manifest: %v\n", err)
				}
			}
		},
	}

	cmd.Flags().IntVarP(&replicas, "replicas", "r", 1, "Number of replicas")
	cmd.Flags().StringVarP(&image, "image", "i", "nginx:latest", "Container image")
	cmd.Flags().IntVarP(&port, "port", "p", 80, "Container port")
	cmd.Flags().StringVarP(&serviceType, "service-type", "s", "ClusterIP", "Service type (ClusterIP, NodePort, etc.)")
	cmd.Flags().BoolVarP(&useHelm, "helm", "h", false, "Generate a Helm chart instead of standalone manifests")

	return cmd
}
