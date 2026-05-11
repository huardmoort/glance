package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nicolasduteil/glance/internal/config"
	"github.com/nicolasduteil/glance/internal/server"
)

const (
	defaultConfigPath = "glance.yml"
	version           = "0.1.0"
)

func main() {
	var (
		configPath  string
		showVersion bool
		showHelp    bool
	)

	flag.StringVar(&configPath, "config", defaultConfigPath, "Path to the configuration file")
	flag.StringVar(&configPath, "c", defaultConfigPath, "Path to the configuration file (shorthand)")
	flag.BoolVar(&showVersion, "version", false, "Print version information and exit")
	flag.BoolVar(&showVersion, "v", false, "Print version information and exit (shorthand)")
	flag.BoolVar(&showHelp, "help", false, "Show help information")
	flag.BoolVar(&showHelp, "h", false, "Show help information (shorthand)")
	flag.Parse()

	if showHelp {
		fmt.Fprintf(os.Stdout, "Glance - A self-hosted dashboard\n\n")
		fmt.Fprintf(os.Stdout, "Usage:\n  glance [flags]\n\nFlags:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if showVersion {
		fmt.Fprintf(os.Stdout, "glance version %s\n", version)
		os.Exit(0)
	}

	// Load configuration from file
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration from %q: %v", configPath, err)
	}

	log.Printf("Starting glance v%s", version)
	log.Printf("Loaded configuration from %s", configPath)

	// Initialize and start the HTTP server
	srv, err := server.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	if err := srv.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
