package shared

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func IOClear() {
	fmt.Printf("\x1bc")
}

func IOStdinRead(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	key, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	paramSub := strings.Trim(key, "\n")
	return paramSub
}
