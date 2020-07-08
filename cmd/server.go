package cmd

import (
	"deck-api/api"
	"deck-api/service/deck"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
)

func Server() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Starts http server",
		Run: func(cmd *cobra.Command, args []string) {
			port, err := cmd.Flags().GetInt("port")
			if err != nil {
				fmt.Printf("error getting port flag: %s", err)
				os.Exit(1)
			}

			s := &http.Server{
				Addr:        fmt.Sprintf(":%d", port),
				IdleTimeout: 30 * time.Second,
				Handler:     api.New(deck.NewMemory()),
			}

			c := make(chan os.Signal)
			signal.Notify(c, os.Interrupt)

			go func() {
				<-c
				fmt.Printf("Shutting down server....\n")
				s.Close()
				os.Exit(0)
			}()

			fmt.Printf("Server running on %s\n", s.Addr)
			fmt.Printf("Server error: %s\n", s.ListenAndServe())
		},
	}

	cmd.Flags().IntP("port", "p", 9000, "server port")

	return cmd
}
