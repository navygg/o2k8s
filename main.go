package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

// Config config.toml
type Config struct {
	APPID   string
	Host    string
	Port    int
	LogDir  string
	LogFile string
}

// start start http server
func start(fileName string) error {
	var (
		err    error
		config Config
	)

	if _, err = toml.DecodeFile(fileName, &config); err != nil {
		log.Fatalf("config file parse error: %v", err)
	}

	config.LogDir, err = filepath.Abs(path.Clean(config.LogDir))
	if err != nil {
		log.Fatalf("abs error: %v", err)
	}

	if err := os.MkdirAll(config.LogDir, 0755); err != nil {
		log.Fatalf("mkdirall error: %v", err)
	}

	if config.LogFile == "" {
		log.Fatalf("LogFile is nil")
	}

	logFileName := path.Join(config.LogDir, path.Clean(config.LogFile))
	logFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("open %s error: %v", logFileName, err)
	}
	defer logFile.Close()
	logger := log.New(logFile, config.APPID+" ", log.LstdFlags|log.Lshortfile)

	hiHandler := HiHandler{logger: logger}
	http.HandleFunc("/hi", hiHandler.ServeHTTP)

	server := http.Server{
		Addr: fmt.Sprintf("%v:%v", config.Host, config.Port),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("start server error: %v", err)
		}
	}()

	// timeout context for shutdown
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	for s := range c {
		log.Printf("got signal: %v", s)
		server.Shutdown(ctx)
		log.Printf("graceful shutdown")
		return nil
	}

	return nil
}

func main() {
	cmdRun := &cobra.Command{
		Use:   "run config.toml",
		Short: "run server",
		Long:  "run server",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("need config file")
			}
			if err := start(args[0]); err != nil {
				return fmt.Errorf("start server error: %+v", err)
			}
			return nil
		},
	}

	rootCmd := &cobra.Command{
		Use:  "o2k8s server",
		Long: "o2k8s server",
	}

	rootCmd.AddCommand(cmdRun)
	rootCmd.Execute()
}
