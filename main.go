// Declare this file to be part of the main package so it can be compiled into
// an executable.
package main

// Import all Go packages required for this file.
import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Version is a constant that stores the Disgord version information.
const Version = "v0.0.0-alpha"

// Session is declared in the global space so it can be easily used
// throughout this program.
// In this use case, there is no error that would be returned.

var Token = os.Getenv("DG_TOKEN")
var Session, _ = discordgo.New()

// Read in all configuration options from both environment variables and
// command line arguments.
func init() {

	// Discord Authentication Token
	Session.Token = "Bot " + os.Getenv("DG_TOKEN")
	if Session.Token == "" {
		flag.StringVar(&Session.Token, "t", "", "Discord Authentication Token")
	}
	fmt.Printf("Initialsizing bot token to Bot %v\n", Token)
}

func main() {

	// Declare any variables needed later.
	var err error

	// Print out a fancy logo!
	fmt.Printf(` 
	________  .__                               .___
	\______ \ |__| ______ ____   ___________  __| _/
	||    |  \|  |/  ___// ___\ /  _ \_  __ \/ __ | 
	||    '   \  |\___ \/ /_/  >  <_> )  | \/ /_/ | 
	||______  /__/____  >___  / \____/|__|  \____ | 
	\_______\/        \/_____/   %-16s\/`+"\n\n", Version)

	// Parse command line arguments
	flag.Parse()

	// Verify a Token was provided
	if Session.Token == "" {
		log.Println("You must provide a Discord authentication token.")
		return
	}
	// Open a websocket connection to Discord
	err = Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}
	mainChan, err := Session.Channel("608346331605368835")
	fmt.Printf("I maybe joined %v\n", mainChan)
	if err != nil {
		log.Printf("error joining channel, %s\n", err)
	}
	Session.AddHandler(messageCreate)
	// Wait for a CTRL-C
	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Clean up
	Session.Close()

	// Exit Normally.
}
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
