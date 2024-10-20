package pod

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mrangelba/solid_cli/cmd/models"
	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	listCmd = &cobra.Command{
		Use:   "ls",
		Short: "List all SOLID Pods",
		Long:  `List all SOLID Pods.`,
		Run: func(cmd *cobra.Command, args []string) {
			entries, err := os.ReadDir("/data/.internal/accounts/data")

			if err != nil {
				log.Fatal(err)
			}

			rows := make([]table.Row, 0)

			for _, entry := range entries {
				accountData := models.AccountData{}

				dataFile, err := os.Open("/data/.internal/accounts/data/" + entry.Name())

				row := table.Row{}

				if err != nil {
					log.Fatal(err)
				}
				defer dataFile.Close()

				bytes, err := io.ReadAll(dataFile)

				if err != nil {
					log.Fatal(err)
				}

				err = json.Unmarshal(bytes, &accountData)

				if err != nil {
					log.Fatal(err)
				}

				for _, pod := range accountData.Payload.Pod {
					row = append(row, len(rows)+1)
					row = append(row, pod.ID)
					row = append(row, pod.BaseURL)

					owner := ""
					for _, password := range accountData.Payload.Password {
						if owner != "" {
							owner += "\n"
						}

						owner = password.Email
					}
					row = append(row, owner)
				}
				rows = append(rows, row)
			}

			tw := table.NewWriter()
			tw.AppendHeader(table.Row{"#", "Pod ID", "Pod URL", "Owner"})
			tw.AppendRows(rows)
			tw.SetIndexColumn(1)
			tw.SetTitle("Listing all SOLID Pods")
			fmt.Println(tw.Render())
		},
	}
)

func init() {
	PodCmd.AddCommand(listCmd)
}
