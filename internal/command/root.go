package cmd

import (
	"log"

	"github.com/farhaniupr/dating-api/internal/controller"
	"github.com/farhaniupr/dating-api/internal/helper"
	"github.com/farhaniupr/dating-api/internal/middleware"
	"github.com/farhaniupr/dating-api/internal/repository"
	"github.com/farhaniupr/dating-api/internal/routes"
	"github.com/farhaniupr/dating-api/internal/service"
	"github.com/farhaniupr/dating-api/package/eksternal"
	"github.com/farhaniupr/dating-api/package/library"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var RootCmd = &cobra.Command{
	Use:   "Introduction",
	Short: "Go Clean Echo",
	Long:  "Go Clean Echo",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Println("err root command execute :", err.Error())
	}
}

var CommonModules = fx.Options(
	library.Module,
	eksternal.Module,
	routes.Module,
	middleware.Module,
	controller.Module,
	service.Module,
	repository.Module,
	helper.Module,
)

var RabbitModules = fx.Options(
	library.Module,
	eksternal.Module,
	helper.Module,
)

var MigrateModules = fx.Options(
	library.Module,
)
