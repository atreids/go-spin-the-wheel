package bot

import (
	"log"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func pickRandomHero(heroes []string) string {
	if len(heroes) == 0 {
		return "No heroes available"
	}
	pick := heroes[rand.Intn(len(heroes))]
	log.Printf("Randomly picked: %s", pick)
	return pick
}

func randomHeroHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Don't respond to own message
	if message.Author.ID == discord.State.User.ID {
		return
	}

	heroes := []string{
		"Captain America", "Doctor Strange", "Emma Frost", "Groot", "Hulk", "Magneto", "Peni Parker", "The Thing", "Thor", "Venom",
		"Black Panther", "Black Widow", "Hawkeye", "Hela", "Human Torch", "Iron Fist", "Iron Man", "Magik", "Mr Fantastic", "Moon Knight", "Namor",
		"Psylocke", "Scarlet Witch", "Spider-Man", "Squirrel Girl", "Star-Lord", "Storm", "The Punisher",
		"Winter Soldier", "Wolverine", "Adam Warlock", "Cloak and Dagger", "Invisible Woman", "Jeff the Land Shark", "Loki",
		"Luna Snow", "Mantis", "Rocket Raccoon",
	}

	switch {
	case strings.Contains(message.Content, "!spin"):
		log.Printf("Saw message: %s", message.Content)
		log.Print("Picking a random hero.")
		pick := pickRandomHero(heroes)
		mention := message.Author.Mention()
		returnMessage := "Spinning...\n" + mention + " your pick is: " + pick + "!"
		discord.ChannelMessageSend(message.ChannelID, returnMessage)
	}
}
