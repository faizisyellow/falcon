package cmd

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/faizisyellow/falcon/internal/tui"
	"github.com/spf13/cobra"
)

type Options struct {
	Router string
	Db     string
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init {name project}",
	Args:  cobra.ExactArgs(1),
	Short: "Create new project",
	RunE: func(cmd *cobra.Command, args []string) error {

		options := &Options{}

		withDefault, err := cmd.Flags().GetBool("yes")
		if err != nil {
			return err
		}

		if withDefault {
			options.Router = "Echo"
			options.Db = "Mysql"

		} else {

			p := tea.NewProgram(tui.ModelResult{})

			// Run returns the model as a tea.Model.
			m, err := p.Run()
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			// Assert the final tea.Model to our local model and print the choice.
			if m, ok := m.(tui.ModelResult); ok && len(m.Choices) != 0 {
				options.Router = m.Choices[0]
				options.Db = m.Choices[1]
			}

		}

		fmt.Printf("You choose router: %v and db: %v\n", options.Router, options.Db)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().BoolP("yes", "y", false, "Using default options")
}
