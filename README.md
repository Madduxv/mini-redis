# mini-redis
**This is a mini redis server created in Go.**
To interract with the server, you can do so with netcat:
```
nc 127.0.0.1 6379
```
> [!WARNING]
> The server will only read 1024 bytes for each command

## Commands and Responses
**Command Format: CMD\nitem1\n...\nitemn** 

### HSET
- Format: HSET\nKey\nField\nValue
```netcat
HSET
843c1744-f6c2-6118-6a62-96ea50c2ea1d
name
Maddux
```

```netcat
HSET
843c1744-f6c2-6118-6a62-96ea50c2ea1d
group
Maddux's Group
```

### HGET
- Format: HGET\nKey\nField
```netcat
# example 1: Get Session Name
HGET
843c1744-f6c2-6118-6a62-96ea50c2ea1d
name
```

```netcat
# example 2: Get Session Group
HGET
843c1744-f6c2-6118-6a62-96ea50c2ea1d
group
```

### HSETLIST
- Format: HSETLIST\nKey\nField\nCSVs
```netcat
# example 1: Set Session Requested Genres
HSETLIST
843c1744-f6c2-6118-6a62-96ea50c2ea1d
genres
ITALIAN, AMERICAN, JAPANESE
```
```
# example response:
OK
```

### HGETLIST
- Format: HGETLIST\nKey\nField
```netcat
# example 1: Get Session Requested Genres
HGETLIST
843c1744-f6c2-6118-6a62-96ea50c2ea1d
genres
```
```
# example response:
ITALIAN, AMERICAN, JAPANESE
```

### HREMOVELISTFIELD
- Format: HREMOVELISTFIELD\nKey\nField
```netcat
# example 1: Remove Session Requested Genres
HREMOVELIST
843c1744-f6c2-6118-6a62-96ea50c2ea1d
genres
```
```
# example response:
OK
```

### HREMOVESTRINGFIELD
- Format: HREMOVESTRINGFIELD\nKey\nField
```netcat
# example 1: Remove Session Name
HREMOVESTRINGFIELD
843c1744-f6c2-6118-6a62-96ea50c2ea1d
name
```
```
# example response:
OK
```

### HREMOVE
- Format: HREMOVE\nKey\nField
```netcat
# example 1: Remove Session
HREMOVELIST
843c1744-f6c2-6118-6a62-96ea50c2ea1d
```
```
# example response:
OK
```
