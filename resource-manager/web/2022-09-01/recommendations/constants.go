package recommendations

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Channels string

const (
	ChannelsAll          Channels = "All"
	ChannelsApi          Channels = "Api"
	ChannelsEmail        Channels = "Email"
	ChannelsNotification Channels = "Notification"
	ChannelsWebhook      Channels = "Webhook"
)

func PossibleValuesForChannels() []string {
	return []string{
		string(ChannelsAll),
		string(ChannelsApi),
		string(ChannelsEmail),
		string(ChannelsNotification),
		string(ChannelsWebhook),
	}
}

func parseChannels(input string) (*Channels, error) {
	vals := map[string]Channels{
		"all":          ChannelsAll,
		"api":          ChannelsApi,
		"email":        ChannelsEmail,
		"notification": ChannelsNotification,
		"webhook":      ChannelsWebhook,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Channels(input)
	return &out, nil
}

type NotificationLevel string

const (
	NotificationLevelCritical            NotificationLevel = "Critical"
	NotificationLevelInformation         NotificationLevel = "Information"
	NotificationLevelNonUrgentSuggestion NotificationLevel = "NonUrgentSuggestion"
	NotificationLevelWarning             NotificationLevel = "Warning"
)

func PossibleValuesForNotificationLevel() []string {
	return []string{
		string(NotificationLevelCritical),
		string(NotificationLevelInformation),
		string(NotificationLevelNonUrgentSuggestion),
		string(NotificationLevelWarning),
	}
}

func parseNotificationLevel(input string) (*NotificationLevel, error) {
	vals := map[string]NotificationLevel{
		"critical":            NotificationLevelCritical,
		"information":         NotificationLevelInformation,
		"nonurgentsuggestion": NotificationLevelNonUrgentSuggestion,
		"warning":             NotificationLevelWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotificationLevel(input)
	return &out, nil
}

type ResourceScopeType string

const (
	ResourceScopeTypeServerFarm   ResourceScopeType = "ServerFarm"
	ResourceScopeTypeSubscription ResourceScopeType = "Subscription"
	ResourceScopeTypeWebSite      ResourceScopeType = "WebSite"
)

func PossibleValuesForResourceScopeType() []string {
	return []string{
		string(ResourceScopeTypeServerFarm),
		string(ResourceScopeTypeSubscription),
		string(ResourceScopeTypeWebSite),
	}
}

func parseResourceScopeType(input string) (*ResourceScopeType, error) {
	vals := map[string]ResourceScopeType{
		"serverfarm":   ResourceScopeTypeServerFarm,
		"subscription": ResourceScopeTypeSubscription,
		"website":      ResourceScopeTypeWebSite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceScopeType(input)
	return &out, nil
}
