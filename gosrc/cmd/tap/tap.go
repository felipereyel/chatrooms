package tap

import (
	"chatrooms/gosrc/models"
	"chatrooms/gosrc/repositories/broker"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func Tap(cmd *cobra.Command, args []string) {
	roomId := args[0]
	broker, err := broker.NewBrokerRepo()
	if err != nil {
		panic(err.Error())
	}

	subscription, err := broker.Subscribe(roomId)
	if err != nil {
		panic(err.Error())
	}
	defer subscription.Close()

	for msg := range subscription.MessageChan {
		body := msg.Body
		var post models.PostView
		if err := json.Unmarshal(body, &post); err != nil {
			fmt.Println("Bad Message: ", string(body))
			continue
		}

		fmt.Printf("%s: %s\n", post.Username, post.Content)
	}
}
