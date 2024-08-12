package protocol

import (
	"fmt"
	"strings"
)

func ParseRESP(data []byte) (command string, args []string, err error) {
    input := strings.Split(string(data), "\n") 
    args = input[1:]
    fmt.Printf("Args: %s \r\n", args)

    command = string(strings.ToUpper(input[0]))
    fmt.Printf("Command: %s \r\n", command)

    return command, args, nil
}

/*
HSET
843c1744-f6c2-6118-6a62-96ea50c2ea1d
name
Maddux
group
Maddux's Group

*/
