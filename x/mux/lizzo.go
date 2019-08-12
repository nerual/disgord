package mux

import (
	"github.com/bwmarrin/discordgo"
)

// Lizzo ...
// To use this function it must first be
// registered with the Mux.Route function.
func (m *Mux) Lizzo(ds *discordgo.Session, dm *discordgo.Message, ctx *Context) {

	// Set command prefix to display.

	resp := "\nGo to https://www.youtube.com/channel/UCXVMHu5xDH1oOfUGvaLyjGg for your Lizzo needs\n"

	ds.ChannelMessageSend(dm.ChannelID, resp)

	return
}
