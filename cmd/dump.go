package cmd

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/alecthomas/kingpin.v2"

	cmdenv "github.com/previousnext/mysql-toolkit/cmd/env"
	"github.com/previousnext/mysql-toolkit/dumper"
)

type cmdDump struct {
	params dumper.RunParams
}

func (cmd *cmdDump) run(c *kingpin.ParseContext) error {
	return dumper.Dump(os.Stdout, cmd.params)
}

// Dump declares the "dump" subcommand.
func Dump(app *kingpin.Application) {
	c := new(cmdDump)

	cmd := app.Command("dump", "Dump the database").Action(c.run)

	cmd.Flag("mysql-hostname", "Hostname for connecting to Mysql").Required().Envar(cmdenv.MySQLHostname).StringVar(&c.params.Connection.Hostname)
	cmd.Flag("mysql-username", "Username for connecting to Mysql").Required().Envar(cmdenv.MySQLUsername).StringVar(&c.params.Connection.Username)
	cmd.Flag("mysql-password", "Password for connecting to Mysql").Required().Envar(cmdenv.MySQLPassword).StringVar(&c.params.Connection.Password)
	cmd.Flag("mysql-database", "Database for connecting to Mysql").Required().Envar(cmdenv.MySQLDatabase).StringVar(&c.params.Connection.Database)
	cmd.Flag("mysql-protocol", "Protocol for connecting to Mysql").Default("tcp").Envar(cmdenv.MySQLProtocol).StringVar(&c.params.Connection.Protocol)
	cmd.Flag("mysql-port", "Port for connecting to Mysql").Default("3306").Envar(cmdenv.MySQLPort).StringVar(&c.params.Connection.Port)
	cmd.Flag("mysql-max-conn", "Maximum amount of open connections").Default("50").Envar(cmdenv.MySQLMaxConn).IntVar(&c.params.Connection.MaxConn)
	cmd.Flag("mysql-config", "Policy for dumping the database").Default("config.yml").Envar(cmdenv.MySQLConfig).StringVar(&c.params.Config)

	cmd.Flag("file", "Location to save the dumped database").Required().Envar(cmdenv.MySQLFile).StringVar(&c.params.File)
}