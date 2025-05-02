package bot

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func randomisePlayers(players []string) []string {
	if len(players) == 0 {
		log.Print("No players found")
	}
	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})
	return players
}

func errorLog(err error) {
	fmt.Printf("ERROR: %+v\n", err)
}

func randomTeamsHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	if !strings.HasPrefix(message.Content, "!random_teams") {
		return
	}

	log.Printf("Saw message: %s", message.Content)

	guildID := message.GuildID
	userID := message.Author.ID

	if guildID == "" {
		log.Print("User was not in a server, returning early.")
		discord.ChannelMessageSend(message.ChannelID, "You must be in a server (and in a voice channel) to send this command.")
		return
	}

	// Get voice channel that the requesting user is in.
	requestingUserVoiceState, err := discord.State.VoiceState(guildID, userID)
	if err == discordgo.ErrStateNotFound {
		discord.ChannelMessageSend(message.ChannelID, "You must be in a voice channel to use this command.")
		return
	} else if err != nil {
		errorLog(err)
		return
	}

	// Get whole guild (aka server) details
	guild, err := discord.State.Guild(guildID)
	if err != nil {
		errorLog(err)
		return
	}

	// Loop through all connected VoiceStates (members in a voice channel) in the server, if they are in the same
	// voice channel as requesting user, grab their username and append to array
	var players []string
	for _, vs := range guild.VoiceStates {
		if vs.ChannelID == requestingUserVoiceState.ChannelID {
			// Get specific member details
			member, err := discord.State.Member(guildID, vs.UserID)
			if err != nil {
				errorLog(err)
				continue
			}
			players = append(players, member.User.GlobalName)
		}
	}

	if len(players) < 2 {
		discord.ChannelMessageSend(message.ChannelID, "There's only one of you! You can figure it out ;)")
		return
	}

	randomisedList := randomisePlayers(players)
	mid := len(randomisedList) / 2
	team1 := players[:mid]
	team2 := players[mid:]

	msg := "**Random Teams:**\n"
	msg += "__Team 1:__ " + strings.Join(team1, ", ") + "\n"
	msg += "__Team 2:__ " + strings.Join(team2, ", ")

	log.Printf("Generated random teams:\n%s", msg)
	discord.ChannelMessageSend(message.ChannelID, msg)
}
