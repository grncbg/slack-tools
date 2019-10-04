package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var channel string
var user string
var list string

// inviteCmd represents the invite command
var inviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invite user to channel",
	Long:  `Invite user to channel.`,
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		client := slack.New(token)

		if user != "" {
			err := inviteToChannel(client, channel, user)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		} else {
			var (
				f   *os.File
				err error
			)

			f = os.Stdin
			if list != "-" {
				f, err = os.Open(list)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
				defer f.Close()
			}

			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				user = scanner.Text()
				err = inviteToChannel(client, channel, user)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		}
	},
}

func inviteToChannel(client *slack.Client, channelID string, userID string) error {
	_, err := client.InviteUsersToConversation(channelID, userID)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(inviteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inviteCmd.PersistentFlags().String("foo", "", "A help for foo")
	inviteCmd.PersistentFlags().StringVarP(&channel, "channel", "c", "", "channel ID")
	inviteCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "user ID")
	inviteCmd.PersistentFlags().StringVarP(&list, "list", "l", "-", "user list")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inviteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
