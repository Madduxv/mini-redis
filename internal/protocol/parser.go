package protocol

import (
  // "fmt"
  "strings"
)

func ParseRESP(data string) (command string, args []string, err error) {
  input := strings.Split(data, "\n") 
  args = input[1:]

  command = strings.ToUpper(input[0])

  return command, args, nil
}

/*
HSET
843c1744-f6c2-6118-6a62-96ea50c2ea1d
name
Maddux

HSET
843c1744-f6c2-6118-6a62-96ea50c2ea1d
group
Maddux's Group

HSETLIST
843c1744-f6c2-6118-6a62-96ea50c2ea1d
genres
ITALIAN, AMERICAN, JAPANESE

HGET
843c1744-f6c2-6118-6a62-96ea50c2ea1d
name

HGET
843c1744-f6c2-6118-6a62-96ea50c2ea1d
group

HGETLIST
843c1744-f6c2-6118-6a62-96ea50c2ea1d
genres
*/
