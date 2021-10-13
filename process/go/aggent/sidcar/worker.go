package sidcar

import (
	"context"
	"github.com/spf13/cobra"
	"hb_process/utils"
	"hb_process/worker"
	"os"
	"os/signal"
	"syscall"
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
	side  := worker.NewSidecar(context.TODO())

	isSignalExit := false

	// hook exit signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGSTOP, syscall.SIGHUP, syscall.SIGUSR1)

	go func() {
		<-sigs
		isSignalExit = true
		side.Close()
		os.Exit(0)
	}()

	err := side.Init()
	if err != nil {
		utils.PanicError(err)
	}

	err = side.Server()

	if err != nil  && isSignalExit{
		<-make(chan interface{})
	}
}
