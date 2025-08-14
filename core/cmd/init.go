package cmd

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/faizisyellow/falcon/internal/generate"
	"github.com/faizisyellow/falcon/internal/tools"
	"github.com/faizisyellow/falcon/internal/tui"
	"github.com/faizisyellow/falcon/internal/utils"
	"github.com/spf13/cobra"
)

type Options struct {
	Router string
	Db     string
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init ",
	Short: "Create new project",

	RunE: func(cmd *cobra.Command, args []string) error {

		_, err := utils.IsModuleExist()
		if err != nil {
			return err
		}

		options := &Options{}

		withDefault, err := cmd.Flags().GetBool("yes")
		if err != nil {
			return err
		}

		if withDefault {
			options.Router = "Chi"
			options.Db = "Mysql"

		} else {

			p := tea.NewProgram(tui.ModelResult{})

			// Run returns the model as a tea.Model.
			m, err := p.Run()
			if err != nil {
				log.Fatal(err)
			}

			mr, ok := m.(tui.ModelResult)
			if !ok || len(mr.Choices) < 2 {
				fmt.Println("Canceled")
				return nil
			}

			options.Router = mr.Choices[0]
			options.Db = mr.Choices[1]

		}

		if !tools.HasCommand("migrate") {
			return fmt.Errorf("falcon need golang migrate version 4.18.3")
		}

		err = tools.InstallGoTool("go install github.com/air-verse/air@latest", "air")
		if err != nil {
			return fmt.Errorf("failed to install air hot reload: %w", err)
		}

		err = tools.InstallGoTool("github.com/swaggo/swag/cmd/swag@latest", "swag")
		if err != nil {
			return fmt.Errorf("failed to install swag: %w", err)
		}

		err = generate.GenerateNewProject(generate.Options{
			Db:     strings.ToLower(options.Db),
			Router: strings.ToLower(options.Router),
		})
		if err != nil {
			return err
		}

		fmt.Println(
			lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).SetString("\nDone. Now run: \ngo mod tidy\n").String())

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
