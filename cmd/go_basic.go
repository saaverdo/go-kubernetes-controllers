package cmd

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Run golang basic code",
	Run: func(cmd *cobra.Command, args []string) {
		//Go basic code to run functions
		log.Debug().Msg("Initializing k8s cluster structure")
		k8s := Kubernetes{
			Name:    "k8s-demo-cluster",
			Version: "1.31",
			Users:   []string{"arthas", "illidan"},
			NodeNumber: func() int {
				return 10
			},
		}
		log.Debug().Msg("k8s cluster structure initialized")
		//print users
		k8s.GetUsers()

		//add new user to struct
		k8s.AddNewUser("nameless_one")

		//print users one more time
		k8s.GetUsers()
	},
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	rootCmd.AddCommand(startCmd)

}

// My go basic fucntions here
type Kubernetes struct {
	Name       string     `json:"name"`
	Version    string     `json:"version"`
	Users      []string   `json:"users,omitempty"`
	NodeNumber func() int `json:"-"`
}

func (k8s Kubernetes) GetUsers() {
	log.Debug().Msg("Printing users")
	for _, user := range k8s.Users {
		fmt.Println(user)
	}
}

func (k8s *Kubernetes) AddNewUser(user string) {
	log.Debug().Msgf("Adding user %s", user)
	k8s.Users = append(k8s.Users, user)
}
