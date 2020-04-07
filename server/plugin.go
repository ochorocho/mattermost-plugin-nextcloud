package main

import (
	"fmt"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"io/ioutil"
	"net/http"
	//"github.com/ochorocho/mattermost-plugin-nextcloud/server/nextcloud"
	"path/filepath"
	"sync"
)

const (
	botUserName    = "Nextcloud"
	botDisplayName = "Nextcloud Talk"
	botDescription = "Created by the Nextcloud plugin."
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	botUserID string

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}



type Rooms struct {
	Name string
	Id  int
}


// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello, world!")

	switch r.URL.Path {
	case "/status":
		roomList := nextcloud.Request{"spreed", "room", "GET", ""}

		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write([]byte(roomList.Request())); err != nil {
			p.API.LogError("failed to write rooms", "err", err.Error())
		}

		userID := r.Header.Get("Mattermost-User-Id")
		if userID == "" {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		user, _ := p.API.GetUser(userID)

		_, appErr := p.postMeeting(user.Username, 3, "off-topic", "Look at this amazing meeting")
		if appErr != nil {
			http.Error(w, appErr.Error(), appErr.StatusCode)
			return
		}

	}
}

func (p *Plugin) postMeeting(creatorUsername string, meetingID int, channelID string, topic string) (*model.Post, *model.AppError) {

	meetingURL := "sadsadsadas"

	post := &model.Post{
		UserId:    p.botUserID,
		ChannelId: "sb4u9sxabpbq3qtaf1qod4fkuc",
		Message:   fmt.Sprintf("Meeting started at %s.", meetingURL),
		Type:      "custom_zoom",
		Props: map[string]interface{}{
			"meeting_id":               "meetingID",
			"meeting_link":             "meetingURL",
			"meeting_status":           creatorUsername + "########",
			"meeting_personal":         true,
			"meeting_topic":            topic + "<h1>ldkjdslkadj</h1>",
			"meeting_creator_username": creatorUsername,
		},
	}

	return p.API.CreatePost(post)
}

// OnActivate checks if the configurations is valid and ensures the bot account exists
func (p *Plugin) OnActivate() error {
	//config := p.getConfiguration()
	//if err := config.IsValid(); err != nil {
	//	return err
	//}

	botUserID, err := p.Helpers.EnsureBot(&model.Bot{
		Username:    botUserName,
		DisplayName: botDisplayName,
		Description: botDescription,
	})
	if err != nil {
		//return errors.Wrap(err, "failed to ensure bot account")
	}
	p.botUserID = botUserID

	bundlePath, err := p.API.GetBundlePath()
	if err != nil {
		//return errors.Wrap(err, "couldn't get bundle path")
	}

	//if err = p.API.RegisterCommand(getCommand()); err != nil {
	//	//return errors.WithMessage(err, "OnActivate: failed to register command")
	//}

	profileImage, err := ioutil.ReadFile(filepath.Join(bundlePath, "assets", "profile.png"))
	if err != nil {
		//return errors.Wrap(err, "couldn't read profile image")
	}

	if appErr := p.API.SetProfileImage(botUserID, profileImage); appErr != nil {
		//return errors.Wrap(appErr, "couldn't set profile image")
	}

	//p.zoomClient = zoom.NewClient(config.ZoomAPIURL, config.APIKey, config.APISecret)

	return nil
}
