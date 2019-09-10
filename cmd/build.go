package cmd

import (
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "k8comp",
	Long:    "k8comp is a tool which substitutes any templates variables declared with values from a files hierarchy using hiera.",
	Example: "k8comp [-h | pull | -p <project_name> -a <application> -e <environment> -b <git_branch>]",
	Run:     buildTemplate,
}

func init() {
	addApplicationFlag(buildCmd)
	addBranchFlag(buildCmd)
	addEnvironmentFlag(buildCmd)
	addNamespaceFlag(buildCmd)
	addLocationFlag(buildCmd)
	addProductFlag(buildCmd)
	addProfileFlag(buildCmd)
	addProjectFlag(buildCmd)
	addRoleFlag(buildCmd)
	addTemplateFlag(buildCmd)
	addXtraFlag(buildCmd)
}

func buildTemplate(cmd *cobra.Command, args []string) {
}
