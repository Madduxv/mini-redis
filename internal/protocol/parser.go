package protocol

import ("fmt"; "bytes")

/* RESP:
   *x-> An array with x elements
   $x-> A bulk string with x elements
*/

/* example Requests:
SET: {
  *3\r\n$3\r\nSET\r\n$4\r\nmykey\r\n$13\r\nHello, World!\r\n
}

GET: {
  *2\r\n$3r\nGET\r\n$4\r\nmykey\r\n
}

DEL: {
  *2\r\n$3\r\nDEL\r\n$4\r\nmykey\r\n
}

*/

type Request struct {
  requestType byte
  requestLength byte
  requestBytes byte
}

func ParseRESP(data []byte) (command string, args []string, err error) {
    if len(data) == 0 || data[0] != '*' {
        return "", nil, fmt.Errorf("invalid RESP format")
    }
    data_split := bytes.Split(data, []byte("\r\n")) 
    fmt.Print(data_split)

    // Parse number of elements in the array
    // Your parsing logic here...

    // Example: Convert a bulk string part to a string
    command = string(data[3:6]) // This would extract "SET" from a properly formatted input

    return command, args, nil
}
