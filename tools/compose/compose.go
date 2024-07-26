package compose

import (
	"context"
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/opannapo/pubsub-emulator-tools/shared"
)

func Setup(ctx context.Context, fs embed.FS) {
	shared.IOClear()
	projectID := shared.IOStdinRead("[setup-emulator-compose] project ID  : ")
	port := shared.IOStdinRead("[setup-emulator-compose] port  : ")

	result := shared.CreateDockerCompose(projectID, port, fs)
	fmt.Println("Docker Compose created at ", result)
	targetFolder := filepath.Dir(result)

	cmd := exec.Command("docker", "compose", "up", "-d")
	cmd.WaitDelay = 500
	cmd.Dir = targetFolder
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
}
