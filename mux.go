package main

// This file adds the Disgord message route multiplexer, aka "command router".
// to the Disgord bot. This is an optional addition however it is included
// by default to demonstrate how to extend the Disgord bot.

import (
	"fmt"

	"github.com/nerual/pixelbot/x/mux"
)

// Router is registered as a global variable to allow easy access to the
// multiplexer throughout the bot.
var Router = mux.New()

func init() {
	// Register the mux OnMessageCreate handler that listens for and processes
	// all messages received.
	Session.AddHandler(Router.OnMessageCreate)

	// Register the build-in help command.
	Router.Prefix = "!"
	fmt.Printf("Your new Mux Router is: %#v\n", Router)
	Router.Route("help", "Display this message.", Router.Help)
	fmt.Printf("Router is: %#v\n", Router)

	Router.Route("lizzo", "Do a cool Lizzo thing.", Router.Lizzo)
}
