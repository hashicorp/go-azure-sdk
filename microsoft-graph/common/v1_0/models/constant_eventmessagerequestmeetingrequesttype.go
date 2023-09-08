package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageRequestMeetingRequestType string

const (
	EventMessageRequestMeetingRequestTypefullUpdate          EventMessageRequestMeetingRequestType = "FullUpdate"
	EventMessageRequestMeetingRequestTypeinformationalUpdate EventMessageRequestMeetingRequestType = "InformationalUpdate"
	EventMessageRequestMeetingRequestTypenewMeetingRequest   EventMessageRequestMeetingRequestType = "NewMeetingRequest"
	EventMessageRequestMeetingRequestTypenone                EventMessageRequestMeetingRequestType = "None"
	EventMessageRequestMeetingRequestTypeoutdated            EventMessageRequestMeetingRequestType = "Outdated"
	EventMessageRequestMeetingRequestTypeprincipalWantsCopy  EventMessageRequestMeetingRequestType = "PrincipalWantsCopy"
	EventMessageRequestMeetingRequestTypesilentUpdate        EventMessageRequestMeetingRequestType = "SilentUpdate"
)

func PossibleValuesForEventMessageRequestMeetingRequestType() []string {
	return []string{
		string(EventMessageRequestMeetingRequestTypefullUpdate),
		string(EventMessageRequestMeetingRequestTypeinformationalUpdate),
		string(EventMessageRequestMeetingRequestTypenewMeetingRequest),
		string(EventMessageRequestMeetingRequestTypenone),
		string(EventMessageRequestMeetingRequestTypeoutdated),
		string(EventMessageRequestMeetingRequestTypeprincipalWantsCopy),
		string(EventMessageRequestMeetingRequestTypesilentUpdate),
	}
}

func parseEventMessageRequestMeetingRequestType(input string) (*EventMessageRequestMeetingRequestType, error) {
	vals := map[string]EventMessageRequestMeetingRequestType{
		"fullupdate":          EventMessageRequestMeetingRequestTypefullUpdate,
		"informationalupdate": EventMessageRequestMeetingRequestTypeinformationalUpdate,
		"newmeetingrequest":   EventMessageRequestMeetingRequestTypenewMeetingRequest,
		"none":                EventMessageRequestMeetingRequestTypenone,
		"outdated":            EventMessageRequestMeetingRequestTypeoutdated,
		"principalwantscopy":  EventMessageRequestMeetingRequestTypeprincipalWantsCopy,
		"silentupdate":        EventMessageRequestMeetingRequestTypesilentUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageRequestMeetingRequestType(input)
	return &out, nil
}
