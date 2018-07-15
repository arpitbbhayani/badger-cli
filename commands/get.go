package commands

import (
	"bufio"
	"os"

	"github.com/arpitbbhayani/badger-cli/sessions"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var keysFile string

func init() {
	getCommand.Flags().StringVarP(&databaseDir, "dir", "", "", "Path to Database Directory (required)")
	getCommand.MarkFlagRequired("dir")

	getCommand.Flags().StringVarP(&keysFile, "keys-file", "", "", "Path to Keys File")

	BadgerCommand.AddCommand(getCommand)
}

var getCommand = &cobra.Command{
	Use:   "get",
	Short: "Gets a key",
	Run: func(cmd *cobra.Command, args []string) {
		sessions.InitBadgerDB("session", databaseDir)

		db := sessions.BadgerDBSessions["session"]
		txn := db.NewTransaction(false)

		keys := args

		if keysFile != "" {
			fp, err := os.Open(keysFile)
			if err != nil {
				logrus.Errorf("error opening file %q: %v", keysFile, err)
			}
			defer fp.Close()
			scanner := bufio.NewScanner(fp)
			scanner.Split(bufio.ScanLines)

			for scanner.Scan() {
				keys = append(keys, scanner.Text())
			}
		}

		for _, key := range keys {
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
