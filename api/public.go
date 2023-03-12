/*
 * Copyright (c) 2022-2023. Veteran Software
 *
 * Picarto API Wrapper - A custom wrapper for the Picarto REST API developed for a proprietary project.
 *
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 * License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later
 * version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
 * warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 */

package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	log "github.com/veteran-software/nowlive-logging"
)

const (
	apiBase = "https://api.picarto.tv/api"
	api     = apiBase + "/v1"
)

const (
	Adult = "adult="
)

// GetCategories
//
// Get information about all categories
//
//goland:noinspection GoUnusedExportedFunction
func GetCategories() *[]Category {
	resp, err := Rest.Request(http.MethodGet, api+"/categories", nil)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var categories []Category
	err = json.NewDecoder(resp.Body).Decode(&categories)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}

	return &categories
}

// GetChannelByID
//
// Gets information about a channel by ID - providing a bearer token with permission readpub can get followed status in result
//
//goland:noinspection GoUnusedExportedFunction
func GetChannelByID(channelID int) *Channel {
	resp, err := Rest.Request(http.MethodGet, api+fmt.Sprintf("/channel/id/%d", channelID), nil)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
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
//
//goland:noinspection GoUnusedExportedFunction
func GetChannelByName(channelName string) *Channel {
	resp, err := Rest.Request(http.MethodGet, api+fmt.Sprintf("/channel/name/%s", channelName), nil)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var channel Channel
	err = json.NewDecoder(resp.Body).Decode(&channel)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}

	return &channel
}

// GetAllChannelVideosByChannelID
//
// Get all videos for a channel by id
//
//goland:noinspection GoUnusedExportedFunction
func GetAllChannelVideosByChannelID(channelID int) *[]Video {
	resp, err := Rest.Request(http.MethodGet, api+fmt.Sprintf("/channel/id/%d/videos", channelID), nil)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var videos []Video
	err = json.NewDecoder(resp.Body).Decode(&videos)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}

	return &videos
}

// GetAllChannelVideosByChannelName
//
// Get all videos for a channel by name
//
//goland:noinspection GoUnusedExportedFunction
func GetAllChannelVideosByChannelName(channelName string) *[]Video {
	resp, err := Rest.Request(http.MethodGet, api+fmt.Sprintf("/channel/name/%s/videos", channelName), nil)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var videos []Video
	err = json.NewDecoder(resp.Body).Decode(&videos)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}

	return &videos
}

// GetOnline
//
// Gets all currently online channels - providing a bearer token with permission readpub can get followed status in result
//
//goland:noinspection GoUnusedExportedFunction
func GetOnline(adult *bool, gaming *bool, category ...string) *[]Online {
	var qsp []string
	if adult == nil {
		*adult = false
	}
	if gaming == nil {
		*gaming = false
	}

	qsp = append(qsp, Adult+strconv.FormatBool(*adult), "gaming="+strconv.FormatBool(*gaming))

	if category != nil && len(category) > 0 {
		qsp = append(qsp, "category="+strings.Join(category, ","))
	}

	var params string
	if len(qsp) > 0 {
		params = "?" + strings.Join(qsp, "&")
	}

	resp, err := Rest.Request(http.MethodGet, fmt.Sprintf(api+"/online%s", params), nil)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var online []Online
	err = json.NewDecoder(resp.Body).Decode(&online)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}

	return &online
}

// SearchChannels
//
// Get all channels matching the given search criteria (by name and tags)
//
//goland:noinspection GoUnusedExportedFunction
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
		Adult+strconv.FormatBool(*adult),
		"page="+strconv.FormatUint(*page, 10),
		"commissions="+strconv.FormatBool(*commissions))

	var params string
	if len(qsp) > 0 {
		params = "?" + strings.Join(qsp, "&")
	}

	resp, err := Rest.Request(http.MethodGet, fmt.Sprintf(api+"/search/channels%s", params), nil)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var channels []Channel
	err = json.NewDecoder(resp.Body).Decode(&channels)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}

	return &channels
}

// SearchVideos
//
// Get all channels matching the given search criteria (by name and tags)
//
//goland:noinspection GoUnusedExportedFunction
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
		Adult+strconv.FormatBool(*adult),
		"page="+strconv.FormatUint(*page, 10))

	var params string
	if len(qsp) > 0 {
		params = "?" + strings.Join(qsp, "&")
	}

	resp, err := Rest.Request(http.MethodGet, fmt.Sprintf(api+"/search/videos%s", params), nil)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var videos []Video
	err = json.NewDecoder(resp.Body).Decode(&videos)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}

	return &videos
}

// GetStreamByChannelID
//
// Get stream
//
//goland:noinspection GoUnusedExportedFunction
func GetStreamByChannelID(channelID int) *Stream {
	resp, err := Rest.Request(http.MethodGet, api+fmt.Sprintf("/channel/id/%d/streams", channelID), nil)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var stream Stream
	err = json.NewDecoder(resp.Body).Decode(&stream)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}

	return &stream
}

// GetStreamByChannelName
//
// Get stream
//
//goland:noinspection GoUnusedExportedFunction
func GetStreamByChannelName(channelName string) *Stream {
	resp, err := Rest.Request(http.MethodGet, api+fmt.Sprintf("/channel/name/%s/streams", channelName), nil)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var stream Stream
	err = json.NewDecoder(resp.Body).Decode(&stream)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}

	return &stream
}

// GetNotifications
//
// Get all global notifications/announcements
//
//goland:noinspection GoUnusedExportedFunction
func GetNotifications() *Notification {
	resp, err := Rest.Request(http.MethodGet, api+"/notifications", nil)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var notification Notification
	err = json.NewDecoder(resp.Body).Decode(&notification)
	if err != nil {
		log.Errorln(log.Picarto, log.FuncName(), err)
		return nil
	}

	return &notification
}
