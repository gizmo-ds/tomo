/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package models

import "time"

type UserStatus = string

const (
	UserStatusActive  UserStatus = "active"
	UserStatusJoinMe  UserStatus = "join me"
	UserStatusAskMe   UserStatus = "ask me"
	UserStatusBusy    UserStatus = "busy"
	UserStatusOffline UserStatus = "offline"
)

type DeveloperType = string

const (
	DeveloperTypeNone      DeveloperType = "none"
	DeveloperTypeTrusted   DeveloperType = "trusted"
	DeveloperTypeInternal  DeveloperType = "internal"
	DeveloperTypeModerator DeveloperType = "moderator"
)

type (
	LimitedUser struct {
		ID                             string        `json:"id"`
		Username                       string        `json:"username"`
		DisplayName                    string        `json:"displayName"`
		Bio                            string        `json:"bio"`
		BioLinks                       []string      `json:"bioLinks"`
		CurrentAvatarImageURL          string        `json:"currentAvatarImageUrl"`
		CurrentAvatarThumbnailImageURL string        `json:"currentAvatarThumbnailImageUrl"`
		LastPlatform                   string        `json:"last_platform"`
		Tags                           []string      `json:"tags"`
		DeveloperType                  DeveloperType `json:"developerType"`
		IsFriend                       bool          `json:"isFriend"`
	}

	CurrentUser struct {
		LimitedUser
		UserIcon                  string               `json:"userIcon"`
		ProfilePicOverride        string               `json:"profilePicOverride"`
		StatusDescription         string               `json:"statusDescription"`
		Badges                    []UserBadge          `json:"badges"`
		PastDisplayNames          []PastDisplayName    `json:"pastDisplayNames"`
		HasEmail                  bool                 `json:"hasEmail"`
		HasPendingEmail           bool                 `json:"hasPendingEmail"`
		ObfuscatedEmail           string               `json:"obfuscatedEmail"`
		ObfuscatedPendingEmail    string               `json:"obfuscatedPendingEmail"`
		EmailVerified             bool                 `json:"emailVerified"`
		HasBirthday               bool                 `json:"hasBirthday"`
		HideContentFilterSettings bool                 `json:"hideContentFilterSettings"`
		Unsubscribe               bool                 `json:"unsubscribe"`
		StatusHistory             []string             `json:"statusHistory"`
		StatusFirstTime           bool                 `json:"statusFirstTime"`
		Friends                   []string             `json:"friends"`
		FriendGroupNames          []string             `json:"friendGroupNames"`
		UserLanguage              string               `json:"userLanguage"`
		UserLanguageCode          string               `json:"userLanguageCode"`
		CurrentAvatarTags         []string             `json:"currentAvatarTags"`
		CurrentAvatar             string               `json:"currentAvatar"`
		CurrentAvatarAssetURL     string               `json:"currentAvatarAssetUrl"`
		FallbackAvatar            string               `json:"fallbackAvatar"`
		AccountDeletionDate       string               `json:"accountDeletionDate"`
		AccountDeletionLog        []AccountDeletionLog `json:"accountDeletionLog"`
		AcceptedTOSVersion        int                  `json:"acceptedTOSVersion"`
		AcceptedPrivacyVersion    int                  `json:"acceptedPrivacyVersion"`
		SteamID                   string               `json:"steamId"`
		SteamDetails              SteamDetails         `json:"steamDetails"`
		GoogleID                  string               `json:"googleId"`
		GoogleDetails             struct{}             `json:"googleDetails"`
		OculusID                  string               `json:"oculusId"`
		PicoID                    string               `json:"picoId"`
		ViveID                    string               `json:"viveId"`
		HasLoggedInFromClient     bool                 `json:"hasLoggedInFromClient"`
		HomeLocation              string               `json:"homeLocation"`
		TwoFactorAuthEnabled      bool                 `json:"twoFactorAuthEnabled"`
		TwoFactorAuthEnabledDate  time.Time            `json:"twoFactorAuthEnabledDate"`
		UpdatedAt                 time.Time            `json:"updated_at"`
		State                     string               `json:"state"`
		LastLogin                 time.Time            `json:"last_login"`
		AllowAvatarCopying        bool                 `json:"allowAvatarCopying"`
		Status                    string               `json:"status"`
		DateJoined                string               `json:"date_joined"`
		FriendKey                 string               `json:"friendKey"`
		LastActivity              time.Time            `json:"last_activity"`
		OnlineFriends             []string             `json:"onlineFriends"`
		ActiveFriends             []string             `json:"activeFriends"`
		Presence                  UserPresence         `json:"presence"`
		OfflineFriends            []string             `json:"offlineFriends"`
	}

	UserBadge struct {
		BadgeID          string    `json:"badgeId"`
		Showcased        bool      `json:"showcased"`
		BadgeName        string    `json:"badgeName"`
		BadgeDescription string    `json:"badgeDescription"`
		BadgeImageURL    string    `json:"badgeImageUrl"`
		Hidden           bool      `json:"hidden"`
		AssignedAt       time.Time `json:"assignedAt"`
		UpdatedAt        time.Time `json:"updatedAt"`
	}
	AccountDeletionLog struct {
		Message           string  `json:"message"`
		DeletionScheduled *string `json:"deletionScheduled"`
		DateTime          string  `json:"dateTime"`
	}
	SteamDetails struct {
		SteamID                  string `json:"steamid"`
		CommunityVisibilityState int    `json:"communityvisibilitystate"`
		ProfileState             int    `json:"profilestate"`
		PersonaName              string `json:"personaname"`
		ProfileUrl               string `json:"profileurl"`
		Avatar                   string `json:"avatar"`
		AvatarMedium             string `json:"avatarmedium"`
		AvatarFull               string `json:"avatarfull"`
		PersonaState             int    `json:"personastate"`
		RealName                 string `json:"realname"`
		PrimaryClanID            string `json:"primaryclanid"`
		TimeCreated              int    `json:"timecreated"`
		PersonaStateFlags        int    `json:"personastateflags"`
		AvatarHash               string `json:"avatarhash"`
		CommentPermission        int    `json:"commentpermission"`
		GameExtraInfo            string `json:"gameextrainfo"`
		GameID                   string `json:"gameid"`
	}
	PastDisplayName struct {
		DisplayName string    `json:"displayName"`
		UpdatedAt   time.Time `json:"updated_at"`
		Reverted    bool      `json:"reverted"`
	}
	UserPresence struct {
		InstanceType        string        `json:"instanceType"`
		AvatarThumbnail     string        `json:"avatarThumbnail"`
		TravelingToWorld    string        `json:"travelingToWorld"`
		World               string        `json:"world"`
		DisplayName         string        `json:"displayName"`
		Instance            string        `json:"instance"`
		Groups              []interface{} `json:"groups"`
		ID                  string        `json:"id"`
		ProfilePicOverride  string        `json:"profilePicOverride"`
		UserIcon            string        `json:"userIcon"`
		Platform            string        `json:"platform"`
		TravelingToInstance string        `json:"travelingToInstance"`
		Status              string        `json:"status"`
		DebugFlag           string        `json:"debugflag"`
		CurrentAvatarTags   string        `json:"currentAvatarTags"`
	}

	Friend struct {
		ID                string     `json:"id"`
		DisplayName       string     `json:"displayName"`
		Status            UserStatus `json:"status"`
		StatusDescription string     `json:"statusDescription"`
		Bio               string     `json:"bio"`
		// BioLinks           []string      `json:"bioLinks"`
		Location                       string        `json:"location"`
		UserIcon                       string        `json:"userIcon"`
		Tags                           []string      `json:"tags"`
		DeveloperType                  DeveloperType `json:"developerType"`
		ProfilePicOverride             string        `json:"profilePicOverride"`
		IsFriend                       bool          `json:"isFriend"`
		FriendKey                      string        `json:"friendKey"`
		LastPlatform                   string        `json:"last_platform"`
		FallbackAvatar                 string        `json:"fallbackAvatar"`
		CurrentAvatarImageURL          string        `json:"currentAvatarImageUrl"`
		CurrentAvatarThumbnailImageURL string        `json:"currentAvatarThumbnailImageUrl"`
	}
)
