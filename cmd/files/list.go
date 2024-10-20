package files

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/mrangelba/solid_cli/cmd/models"
	"github.com/spf13/cobra"
)

var (
	email string

	listCmd = &cobra.Command{
		Use:   "ls",
		Short: "List all SOLID Files",
		Long:  `List all SOLID Files.`,
	}
)

func init() {
	FilesCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&email, "email", "e", "", "Email")

	listCmd.Run = func(cmd *cobra.Command, args []string) {
		if email == "" {
			cmd.Help()
			return
		}

		listFilesByEmail(email)
	}
}

func listFilesByEmail(email string) {
	accountEmailDataPath := fmt.Sprintf("/data/.internal/accounts/index/password/email/%s$.json", email)

	_, err := os.Stat(accountEmailDataPath)

	if os.IsNotExist(err) {
		fmt.Println("Account not found")
		return
	}

	emailFile, err := os.Open(accountEmailDataPath)

	if err != nil {
		return
	}
	defer emailFile.Close()

	passwordData := &models.PasswordData{}

	bytes, err := io.ReadAll(emailFile)

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(bytes, &passwordData)

	if err != nil {
		fmt.Println(err)
	}

	accountDataFilePath := fmt.Sprintf("/data/.internal/accounts/data/%s$.json", passwordData.Payload[0])
	_, err = os.Stat(accountDataFilePath)

	if os.IsNotExist(err) {
		fmt.Println("Account not found")
		return
	}

	accountFile, err := os.Open(accountDataFilePath)

	if err != nil {
		return
	}
	defer accountFile.Close()

	accountData := &models.AccountData{}

	bytes, err = io.ReadAll(accountFile)

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(bytes, &accountData)

	if err != nil {
		fmt.Println(err)
	}

	// password
	for _, password := range accountData.Payload.Password {
		passwordPath := fmt.Sprintf("/data/.internal/accounts/index/password/%s$.json", password.ID)

		printFilePath(passwordPath)

		passwordEmailPath := fmt.Sprintf("/data/.internal/accounts/index/password/email/%s$.json", password.Email)

		printFilePath(passwordEmailPath)
	}

	// client credentials
	for _, client := range accountData.Payload.ClientCredentials {
		clientCredentialsPath := fmt.Sprintf("/data/.internal/accounts/index/clientCredentials/%s$.json", client.ID)

		printFilePath(clientCredentialsPath)

		clientCredentialsLabelPath := fmt.Sprintf("/data/.internal/accounts/index/clientCredentials/label/%s$.json", client.Label)

		printFilePath(clientCredentialsLabelPath)
	}

	// pod
	for _, pod := range accountData.Payload.Pod {
		podPath := fmt.Sprintf("/data/.internal/accounts/index/pod/%s$.json", pod.ID)

		printFilePath(podPath)

		podBaseUrlPath := fmt.Sprintf("/data/.internal/accounts/index/pod/baseUrl/%s$.json", url.QueryEscape(pod.BaseURL))

		printFilePath(podBaseUrlPath)

		for _, owner := range pod.Owner {
			webIdLinkPath := fmt.Sprintf("/data/.internal/accounts/index/webIdLink/%s$.json", owner.ID)

			printFilePath(webIdLinkPath)

			webIdPath := fmt.Sprintf("/data/.internal/accounts/index/webIdLink/webid/%s$.json", url.QueryEscape(owner.WebID))

			printFilePath(webIdPath)
		}

		printPodFolder(pod.BaseURL)
	}

	printFilePath(accountDataFilePath)
}

func printFilePath(filePath string) {
	fmt.Printf("%s\n", filePath)
}

func printPodFolder(podBaseURL string) {
	uri, err := url.Parse(podBaseURL)

	if err != nil {
		fmt.Println(err)
		return
	}
	podID := strings.ReplaceAll(uri.Path, "/", "")
	podFolderPath := fmt.Sprintf("/data/%s", podID)

	printFolderPath(podFolderPath)
}

func printFolderPath(folderPath string) {
	entries, err := os.ReadDir(folderPath)

	if err != nil {
		fmt.Println(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			printFolderPath(fmt.Sprintf("%s/%s", folderPath, entry.Name()))
		} else {
			fmt.Printf("%s/%s\n", folderPath, entry.Name())
		}
	}
}
