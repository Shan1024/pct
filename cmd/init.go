//todo: add copyright notice

package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/wso2/wum-uc/constant"
	"github.com/wso2/wum-uc/util"
	"gopkg.in/yaml.v2"
)

// initCmd represents the validate command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			initCurrentDirectory()
		} else if len(args) == 1 {
			initDirectory(args[0])
		} else {
			util.PrintErrorAndExit("Invalid number of argumants. Run with --help for more details about the argumants")
		}

	},
}

func init() {
	RootCmd.AddCommand(initCmd)
	RootCmd.Flags().BoolVarP(&isDebugLogsEnabled, "debug", "d", false, "Enable debug logs")
	RootCmd.Flags().BoolVarP(&isTraceLogsEnabled, "trace", "t", false, "Enable trace logs")
}

func initCurrentDirectory() {
	currentDirectory := "./"
	initDirectory(currentDirectory)
}

func initDirectory(destination string) {
	err := util.CreateDirectory(destination)
	util.HandleError(err)

	updateDescriptorFile := filepath.Join(destination, constant.UPDATE_DESCRIPTOR_FILE)
	updateDescriptor := util.UpdateDescriptor{}
	data, err := yaml.Marshal(&updateDescriptor)
	if err != nil {
		util.HandleError(err)
	}

	file, err := os.OpenFile(
		updateDescriptorFile,
		os.O_WRONLY | os.O_TRUNC | os.O_CREATE,
		0600,
	)
	if err != nil {
		util.HandleError(err)
	}
	defer file.Close()

	// Write bytes to file
	_, err = file.Write(data)
	if err != nil {
		util.HandleError(err)
	}
}
