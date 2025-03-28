# DiscordGo

[![Go Reference](https://pkg.go.dev/badge/github.com/bmrgcorp/discordgo.svg)](https://pkg.go.dev/github.com/bmrgcorp/discordgo) [![Go Report Card](https://goreportcard.com/badge/github.com/bmrgcorp/discordgo)](https://goreportcard.com/report/github.com/bmrgcorp/discordgo) [![Discord Gophers](https://img.shields.io/badge/Discord%20Gophers-%23discordgo-blue.svg)](https://discord.gg/golang) [![Discord API](https://img.shields.io/badge/Discord%20API-%23go_discordgo-blue.svg)](https://discord.com/invite/discord-api)

<img align="right" alt="DiscordGo logo" src="docs/img/discordgo.svg" width="400">

DiscordGo is a [Go](https://golang.org/) package that provides low level 
bindings to the [Discord](https://discord.com/) chat client API. DiscordGo 
has nearly complete support for all of the Discord API endpoints, websocket
interface, and voice interface.

If you would like to help the DiscordGo package please use 
[this link](https://discord.com/oauth2/authorize?client_id=173113690092994561&scope=bot)
to add the official DiscordGo test bot **dgo** to your server. This provides 
indispensable help to this project.

* See [dgVoice](https://github.com/bwmarrin/dgvoice) package for an example of
additional voice helper functions and features for DiscordGo.

* See [dca](https://github.com/bwmarrin/dca) for an **experimental** stand alone
tool that wraps `ffmpeg` to create opus encoded audio appropriate for use with
Discord (and DiscordGo).

> [!WARNING]  
> Boomerang does not provide support for this package. Please visit the official repository [here](<https://github.com/bwmarrin/discordgo>) instead.