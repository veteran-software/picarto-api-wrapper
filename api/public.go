package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

import (
	"os"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var (
	log = logrus.Logger{
		Out: os.Stderr,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05.999",
			LogFormat:       "[%lvl%]: %msg%\n",
		},
		ReportCaller: true,
	}
)

const (
	apiBase = "https://api.picarto.tv/api"
	api     = apiBase + "/v1"
)

// GetCategories
//
// Get information about all categories
func GetCategories() *[]Category {
	resp, err := Rest.Request(http.MethodGet, api+"/categories", nil, &token)
	if err != nil {
		log.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var categories []Category
	err = json.NewDecoder(resp.Body).Decode(&categories)
	if err != nil {
		log.Errorln(err)
		return nil
	}

	return &categories
}

// GetChannelByID
//
// Gets information about a channel by ID - providing a bearer token with permission readpub can get followed status in result
func GetChannelByID(channelID int) *Channel {
	resp, err := Rest.Request(http.MethodGet, api+fmt.Sprintf("/channel/id/%d", channelID), nil, &token)
	if err != nil {
		log.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var channel Channel
	err = json.NewDecoder(resp.Body).Decode(&channel)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &channel
}

// GetChannelByName
//
// Gets information about a channel by name - providing a bearer token with permission readpub can get followed status in result
func GetChannelByName(channelName string) *Channel {
	resp, err := Rest.Request(http.MethodGet, api+fmt.Sprintf("/channel/name/%s", channelName), nil, &token)
	if err != nil {
		log.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var channel Channel
	err = json.NewDecoder(resp.Body).Decode(&channel)
	if err != nil {
		log.Errorln(err)
		return nil
	}

	return &channel
}

// GetAllChannelVideosByID
//
// Get all videos for a channel by id
func GetAllChannelVideosByID(channelID int) *[]Video {
	resp, err := Rest.Request(http.MethodGet, api+fmt.Sprintf("/channel/id/%d/videos", channelID), nil, &token)
	if err != nil {
		log.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var videos []Video
	err = json.NewDecoder(resp.Body).Decode(&videos)
	if err != nil {
		log.Errorln(err)
		return nil
	}

	return &videos
}

//
// GetAllChannelVideosByName
//
// Get all videos for a channel by name
func GetAllChannelVideosByName(channelName string) *[]Video {
	resp, err := Rest.Request(http.MethodGet, api+fmt.Sprintf("/channel/name/%s/videos", channelName), nil, &token)
	if err != nil {
		log.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var videos []Video
	err = json.NewDecoder(resp.Body).Decode(&videos)
	if err != nil {
		log.Errorln(err)
		return nil
	}

	return &videos
}

// GetOnline
//
// Gets all currently online channels - providing a bearer token with permission readpub can get followed status in result
func GetOnline(adult *bool, gaming *bool, category ...string) *[]Online {
	var qsp []string
	if adult == nil {
		*adult = false
	}
	if gaming == nil {
		*gaming = false
	}

	qsp = append(qsp, "adult="+strconv.FormatBool(*adult), "gaming="+strconv.FormatBool(*gaming))

	if category != nil || len(category) > 0 {
		qsp = append(qsp, "category="+strings.Join(category, ","))
	}

	var params string
	if len(qsp) > 0 {
		params = "?" + strings.Join(qsp, "&")
	}

	resp, err := Rest.Request(http.MethodGet, fmt.Sprintf(api+"/online%s", params), nil, &token)
	if err != nil {
		log.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var online []Online
	err = json.NewDecoder(resp.Body).Decode(&online)
	if err != nil {
		log.Errorln(err)
		return nil
	}

	return &online
}

// SearchChannels
//
// Get all channels matching the given search criteria (by name and tags)
func SearchChannels(q string, adult *bool, page *uint64, commissions *bool) *[]Channel {
	var qsp []string

	if q == "" {
		return nil
	}

	qsp = append(qsp, q)

	if adult == nil {
		*adult = false
	}
	if page == nil {
		*page = 1
	}
	if commissions == nil {
		*commissions = false
	}

	qsp = append(qsp,
		"adult="+strconv.FormatBool(*adult),
		"page="+strconv.FormatUint(*page, 10),
		"commissions="+strconv.FormatBool(*commissions))

	var params string
	if len(qsp) > 0 {
		params = "?" + strings.Join(qsp, "&")
	}

	resp, err := Rest.Request(http.MethodGet, fmt.Sprintf(api+"/search/channels%s", params), nil, &token)
	if err != nil {
		log.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var channels []Channel
	err = json.NewDecoder(resp.Body).Decode(&channels)
	if err != nil {
		log.Errorln(err)
		return nil
	}

	return &channels
}

// SearchVideos
//
// Get all channels matching the given search criteria (by name and tags)
func SearchVideos(q string, adult *bool, page *uint64) *[]Video {
	var qsp []string

	if q == "" {
		return nil
	}

	qsp = append(qsp, q)

	if adult == nil {
		*adult = false
	}
	if page == nil {
		*page = 1
	}

	qsp = append(qsp,
		"adult="+strconv.FormatBool(*adult),
		"page="+strconv.FormatUint(*page, 10))

	var params string
	if len(qsp) > 0 {
		params = "?" + strings.Join(qsp, "&")
	}

	resp, err := Rest.Request(http.MethodGet, fmt.Sprintf(api+"/search/videos%s", params), nil, &token)
	if err != nil {
		log.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var videos []Video
	err = json.NewDecoder(resp.Body).Decode(&videos)
	if err != nil {
		log.Errorln(err)
		return nil
	}

	return &videos
}

// GetStreamByChannelID
//
// Get stream
func GetStreamByChannelID(channelID int) *Stream {
	resp, err := Rest.Request(http.MethodGet, api+fmt.Sprintf("/channel/id/%d/streams", channelID), nil, &token)
	if err != nil {
		log.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var stream Stream
	err = json.NewDecoder(resp.Body).Decode(&stream)
	if err != nil {
		log.Errorln(err)
		return nil
	}

	return &stream
}

// GetStreamByChannelName
//
// Get stream
func GetStreamByChannelName(channelName string) *Stream {
	resp, err := Rest.Request(http.MethodGet, api+fmt.Sprintf("/channel/name/%s/streams", channelName), nil, &token)
	if err != nil {
		log.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var stream Stream
	err = json.NewDecoder(resp.Body).Decode(&stream)
	if err != nil {
		log.Errorln(err)
		return nil
	}

	return &stream
}

// GetNotifications
//
// Get all global notifications/announcements
func GetNotifications() *Notification {
	resp, err := Rest.Request(http.MethodGet, api+"/notifications", nil, &token)
	if err != nil {
		log.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var notification Notification
	err = json.NewDecoder(resp.Body).Decode(&notification)
	if err != nil {
		log.Errorln(err)
		return nil
	}

	return &notification
}
