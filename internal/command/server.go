package cmd

import (
	"context"
	"log"

	"github.com/farhaniupr/dating-api/internal/middleware"
	"github.com/farhaniupr/dating-api/internal/routes"
	"github.com/farhaniupr/dating-api/package/library"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s", "sr", "srg", "sg"},
	Short:   "run server",
	Long:    "Start running the server",
	Run:     server,
}

func init() {
	RootCmd.AddCommand(serverCmd)
	RootCmd.AddCommand(migrateCmd)
}

func server(cmd *cobra.Command, args []string) {
	opts := fx.Options(
		fx.Invoke(func(
			route routes.Routes,
			router library.RequestHandler,
			middleware middleware.Middlewares,
			env library.Env,

		) {

			route.Setup()
			middleware.Setup()

		}),
	)
	ctx := context.Background()
	app := fx.New(CommonModules, opts)
	err := app.Start(ctx)
	defer app.Stop(ctx)
	if err != nil {
		log.Println(err.Error())
	}
}
