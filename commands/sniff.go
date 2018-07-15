package commands

import (
	"strconv"

	"github.com/arpitbbhayani/badger-cli/sessions"
	"github.com/dgraph-io/badger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var count string

func init() {
	sniffCommand.Flags().StringVarP(&databaseDir, "dir", "", "", "Path to Database Directory (required)")
	sniffCommand.Flags().StringVarP(&count, "count", "", "", "Number of items to sniff (default 10)")
	sniffCommand.MarkFlagRequired("dir")

	BadgerCommand.AddCommand(sniffCommand)
}

var sniffCommand = &cobra.Command{
	Use:   "sniff",
	Short: "sniffs items",
	Run: func(cmd *cobra.Command, args []string) {
		sessions.InitBadgerDB("session", databaseDir)

		var err error
		db := sessions.BadgerDBSessions["session"]

		sniffCount, err := strconv.Atoi(count)
		if err != nil {
			sniffCount = 10
		}

		txn := db.NewTransaction(false)

		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid() && sniffCount != 0; it.Next() {
			item := it.Item()

			k := item.Key()
			v, err := item.Value()
			if err != nil {
				logrus.Error(err)
			}

			logrus.Infof("%s -> %s", k, v)

			sniffCount--
		}

		txn.Discard()
	},
}
