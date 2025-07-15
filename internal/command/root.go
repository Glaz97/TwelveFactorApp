package command

import (
	"github.com/Glaz97/twelvefactorapp/internal/app"
	"github.com/Glaz97/twelvefactorapp/internal/server/server_http"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func GetRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "twelvefactorapp",
		Short: "twelvefactorapp is a test service",
		Run: func(cmd *cobra.Command, args []string) {
			serve()
		},
	}

	return cmd
}

func serve() {
	fx.New(
		app.Module,
		fx.Invoke(func(_ *server_http.Server) {}),
	).Run()
}
