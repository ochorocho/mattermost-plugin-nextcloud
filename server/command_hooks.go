package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	//	"github.com/ochorocho/mattermost-plugin-nextcloud/server/nextcloud"
)

const (
	commandTriggerNextcloud = "nextcloud"

	loginSuffixUrl      = "_nextcloud_url"
	loginSuffixUsername = "_nextcloud_username"
	loginSuffixPassword = "_nextcloud_password"

	commandDialogHelp = "###### Nextcloud slash command Help\n" +
		"| Command| Description |\n" +
		"|--------|-------------|\n" +
		"| `/nextcloud` 		| . |\n" +
		"| `/nextcloud login` 	| Login to your Nextcloud instance. |\n" +
		"| `/nextcloud new` 	| Create a new Meeting. |\n" +
		"| `/nextcloud help` | Show this help text |"
)

func (p *Plugin) registerCommands() error {
	if err := p.API.RegisterCommand(&model.Command{
		Trigger:          commandTriggerNextcloud,
		AutoComplete:     true,
		AutoCompleteDesc: "Manage Nextcloud",
		DisplayName:      "Manage Nextcloud conversations and meetings",
	}); err != nil {
		return errors.Wrapf(err, "failed to register %s command", commandTriggerNextcloud)
	}

	return nil
}

func (p *Plugin) emitStatusChange() {
	configuration := p.getConfiguration()

	p.API.PublishWebSocketEvent("status_change", map[string]interface{}{
		"enabled": !configuration.disabled,
	}, &model.WebsocketBroadcast{})
}

// ExecuteCommand executes a command that has been previously registered via the RegisterCommand
// API.
func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	trigger := strings.TrimPrefix(strings.Fields(args.Command)[0], "/")
	switch trigger {
	case commandTriggerNextcloud:
		return p.executeCommandNextcloud(args), nil
	default:
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         fmt.Sprintf("Unknown command: " + args.Command),
		}, nil
	}
}

func (p *Plugin) setNextcloudCredentials(user *model.User, url string, username string, password string) [3]string {
	var keyUrl = user.Id + loginSuffixUrl
	p.API.KVSet(keyUrl, []byte(url))

	var keyUsername = user.Id + loginSuffixUsername
	p.API.KVSet(keyUsername, []byte(username))

	var keyPassword = user.Id + loginSuffixPassword
	p.API.KVSet(keyPassword, []byte(password))

	return p.getNextcloudCredentials(user)
}

func (p *Plugin) getNextcloudCredentials(user *model.User) [3]string {
	var keyUrl = user.Id + loginSuffixUrl
	var keyUsername = user.Id + loginSuffixUsername
	var keyPassword = user.Id + loginSuffixPassword
	var credentials [3]string

	url, err := p.API.KVGet(keyUrl)
	if err != nil {
		p.API.LogError("Could not get Nextcloud url")
	}

	username, err := p.API.KVGet(keyUsername)
	if err != nil {
		p.API.LogError("Could not get Nextcloud username")
	}

	password, err := p.API.KVGet(keyPassword)
	if err != nil {
		p.API.LogError("Could not get Nextcloud app password")
	}

	credentials[0] = string(url)
	credentials[1] = string(username)
	credentials[2] = string(password)

	return credentials
}

func (p *Plugin) executeCommandNextcloud(args *model.CommandArgs) *model.CommandResponse {
	//serverConfig := p.API.GetConfig()

	user, _ := p.API.GetUser(args.UserId)
	fields := strings.Fields(args.Command)

	command := ""
	if len(fields) >= 2 {
		command = fields[1]
	}

	switch command {
	case "help":
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         commandDialogHelp,
		}
	case "":
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         commandDialogHelp,
		}
	case "status":
		var msg = ""

		if len(fields) != 2 {
			msg = "Too many arguments use `/nextcloud status`"
		} else {
			var credentials = p.getNextcloudCredentials(user)
			var maskedPassword = p.mask(credentials[2])

			msg = "|    |\n" // Required to get rid of header row
			msg += "|--------------|------------------------|\n"
			msg += "| **URL**      | " + credentials[0] + " |\n"
			msg += "| **Username** | " + credentials[1] + " |\n"
			msg += "| **Password** | " + maskedPassword + " |\n"
		}

		p.API.SendEphemeralPost(user.Id, &model.Post{
			UserId:    p.botUserID,
			ChannelId: args.ChannelId,
			Message:   msg,
		})
	case "login":

		var msg = ""

		if len(fields) != 5 {
			msg = "Missing arguments! Does your command look like this?\n`/nextcloud login [url] [username] [password]`"
		} else {
			url := fields[2]
			username := fields[3]
			password := fields[4]

			var credentials = p.setNextcloudCredentials(user, url, username, password)
			msg = "Credentials for **" + credentials[1] + "** on **" + credentials[0] + "** set."
		}

		p.API.SendEphemeralPost(user.Id, &model.Post{
			UserId:    p.botUserID,
			ChannelId: args.ChannelId,
			Message:   msg,
		})
	default:
		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         fmt.Sprintf("Unknown command: " + command),
		}
	}

	//if err := p.API.OpenInteractiveDialog(dialogRequest); err != nil {
	//	errorMessage := "Failed to open Interactive Dialog"
	//	p.API.LogError(errorMessage, "err", err.Error())
	//	return &model.CommandResponse{
	//		ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
	//		Text:         errorMessage,
	//	}
	//}

	return &model.CommandResponse{}
}

//var keyUrl = user.Id + "_nextcloud_url"
//var keyUsername = user.Id + "_nextcloud_username"
//var keyPassword = user.Id + "_nextcloud_password"
//
//url, err := p.API.KVGet(keyUrl)
//if err != nil {
//p.API.LogError("Could not get Nextcloud url")
//}
//
//username, err := p.API.KVGet(keyUsername)
//if err != nil {
//p.API.LogError("Could not get Nextcloud username")
//}
//
//password, err := p.API.KVGet(keyPassword)
//if err != nil {
//p.API.LogError("Could not get Nextcloud app password")
//}
//
//profile, _ := p.API.GetBundlePath()
//
//p.API.LogWarn(profile)
