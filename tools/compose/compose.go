package compose

import (
	"context"
	"fmt"
	"github.com/opannapo/pubsub-emulator-tools/shared"
	"os"
	"os/exec"
	"path/filepath"
)

func Setup(ctx context.Context) {
	shared.IOClear()
	projectID := shared.IOStdinRead("[setup-emulator-compose] project ID  : ")
	port := shared.IOStdinRead("[setup-emulator-compose] port  : ")

	result := shared.CreateDockerCompose(projectID, port)
	fmt.Println("Docker Compose created at ", result)

	//Run compose
	targetFolder := filepath.Dir(result)

	cmd := exec.Command("docker", "compose", "up", "-d")
	cmd.Dir = targetFolder
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Println(string(output))
}
