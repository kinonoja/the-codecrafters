package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a, b float64) float64 { return a + b }
func sub(a, b float64) float64 { return a - b }
func mul(a, b float64) float64 { return a * b }

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func printHelp() {
	fmt.Println()
	fmt.Println("                                           ")
	fmt.Println("             CLI CALCULATOR — HELP         ")
	fmt.Println("  ─────────────────────────────────────────")
	fmt.Println("    add <a> <b>   →  addition              ")
	fmt.Println("    sub <a> <b>   →  subtraction           ")
	fmt.Println("    mul <a> <b>   →  multiplication        ")
	fmt.Println("    div <a> <b>   →  division              ")
	fmt.Println("    help          →  show this menu        ")
	fmt.Println("    quit          →  exit the calculator   ")
	fmt.Println("  ─────────────────────────────────────────")
	fmt.Println("    Negative numbers are supported.        ")
	fmt.Println("     Example:  add -10 -5  →  -15          ")
	fmt.Println("                                           ")
	fmt.Println()
}

func parseArgs(parts []string) (float64, float64, error) {
	if len(parts) < 3 {
		return 0, 0, fmt.Errorf("not enough arguments — usage: %s <a> <b>", parts[0])
	}
	if len(parts) > 3 {
		return 0, 0, fmt.Errorf("too many arguments — usage: %s <a> <b>", parts[0])
	}

	a, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("'%s' is not a valid number", parts[1])
	}

	b, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("'%s' is not a valid number", parts[2])
	}

	return a, b, nil
}

func formatResult(result float64) string {
	// Print as integer if there's no fractional part
	if result == float64(int64(result)) {
		return fmt.Sprintf("%d", int64(result))
	}
	return fmt.Sprintf("%g", result)
}

func main() {
	fmt.Println()
	fmt.Println("  ✦ CLI Calculator — type 'help' to get started")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("  > ")
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		command := strings.ToLower(parts[0])

		switch command {
		case "quit", "exit":
			fmt.Println()
			fmt.Println("  ✦ Goodbye. Happy calculating!")
			fmt.Println()
			os.Exit(0)

		case "help":
			printHelp()

		case "add", "sub", "mul", "div":
			a, b, err := parseArgs(parts)
			if err != nil {
				fmt.Printf("  ✗ Error: %s\n\n", err)
				continue
			}

			var result float64

			switch command {
			case "add":
				result = add(a, b)
			case "sub":
				result = sub(a, b)
			case "mul":
				result = mul(a, b)
			case "div":
				result, err = div(a, b)
				if err != nil {
					fmt.Printf("  ✗ Error: %s\n\n", err)
					continue
				}
			}

			fmt.Printf("  ✦ Result: %s\n\n", formatResult(result))

		default:
			fmt.Printf("  ✗ Unknown command '%s'. Type 'help' to see available commands.\n\n", command)
		}
	}
}
