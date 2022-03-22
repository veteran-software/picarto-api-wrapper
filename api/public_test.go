/*
 * Copyright (c) 2022. Veteran Software
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
	"reflect"
	"testing"
)

const (
	nilError         = "an error occurred and nil was returned"
	returnTypeBool   = "unexpected return type; expected: bool, got: %T"
	returnTypeInt    = "unexpected return type; expected: int, got: %T"
	returnTypeString = "unexpected return type; expected: string, got: %T"
)

func TestGetCategories(t *testing.T) {
	Rest = NewPicarto("TEST_TOKEN")

	categories := GetCategories()

	if categories == nil {
		t.Error(nilError)
		return
	}

	if len(*categories) == 0 {
		t.Error("empty category list returned")
		return
	}

	for _, cat := range *categories {
		if reflect.ValueOf(cat.ID).Kind() != reflect.Int {
			t.Errorf(returnTypeInt, reflect.ValueOf(cat.ID).Kind())
		}
		if reflect.ValueOf(cat.Name).Kind() != reflect.String {
			t.Errorf(returnTypeString, reflect.ValueOf(cat.Name).Kind())
		}
		if reflect.ValueOf(cat.TotalChannels).Kind() != reflect.Int {
			t.Errorf(returnTypeInt, reflect.ValueOf(cat.TotalChannels).Kind())
		}
		if reflect.ValueOf(cat.OnlineChannels).Kind() != reflect.Int {
			t.Errorf(returnTypeInt, reflect.ValueOf(cat.OnlineChannels).Kind())
		}
		if reflect.ValueOf(cat.Viewers).Kind() != reflect.Int {
			t.Errorf(returnTypeInt, reflect.ValueOf(cat.Viewers).Kind())
		}
		if reflect.ValueOf(cat.Adult).Kind() != reflect.Bool {
			t.Errorf(returnTypeBool, reflect.ValueOf(cat.Adult).Kind())
		}
	}
}

func TestGetChannelByID(t *testing.T) {
	Rest = NewPicarto("TEST_TOKEN")

	c := GetChannelByID(527732)

	if c == nil {
		t.Error(nilError)
		return
	}

	evalChannelFields(t, c)
}

func TestGetChannelByName(t *testing.T) {
	Rest = NewPicarto("TEST_TOKEN")

	c := GetChannelByName("AgueMort")

	if c == nil {
		t.Error(nilError)
		return
	}

	evalChannelFields(t, c)
}

func evalChannelFields(t *testing.T, c *Channel) {
	/* Type check all params */
	if reflect.ValueOf(c.UserId).Kind() != reflect.Int64 {
		t.Errorf(returnTypeInt, reflect.ValueOf(c.UserId).Kind())
	}
	if reflect.ValueOf(c.Name).Kind() != reflect.String {
		t.Errorf(returnTypeString, reflect.ValueOf(c.Name).Kind())
	}
	if reflect.ValueOf(c.Avatar).Kind() != reflect.String {
		t.Errorf(returnTypeString, reflect.ValueOf(c.Avatar).Kind())
	}
	if reflect.ValueOf(c.Online).Kind() != reflect.Bool {
		t.Errorf(returnTypeBool, reflect.ValueOf(c.Online).Kind())
	}
	if reflect.ValueOf(c.Viewers).Kind() != reflect.Int64 {
		t.Errorf(returnTypeInt, reflect.ValueOf(c.Viewers).Kind())
	}
	if reflect.ValueOf(c.ViewersTotal).Kind() != reflect.Int64 {
		t.Errorf(returnTypeInt, reflect.ValueOf(c.ViewersTotal).Kind())
	}
	if reflect.ValueOf(c.Thumbnails.Web).Kind() != reflect.String {
		t.Errorf(returnTypeString, reflect.ValueOf(c.Thumbnails.Web).Kind())
	}
	if reflect.ValueOf(c.Thumbnails.WebLarge).Kind() != reflect.String {
		t.Errorf(returnTypeString, reflect.ValueOf(c.Thumbnails.WebLarge).Kind())
	}
	if reflect.ValueOf(c.Thumbnails.Mobile).Kind() != reflect.String {
		t.Errorf(returnTypeString, reflect.ValueOf(c.Thumbnails.Mobile).Kind())
	}
	if reflect.ValueOf(c.Thumbnails.Tablet).Kind() != reflect.String {
		t.Errorf(returnTypeString, reflect.ValueOf(c.Thumbnails.Tablet).Kind())
	}
	if reflect.ValueOf(c.Followers).Kind() != reflect.Int64 {
		t.Errorf(returnTypeInt, reflect.ValueOf(c.Followers).Kind())
	}
	if reflect.ValueOf(c.Subscribers).Kind() != reflect.Int64 {
		t.Errorf(returnTypeInt, reflect.ValueOf(c.Subscribers).Kind())
	}
	if reflect.ValueOf(c.Adult).Kind() != reflect.Bool {
		t.Errorf(returnTypeBool, reflect.ValueOf(c.Adult).Kind())
	}
	if reflect.ValueOf(c.Category).Kind() != reflect.Slice {
		t.Errorf(returnTypeString, reflect.ValueOf(c.Category).Kind())
	}
	if reflect.ValueOf(c.AccountType).Kind() != reflect.String {
		t.Errorf(returnTypeString, reflect.ValueOf(c.AccountType).Kind())
	}
	if reflect.ValueOf(c.Commissions).Kind() != reflect.Bool {
		t.Errorf(returnTypeBool, reflect.ValueOf(c.Commissions).Kind())
	}
	if reflect.ValueOf(c.Recordings).Kind() != reflect.Bool {
		t.Errorf(returnTypeBool, reflect.ValueOf(c.Recordings).Kind())
	}
}
