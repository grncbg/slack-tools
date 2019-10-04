package cmd

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// idsCmd represents the ids command
var idsCmd = &cobra.Command{
	Use:   "ids",
	Short: "show ids",
	Long:  `show ids`,
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		client := slack.New(token)

		users, err := client.GetUsers()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}

		for _, u := range users {
			if bots || (!(u.IsBot || u.IsAppUser) && u.Name != "slackbot") {
				fmt.Println(u.ID)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(idsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// idsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// idsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	idsCmd.Flags().BoolVarP(&bots, "bots", "b", false, "with bots")
}
