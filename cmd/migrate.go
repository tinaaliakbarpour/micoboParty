package cmd

import (
	"fmt"
	"html"
	"micobianParty/client/postgres"
	"micobianParty/config"
	migrate "micobianParty/pkg/migrations"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var migrateCMD = cobra.Command{
	Use:     "migrate",
	Long:    "migrate database strucutures. This will migrate tables",
	Aliases: []string{"m"},
	Run:     Runner.Migrate,
}

// migrate database with fake data
func (c *command) Migrate(cmd *cobra.Command, args []string) {

	if err := config.Confs.Load(cPath); err != nil {
		fmt.Printf("%v %v\n", aurora.White(html.UnescapeString("&#x274C;")), err)
		return
	}
	if err := postgres.Storage.Connect(config.Confs); err != nil {
		fmt.Printf("%v error in connect to postgress: %v\n", aurora.White(html.UnescapeString("&#x274C;")), err)
		return
	}

	err := migrate.AutoMigrateDB()
	if err != nil {
		fmt.Printf("%v error in migration: %v\n", aurora.White(html.UnescapeString("&#x274C;")), err)
		return
	}
}
