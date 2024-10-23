package account

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
	accountID string
	email     string
	webid     string

	deleteCmd = &cobra.Command{
		Use:   "rm",
		Short: "Delete a SOLID account",
		Long:  `Delete a SOLID account.`,
	}
)

func init() {
	AccountCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&accountID, "id", "i", "", "Account ID")
	deleteCmd.Flags().StringVarP(&email, "email", "e", "", "Email")
	deleteCmd.Flags().StringVarP(&webid, "webid", "w", "", "WebID")

	deleteCmd.Run = func(cmd *cobra.Command, args []string) {
		if accountID == "" && email == "" && webid == "" {
			cmd.Help()
			return
		}

		if accountID != "" {
			deleteAccountByID(accountID)
		}

		if email != "" {
			deleteAccountByEmail(email)
		}

		if webid != "" {
			deleteAccountByWebID(webid)
		}
	}
}

func deleteAccountByID(accountID string) {
	accountDataFilePath := fmt.Sprintf("/data/.internal/accounts/data/%s$.json", accountID)
	_, err := os.Stat(accountDataFilePath)

	if os.IsNotExist(err) {
		fmt.Printf("File not found: %s\n", accountDataFilePath)
		fmt.Println("Account not found")
		return
	}

	accountFile, err := os.Open(accountDataFilePath)

	if err != nil {
		return
	}
	defer accountFile.Close()

	accountData := &models.AccountData{}

	bytes, err := io.ReadAll(accountFile)

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(bytes, &accountData)

	if err != nil {
		fmt.Println(err)
	}

	// delete password
	for _, password := range accountData.Payload.Password {
		passwordPath := fmt.Sprintf("/data/.internal/accounts/index/password/%s$.json", password.ID)

		deleteFile(passwordPath)

		passwordEmailPath := fmt.Sprintf("/data/.internal/accounts/index/password/email/%s$.json", password.Email)

		deleteFile(passwordEmailPath)
	}

	// delete client credentials
	for _, client := range accountData.Payload.ClientCredentials {
		clientCredentialsPath := fmt.Sprintf("/data/.internal/accounts/index/clientCredentials/%s$.json", client.ID)

		deleteFile(clientCredentialsPath)

		clientCredentialsLabelPath := fmt.Sprintf("/data/.internal/accounts/index/clientCredentials/label/%s$.json", client.Label)

		deleteFile(clientCredentialsLabelPath)
	}

	// delete pod
	for _, pod := range accountData.Payload.Pod {
		podPath := fmt.Sprintf("/data/.internal/accounts/index/pod/%s$.json", pod.ID)

		deleteFile(podPath)

		podBaseUrlPath := fmt.Sprintf("/data/.internal/accounts/index/pod/baseUrl/%s$.json", url.QueryEscape(pod.BaseURL))

		deleteFile(podBaseUrlPath)

		for _, owner := range pod.Owner {
			webIdLinkPath := fmt.Sprintf("/data/.internal/accounts/index/webIdLink/%s$.json", owner.ID)

			deleteFile(webIdLinkPath)

			webIdPath := fmt.Sprintf("/data/.internal/accounts/index/webIdLink/webid/%s$.json", strings.ReplaceAll(url.QueryEscape(owner.WebID), "%23", "#"))

			deleteFile(webIdPath)
		}

		deletePodFolder(pod.BaseURL)
	}

	deleteFile(accountDataFilePath)
}

func deleteAccountByEmail(email string) {
	accountEmailDataPath := fmt.Sprintf("/data/.internal/accounts/index/password/email/%s$.json", email)

	_, err := os.Stat(accountEmailDataPath)

	if os.IsNotExist(err) {
		fmt.Printf("File not found: %s\n", accountEmailDataPath)
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

	for _, data := range passwordData.Payload {
		deleteAccountByID(data)
	}
}

func deleteAccountByWebID(webid string) {
	webIdPath := fmt.Sprintf("/data/.internal/accounts/index/webIdLink/webId/%s$.json", strings.ReplaceAll(url.QueryEscape(webid), "%23", "#"))

	_, err := os.Stat(webIdPath)

	if os.IsNotExist(err) {
		fmt.Printf("File not found: %s\n", webIdPath)
		fmt.Println("Account not found")
		return
	}

	webIdFile, err := os.Open(webIdPath)

	if err != nil {
		return
	}
	defer webIdFile.Close()

	webIdLink := &models.PasswordData{}

	bytes, err := io.ReadAll(webIdFile)

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(bytes, &webIdLink)

	if err != nil {
		fmt.Println(err)
	}

	for _, data := range webIdLink.Payload {
		deleteAccountByID(data)
	}
}

func deleteFile(filePath string) {
	fmt.Printf("Deleting file %s\n", filePath)

	err := os.Remove(filePath)

	if !os.IsNotExist(err) {
		if err != nil {
			fmt.Println(err)
		}
	}
}

func deletePodFolder(podUrl string) {
	uri, err := url.Parse(podUrl)

	if err != nil {
		fmt.Println(err)
		return
	}
	podID := strings.ReplaceAll(uri.Path, "/", "")
	podFolderPath := fmt.Sprintf("/data/%s", podID)
	fmt.Printf("Deleting folder %s\n", podFolderPath)

	err = os.RemoveAll(podFolderPath)

	if !os.IsNotExist(err) {
		if err != nil {
			fmt.Println(err)
		}
	}
}
