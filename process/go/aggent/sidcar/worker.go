package sidcar

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewWorkerSidecarCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "worker",
		Short: "start worker sidecar",
		Run: func(cmd *cobra.Command, args []string) {
			runWorkerSidecar()
		},
	}
	return cmd
}

func runWorkerSidecar() {
	fmt.Println(1111)
}
