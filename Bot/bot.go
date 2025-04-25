package bot

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotToken string

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

func Run() {
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	discord.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsDirectMessages |
		discordgo.IntentsMessageContent

	discord.AddHandler(newMessage)

	// open session
	discord.Open()
	defer discord.Close() // close session, after function termination

	// keep bot running untill there is NO os interruption (ctrl + C)
	fmt.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func pickRandomHero(heroes []string) string {
	if len(heroes) == 0 {
		return "No heroes available"
	}
	pick := heroes[rand.Intn(len(heroes))]
	log.Printf("Randomly picked: %s", pick)
	return pick
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Don't respond to own message
	if message.Author.ID == discord.State.User.ID {
		return
	}

	log.Printf("Saw message %s", message.Content)

	heroes := []string{
		"Captain America", "Doctor Strange", "Emma Frost", "Groot", "Hulk", "Magneto", "Peni Parker", "The Thing", "Thor", "Venom",
		"Black Panther", "Black Widow", "Hawkeye", "Hela", "Human Torch", "Iron Fist", "Iron Man", "Magik", "Mr Fantastic", "Moon Knight", "Namor",
		"Psylocke", "Scarlet Witch", "Spider-Man", "Squirrel Girl", "Star-Lord", "Storm", "The Punisher",
		"Winter Soldier", "Wolverine", "Adam Warlock", "Cloak and Dagger", "Invisible Woman", "Jeff the Land Shark", "Loki",
		"Luna Snow", "Mantis", "Rocket Raccoon",
	}

	switch {
	case strings.Contains(message.Content, "!spin"):
		pick := pickRandomHero(heroes)
		mention := message.Author.Mention()
		returnMessage := "Spinning...\n" + mention + " your pick is: " + pick + "!"
		discord.ChannelMessageSend(message.ChannelID, returnMessage)
	}
}
