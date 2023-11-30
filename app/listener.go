package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aarsad18/go-supplier-otomax/model"
	"github.com/aarsad18/go-supplier-otomax/resource"
	"github.com/lib/pq"
	"github.com/spf13/viper"
)

type Listener struct {
	Listener *pq.Listener
}

func NewListener(db *resource.DBConn) *Listener {
	listener := pq.NewListener(db.Dsn, 10*time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Printf("Listener error: %s\n", err)
		}
	})
	return &Listener{
		Listener: listener,
	}
}

func (ls *Listener) StartListener(db *resource.DBConn) {
	err := ls.Listener.Listen(viper.GetString("NOTIFY_CHANNEL"))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening for notifications on channel %s...\n", viper.Get("NOTIFY_CHANNEL"))

	// Setup a signal handler to gracefully shutdown the listener
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-sig:
			log.Println("Received signal to shutdown. Closing listener and database connection.")
			ls.Listener.Close()
			db.DB.Close()
			return
		case notify := <-ls.Listener.Notify:
			log.Printf("Received notification on channel %s: %v\n", notify.Channel, notify.Extra)

			payload := model.PgNotificationPayload{}

			err := json.Unmarshal([]byte(notify.Extra), &payload)
			if err != nil {
				fmt.Println("Error:", err)
			}

			fmt.Printf("%v\n", payload)
		}
	}
}
