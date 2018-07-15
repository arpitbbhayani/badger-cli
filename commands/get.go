package commands

import (
	"errors"

	"github.com/arpitbbhayani/badger-cli/sessions"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	getCommand.Flags().StringVarP(&databaseDir, "dir", "", "", "Path to Database Directory (required)")
	getCommand.MarkFlagRequired("dir")

	BadgerCommand.AddCommand(getCommand)
}

var getCommand = &cobra.Command{
	Use:   "get",
	Short: "Gets a key",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one key")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		sessions.InitBadgerDB("session", databaseDir)

		db := sessions.BadgerDBSessions["session"]
		txn := db.NewTransaction(false)

		for _, key := range args {
			item, err := txn.Get([]byte(key))
			if err != nil {
				logrus.Errorf("error while fetching key: '%s', error: %q", key, err)
				continue
			}
			value, err := item.Value()
			if err != nil {
				logrus.Errorf("error while fetching value for key: '%s', error: %q", key, err)
			}
			logrus.Infof("%s -> %s", key, value)
		}
	},
}
