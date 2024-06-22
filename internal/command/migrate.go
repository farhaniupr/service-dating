package cmd

import (
	"context"
	"log"

	"github.com/farhaniupr/dating-api/package/library"
	"github.com/farhaniupr/dating-api/resource/model"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Aliases: []string{"m"},
	Short:   "run migrate db server",
	Long:    "Start running migrate db server",
	Run:     migrate,
}

func migrate(cmd *cobra.Command, args []string) {
	opts := fx.Options(
		fx.Invoke(func(
			env library.Env,
			db library.Database,
		) {
			if err := db.MysqlDB.WithContext(context.Background()).
				AutoMigrate(&model.User{}, &model.UserSetting{}, &model.UserLiked{}); err != nil {
				library.Writelog(context.Background(), env, "err", err.Error())
			}
		}),
	)
	ctx := context.Background()
	app := fx.New(MigrateModules, opts)
	err := app.Start(ctx)
	defer app.Stop(ctx)
	if err != nil {
		log.Println(err.Error())
	}
}
