package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bmrgcorp/discordgo"
)

// Flags
var (
	GuildID        = flag.String("guild", "", "Test guild ID")
	VoiceChannelID = flag.String("voice", "", "Test voice channel ID")
	BotToken       = flag.String("token", "", "Bot token")
)

func init() { flag.Parse() }

func main() {
	s, _ := discordgo.New("Bot " + *BotToken)
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready")
	})

	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")

}
