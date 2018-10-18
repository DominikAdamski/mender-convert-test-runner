package main

import (
	"fmt"
	"os/exec"
	"time"

)

func updateIntegrationRepo(conf *config) error {
	gitcmd := exec.Command("git", "pull", "--rebase", "origin")
	gitcmd.Dir = conf.integrationDirectory

	// timeout and kill process after GIT_OPERATION_TIMEOUT seconds
	t := time.AfterFunc(GIT_OPERATION_TIMEOUT*time.Second, func() {
		gitcmd.Process.Kill()
	})
	defer t.Stop()

	if err := gitcmd.Run(); err != nil {
		return fmt.Errorf("failed to 'git pull' integration folder: %s\n", err.Error())
	}
	return nil
}
