package api

import (
	"reflect"
	"testing"
)

func TestGetCategories(t *testing.T) {
	Rest = NewPicarto("TEST_TOKEN")

	categories := GetCategories()

	if categories == nil {
		t.Error("an error occurred and nil was returned")
		return
	}

	if len(*categories) == 0 {
		t.Error("empty category list returned")
		return
	}

	for _, cat := range *categories {
		if reflect.ValueOf(cat.ID).Kind() != reflect.Int {
			t.Errorf("unexpected return type; expected: int, got: %T", reflect.ValueOf(cat.ID).Kind())
		}
		if reflect.ValueOf(cat.Name).Kind() != reflect.String {
			t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(cat.Name).Kind())
		}
		if reflect.ValueOf(cat.TotalChannels).Kind() != reflect.Int {
			t.Errorf("unexpected return type; expected: int, got: %T", reflect.ValueOf(cat.TotalChannels).Kind())
		}
		if reflect.ValueOf(cat.OnlineChannels).Kind() != reflect.Int {
			t.Errorf("unexpected return type; expected: int, got: %T", reflect.ValueOf(cat.OnlineChannels).Kind())
		}
		if reflect.ValueOf(cat.Viewers).Kind() != reflect.Int {
			t.Errorf("unexpected return type; expected: int, got: %T", reflect.ValueOf(cat.Viewers).Kind())
		}
		if reflect.ValueOf(cat.Adult).Kind() != reflect.Bool {
			t.Errorf("unexpected return type; expected: bool, got: %T", reflect.ValueOf(cat.Adult).Kind())
		}
	}
}

func TestGetChannelByID(t *testing.T) {
	Rest = NewPicarto("TEST_TOKEN")

	c := GetChannelByID(527732)

	if c == nil {
		t.Error("an error occurred and nil was returned")
		return
	}

	/* Type check all params */
	if reflect.ValueOf(c.UserId).Kind() != reflect.Int64 {
		t.Errorf("unexpected return type; expected: int, got: %T", reflect.ValueOf(c.UserId).Kind())
	}
	if reflect.ValueOf(c.Name).Kind() != reflect.String {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.Name).Kind())
	}
	if reflect.ValueOf(c.Avatar).Kind() != reflect.String {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.Avatar).Kind())
	}
	if reflect.ValueOf(c.Online).Kind() != reflect.Bool {
		t.Errorf("unexpected return type; expected: bool, got: %T", reflect.ValueOf(c.Online).Kind())
	}
	if reflect.ValueOf(c.Viewers).Kind() != reflect.Int64 {
		t.Errorf("unexpected return type; expected: int, got: %T", reflect.ValueOf(c.Viewers).Kind())
	}
	if reflect.ValueOf(c.ViewersTotal).Kind() != reflect.Int64 {
		t.Errorf("unexpected return type; expected: int, got: %T", reflect.ValueOf(c.ViewersTotal).Kind())
	}
	if reflect.ValueOf(c.Thumbnails.Web).Kind() != reflect.String {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.Thumbnails.Web).Kind())
	}
	if reflect.ValueOf(c.Thumbnails.WebLarge).Kind() != reflect.String {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.Thumbnails.WebLarge).Kind())
	}
	if reflect.ValueOf(c.Thumbnails.Mobile).Kind() != reflect.String {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.Thumbnails.Mobile).Kind())
	}
	if reflect.ValueOf(c.Thumbnails.Tablet).Kind() != reflect.String {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.Thumbnails.Tablet).Kind())
	}
	if reflect.ValueOf(c.Followers).Kind() != reflect.Int64 {
		t.Errorf("unexpected return type; expected: int, got: %T", reflect.ValueOf(c.Followers).Kind())
	}
	if reflect.ValueOf(c.Subscribers).Kind() != reflect.Int64 {
		t.Errorf("unexpected return type; expected: int, got: %T", reflect.ValueOf(c.Subscribers).Kind())
	}
	if reflect.ValueOf(c.Adult).Kind() != reflect.Bool {
		t.Errorf("unexpected return type; expected: bool, got: %T", reflect.ValueOf(c.Adult).Kind())
	}
	if reflect.ValueOf(c.Category).Kind() != reflect.Slice {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.Category).Kind())
	}
	if reflect.ValueOf(c.AccountType).Kind() != reflect.String {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.AccountType).Kind())
	}
	if reflect.ValueOf(c.Commissions).Kind() != reflect.Bool {
		t.Errorf("unexpected return type; expected: bool, got: %T", reflect.ValueOf(c.Commissions).Kind())
	}
	if reflect.ValueOf(c.Recordings).Kind() != reflect.Bool {
		t.Errorf("unexpected return type; expected: bool, got: %T", reflect.ValueOf(c.Recordings).Kind())
	}
}

