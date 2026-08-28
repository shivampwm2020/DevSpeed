package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"devspeed/cli/internal/benchmark"
	"devspeed/cli/internal/system"
)


func main() {
	// Define flags
	versionFlag := flag.Bool("version", false, "Show version information")
	verboseFlag := flag.Bool("verbose", false, "Enable verbose output")
	helpFlag := flag.Bool("help", false, "Show help information")
	
	// We need to parse flags after the command to handle flags correctly
	// Let's use a different approach: parse global flags first, then command-specific flags
	
	// Parse until we find a command
	args := os.Args[1:]
	var globalArgs []string
	command := ""

	// Debug: print all arguments
	fmt.Printf("DEBUG: All args: %v\n", os.Args)
	fmt.Printf("DEBUG: Raw args: %v\n", args)

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			// This is a flag, add to globalArgs
			globalArgs = append(globalArgs, arg)
			fmt.Printf("DEBUG: Found flag: %s\n", arg)
		} else {
			// This is a command or argument
			command = arg
			fmt.Printf("DEBUG: Found command: %s\n", command)
			break
		}
	}
	fmt.Printf("DEBUG: Final globalArgs: %v\n", globalArgs)
	fmt.Printf("DEBUG: Final command: %s\n", command)

	// Parse global flags
	flagSet := flag.NewFlagSet("devspeed", flag.ContinueOnError)
	flagSet.BoolVar(versionFlag, "version", false, "Show version information")
	flagSet.BoolVar(verboseFlag, "verbose", false, "Enable verbose output")
	flagSet.BoolVar(helpFlag, "help", false, "Show help information")
	
	// Debug: show globalArgs before parsing
	fmt.Printf("DEBUG: globalArgs before parsing: %v\n", globalArgs)
	
	err := flagSet.Parse(globalArgs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		return
	}
	
	// Debug: show flags after parsing
	fmt.Printf("DEBUG: Flags after parsing - version=%v, verbose=%v, help=%v\n", *versionFlag, *verboseFlag, *helpFlag)


	// Debug: print parsed flags
	fmt.Printf("DEBUG: Parsed flags - version=%v, verbose=%v, help=%v\n", *versionFlag, *verboseFlag, *helpFlag)

	if *versionFlag {
		fmt.Println("DevSpeed CLI v0.1.0")
		return
	}

	if *helpFlag {
		help()
		return
	}

	// If no command specified, use 'run' as default
	if command == "" {
		command = "run"
	}

	// Handle subcommands
	switch command {
	case "version":
		fmt.Println("ReqBeam CLI v0.1.0")
	case "system":
		systemInfo()
	case "run":
		// Check if verbose flag was set
		if *verboseFlag {
			fmt.Println("Running benchmarks in verbose mode...")
		}
		runBenchmarks(*verboseFlag)
	default:
		help()
	}
}



func runBenchmarks(verbose bool) {
	fmt.Println("⚡ DevSpeed")
	fmt.Println("Developer Machine Benchmark")
	fmt.Println("Suite v0.1.0")
	fmt.Println()

	// Get system information
	info, err := system.GetSystemInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting system information: %v\n", err)
		return
	}

	fmt.Println("Detected system")
	fmt.Println()
	fmt.Printf("%s\n", info.CPUModel)
	fmt.Printf("%d GB RAM\n", info.MemoryBytes/(1024*1024*1024))
	fmt.Printf("%s\n", info.OS)
	fmt.Printf("%s\n", info.Arch)
	fmt.Printf("%s\n", info.OSVersion)
	fmt.Println()

	if verbose {
		fmt.Println("Using temporary directory: /tmp (actual location may vary)")
		fmt.Println()
	}

	fmt.Println("Running benchmarks...")
	fmt.Println()

	// Create a benchmark runner
	// In a real implementation, we would use the system's temporary directory
	runner := benchmark.NewRunner("/tmp")

	// Create the benchmarks to run
	benchmarks := []benchmark.Benchmark{
		&benchmark.FilesystemSmallFiles{},
	}

	// Run the benchmarks
	results := runner.RunBenchmarks(benchmarks)

	// Display results
	for _, result := range results {
		if !result.Success {
			fmt.Printf("%s: FAILED - %s\n", result.Name, result.ErrorMessage)
			continue
		}
		
		// For now, just display a simple success message
		// In a real implementation, we would calculate a score based on the measurements
		fmt.Printf("%s: PASSED\n", result.Name)
		
		// In verbose mode, show some of the measurements
		if verbose {
			fmt.Println("Verbose mode enabled - showing measurements")
			if ops, ok := result.Measurements["createOpsPerSecond"].(float64); ok {
				fmt.Printf("  Create operations per second: %.0f\n", ops)
			} else {
				fmt.Println("  Create operations per second: not available")
			}
			if ops, ok := result.Measurements["readOpsPerSecond"].(float64); ok {
				fmt.Printf("  Read operations per second: %.0f\n", ops)
			} else {
				fmt.Println("  Read operations per second: not available")
			}
		}
	}

	fmt.Println()
	fmt.Println("Overall Score")
	fmt.Println()
	fmt.Println("[Score calculation not yet implemented]")
	fmt.Println()
	fmt.Println("Run:")
	fmt.Println()
	fmt.Println("devspeed doctor")
	fmt.Println()
	fmt.Println("for detailed diagnostics.")
	fmt.Println()
	fmt.Println("Result:")
	fmt.Println("https://reqbeam.dev/r/x7ad92")
}

func help() {
	fmt.Println("Usage: reqbeam [command] [flags]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  reqbeam           Run benchmarks (default)")
	fmt.Println("  version           Show version information")
	fmt.Println("  system            Show system information")
	fmt.Println("  run               Run benchmarks")
	fmt.Println("  doctor            Run diagnosis on results")
	fmt.Println("  compare           Compare with other machines")
	fmt.Println("  share             Share results")
	fmt.Println("  history           Show benchmark history")
	fmt.Println("  repo              Benchmark a repository")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -version          Show version information")
	fmt.Println("  -verbose         Enable verbose output")
	fmt.Println("  -help            Show help information")
}

func systemInfo() {
	info, err := system.GetSystemInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting system information: %v\n", err)
		return
	}

	fmt.Printf("Detected system\n\n")
	fmt.Printf("%s\n", info.CPUModel)
	fmt.Printf("%d GB RAM\n", info.MemoryBytes/(1024*1024*1024))
	fmt.Printf("%s\n", info.OS)
	fmt.Printf("%s\n", info.Arch)
	fmt.Printf("%s\n", info.OSVersion)
}
