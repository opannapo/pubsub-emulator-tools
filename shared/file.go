package shared

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func CreateDockerCompose(projectID, port string) (result string) {
	type DockerCompose struct {
		Image         string
		ContainerName string
		Port          string
		ProjectID     string
	}

	data := DockerCompose{
		Image:         "google/cloud-sdk:emulators",
		ContainerName: "pubsub-emulator",
		Port:          port,
		ProjectID:     projectID,
	}

	tmpl, err := template.ParseFiles("./emulator/docker-compose.tmpl")
	if err != nil {
		panic(err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	nDir := filepath.Join(homeDir, ".tmp-emulator")
	nFile := filepath.Join(nDir, "docker-compose.yml")

	err = os.MkdirAll(nDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating new folder:", err)
		return
	}

	file, err := os.Create(nFile)
	if err != nil {
		fmt.Println("Error creating new file:", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		panic(err)
	}

	return file.Name()
}