func TestGetChannelByName(t *testing.T) {
	Rest = NewPicarto("TEST_TOKEN")

	c := GetChannelByName("AgueMort")

	if c == nil {
		t.Error("an error occurred and nil was returned")
		return
	}

	/* Type check all params */
	if reflect.ValueOf(c.UserId).Kind() != reflect.Int64 {
		t.Errorf("unexpected return type; expected: int, got: %T", reflect.ValueOf(c.UserId).Kind())
	}
	if reflect.ValueOf(c.Name).Kind() != reflect.String {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.Name).Kind())
	}
	if reflect.ValueOf(c.Avatar).Kind() != reflect.String {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.Avatar).Kind())
	}
	if reflect.ValueOf(c.Online).Kind() != reflect.Bool {
		t.Errorf("unexpected return type; expected: bool, got: %T", reflect.ValueOf(c.Online).Kind())
	}
	if reflect.ValueOf(c.Viewers).Kind() != reflect.Int64 {
		t.Errorf("unexpected return type; expected: int, got: %T", reflect.ValueOf(c.Viewers).Kind())
	}
	if reflect.ValueOf(c.ViewersTotal).Kind() != reflect.Int64 {
		t.Errorf("unexpected return type; expected: int, got: %T", reflect.ValueOf(c.ViewersTotal).Kind())
	}
	if reflect.ValueOf(c.Thumbnails.Web).Kind() != reflect.String {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.Thumbnails.Web).Kind())
	}
	if reflect.ValueOf(c.Thumbnails.WebLarge).Kind() != reflect.String {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.Thumbnails.WebLarge).Kind())
	}
	if reflect.ValueOf(c.Thumbnails.Mobile).Kind() != reflect.String {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.Thumbnails.Mobile).Kind())
	}
	if reflect.ValueOf(c.Thumbnails.Tablet).Kind() != reflect.String {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.Thumbnails.Tablet).Kind())
	}
	if reflect.ValueOf(c.Followers).Kind() != reflect.Int64 {
		t.Errorf("unexpected return type; expected: int, got: %T", reflect.ValueOf(c.Followers).Kind())
	}
	if reflect.ValueOf(c.Subscribers).Kind() != reflect.Int64 {
		t.Errorf("unexpected return type; expected: int, got: %T", reflect.ValueOf(c.Subscribers).Kind())
	}
	if reflect.ValueOf(c.Adult).Kind() != reflect.Bool {
		t.Errorf("unexpected return type; expected: bool, got: %T", reflect.ValueOf(c.Adult).Kind())
	}
	if reflect.ValueOf(c.Category).Kind() != reflect.Slice {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.Category).Kind())
	}
	if reflect.ValueOf(c.AccountType).Kind() != reflect.String {
		t.Errorf("unexpected return type; expected: string, got: %T", reflect.ValueOf(c.AccountType).Kind())
	}
	if reflect.ValueOf(c.Commissions).Kind() != reflect.Bool {
		t.Errorf("unexpected return type; expected: bool, got: %T", reflect.ValueOf(c.Commissions).Kind())
	}
	if reflect.ValueOf(c.Recordings).Kind() != reflect.Bool {
		t.Errorf("unexpected return type; expected: bool, got: %T", reflect.ValueOf(c.Recordings).Kind())
	}
}
