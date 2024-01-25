/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mm-bulk-preference-change input-file preference-category preference-name preference-value",
	Short: "Use mmctl to update the user preferences for several users",
	Long: `Update the preferences of several Mattermost users.
Use a file with the emails of several users separated by whitespaces.
Example:
	mm-bulk-preference-change emails.txt favorite_channel ungsg3mx77gkdxqnacbo1aa6de true`,
	RunE: cmdF,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("mmctlPath", "", "Define the mmctl path if not installed globally")
}

func cmdF(cmd *cobra.Command, args []string) error {
	if len(args) != 4 {
		return errors.New("you must use 4 arguments: input file, preference category, preference name and preference value")
	}

	filePath := args[0]
	preferenceCategory := args[1]
	preferenceName := args[2]
	preferenceValue := args[3]

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return errors.Errorf("failed to open the file: %v", err)
	}

	fileString := string(fileBytes)
	userEmails := strings.Fields(fileString)

	mmctlCommandPath := "mmctl"
	if cmd.Flag("mmctlPath").Value.String() != "" {
		mmctlCommandPath = cmd.Flag("mmctlPath").Value.String()
	}
	mmctlArgs := append([]string{"user", "preference", "update", "--category", preferenceCategory, "--name", preferenceName, "--value", preferenceValue}, userEmails...)

	mmctlCommand := exec.Command(mmctlCommandPath, mmctlArgs...)
	mmctlCommand.Stderr = os.Stderr
	mmctlCommand.Stdout = os.Stdout

	err = mmctlCommand.Run()
	if err != nil {
		return errors.Errorf("failed to execute mmctl: %v", err)
	}

	fmt.Println("Preference updated successfully")
	return nil
}
