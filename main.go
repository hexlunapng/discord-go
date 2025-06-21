package maim

import (
"flag"
"fmt"
"os"
"os/singal"
"syscall"
"github.com/bwmarrin/discordgo"
)

var (
	token = flag.String("token", "", "Discord bot token")
	GuildID = flag.String("guild", "", "Guild ID")
)

var bot *discordgo.Session


func init() { flag.Parse ()}

func init () { 
	var err error
	bot, err = discordgo.NewSession(token)
	if err != nil {
		fmt.Println(err)
		return
}

var handler = func(s *discordgo.Session, m *discordgo.MessageCreate) {
	}
	func main() {
		bot.AddHandler(handler)
 bot.Open()
 if err := bot.Identify(); err != nil {
	fmt.Println(err)
	return
 }
 defer bot.Close()
 
 fmt.Println("bot is running")
 stop := make(chan os.Signal, 1)
 signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
 <-stop
}