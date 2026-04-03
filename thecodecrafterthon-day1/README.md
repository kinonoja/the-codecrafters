# the-codecrafters
My Mission
# CLI Calculator — thecodecrafterthon-day1

A terminal calculator built in Go. Supports addition, subtraction, multiplication, and division. Runs in a loop until you quit.

## How to Run

Make sure you have Go installed (go 1.18+ recommended).

```bash
cd thecodecrafterthon-day1
go run main.go
```

Or build a binary first:

```bash
go build -o calculator main.go
./calculator
```

## Usage

```
  > add 10 4
    ✦ Result: 14

  > sub 20 8
    ✦ Result: 12

  > mul 3 7
    ✦ Result: 21

  > div 10 3
    ✦ Result: 3.333333

  > add -10 -5
    ✦ Result: -15

  > help      — shows all commands
  > quit      — exits the calculator
```

## Supported Commands

| Command       | Description              |
|---------------|--------------------------|
| `add <a> <b>` | Addition                 |
| `sub <a> <b>` | Subtraction              |
| `mul <a> <b>` | Multiplication           |
| `div <a> <b>` | Division                 |
| `help`        | Show usage               |
| `quit`        | Exit                     |

## Error Handling

- Division by zero → clear error message, no crash
- Letters instead of numbers → clear error message, no crash  
- Wrong number of arguments → usage hint shown
- Unknown commands → suggests typing `help`
- Negative numbers → fully supported
