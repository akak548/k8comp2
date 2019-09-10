package cmd

import (
	"fmt"
	"github.com/akak548/k8comp/config"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = buildCmd

var (
	application, branch, environment, location, namespace string
	product, profile, project, role, template             string
	xtra                                                  []string
)

func Execute() {
	rootCmd.AddCommand(pullCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initializeConfig() {
	err := config.Initialize()
	if err != nil {

	}
}

func addProjectFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&project, "project", "p", "", "Project name as specified on the projects folder.")
}

func addApplicationFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&application, "application", "a", "", "The name of the application which need to be deployed.")
}

func addEnvironmentFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&environment, "environment", "e", "", "The environment will be checked from heira.")
}

func addProductFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&product, "product", "d", "", "The prouduct will be checked in hiera.")
}

func addRoleFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&role, "role", "r", "", "The role will be checked for hiera.")
}

func addNamespaceFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&namespace, "namespace", "n", "", "The namespace will be checked in hiera")
}

func addTemplateFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&template, "template", "t", "", "The templates folder that can be configured on the k8comp.conf.")
}

func addLocationFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&location, "location", "l", "", "The location will be checked from hiera.")
}

func addXtraFlag(cmd *cobra.Command) {
	cmd.Flags().StringSliceVarP(&xtra, "xtra", "x", make([]string, 0), "The variable specified on the cmd run will be used to update a value on the final deployment file.")
}

func addBranchFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&branch, "branch", "b", "", "Specify a branch from where to do the deployment.")
}

func addProfileFlag(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&profile, "profile", "arp", "", "Specify the AWS IAM profile to be used for AWS Secret manager secrets retrieval.")
}
