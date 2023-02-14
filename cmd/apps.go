package cmd

import (
	"fmt"
	"os"
	"squarecloud/internal/rest"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var applistCmd = &cobra.Command{
	Use:   "apps",
	Short: "See all your active apps on Square Cloud",
	Run: func(cmd *cobra.Command, args []string) {
		user, err := rest.GetSelfUser()
		if err != nil {
			cmd.PrintErrln(err.Error())
			return
		}
		if len(user.Applications) < 1 {
			fmt.Println("You does not have any application active")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)
		defer w.Flush()

		tags := []string{"TAG", "App ID", "MEMORY", "CLUSTER", "LANG", "TYPE", "WEBSITE"}
		fmt.Fprintln(w, strings.Join(tags, " \t "))

		for _, app := range user.Applications {
			values := []string{
				app.Tag,
				app.ID[:12],
				strconv.Itoa(app.RAM) + "mb",
				app.Cluster,
				app.Lang,
				app.Type,
				strconv.FormatBool(app.IsWesite),
			}
			fmt.Fprintln(w, strings.Join(values, " \t "))
		}
	},
}

func init() {
	rootCmd.AddCommand(applistCmd)
}
