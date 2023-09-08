package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsServiceUserAgentRole string

const (
	CallRecordsServiceUserAgentRoleaudioTeleconferencerController                          CallRecordsServiceUserAgentRole = "AudioTeleconferencerController"
	CallRecordsServiceUserAgentRoleconferencingAnnouncementService                         CallRecordsServiceUserAgentRole = "ConferencingAnnouncementService"
	CallRecordsServiceUserAgentRoleconferencingAttendant                                   CallRecordsServiceUserAgentRole = "ConferencingAttendant"
	CallRecordsServiceUserAgentRolecustomBot                                               CallRecordsServiceUserAgentRole = "CustomBot"
	CallRecordsServiceUserAgentRoleexchangeUnifiedMessagingService                         CallRecordsServiceUserAgentRole = "ExchangeUnifiedMessagingService"
	CallRecordsServiceUserAgentRolegateway                                                 CallRecordsServiceUserAgentRole = "Gateway"
	CallRecordsServiceUserAgentRolemediaController                                         CallRecordsServiceUserAgentRole = "MediaController"
	CallRecordsServiceUserAgentRolemediationServer                                         CallRecordsServiceUserAgentRole = "MediationServer"
	CallRecordsServiceUserAgentRolemediationServerCloudConnectorEdition                    CallRecordsServiceUserAgentRole = "MediationServerCloudConnectorEdition"
	CallRecordsServiceUserAgentRoleresponseGroupService                                    CallRecordsServiceUserAgentRole = "ResponseGroupService"
	CallRecordsServiceUserAgentRoleresponseGroupServiceAnnouncementService                 CallRecordsServiceUserAgentRole = "ResponseGroupServiceAnnouncementService"
	CallRecordsServiceUserAgentRoleskypeForBusinessApplicationSharingMcu                   CallRecordsServiceUserAgentRole = "SkypeForBusinessApplicationSharingMcu"
	CallRecordsServiceUserAgentRoleskypeForBusinessAttendant                               CallRecordsServiceUserAgentRole = "SkypeForBusinessAttendant"
	CallRecordsServiceUserAgentRoleskypeForBusinessAudioVideoMcu                           CallRecordsServiceUserAgentRole = "SkypeForBusinessAudioVideoMcu"
	CallRecordsServiceUserAgentRoleskypeForBusinessAutoAttendant                           CallRecordsServiceUserAgentRole = "SkypeForBusinessAutoAttendant"
	CallRecordsServiceUserAgentRoleskypeForBusinessCallQueues                              CallRecordsServiceUserAgentRole = "SkypeForBusinessCallQueues"
	CallRecordsServiceUserAgentRoleskypeForBusinessMicrosoftTeamsGateway                   CallRecordsServiceUserAgentRole = "SkypeForBusinessMicrosoftTeamsGateway"
	CallRecordsServiceUserAgentRoleskypeForBusinessUnifiedCommunicationApplicationPlatform CallRecordsServiceUserAgentRole = "SkypeForBusinessUnifiedCommunicationApplicationPlatform"
	CallRecordsServiceUserAgentRoleskypeTranslator                                         CallRecordsServiceUserAgentRole = "SkypeTranslator"
	CallRecordsServiceUserAgentRoleunknown                                                 CallRecordsServiceUserAgentRole = "Unknown"
	CallRecordsServiceUserAgentRolevoicemail                                               CallRecordsServiceUserAgentRole = "Voicemail"
)

func PossibleValuesForCallRecordsServiceUserAgentRole() []string {
	return []string{
		string(CallRecordsServiceUserAgentRoleaudioTeleconferencerController),
		string(CallRecordsServiceUserAgentRoleconferencingAnnouncementService),
		string(CallRecordsServiceUserAgentRoleconferencingAttendant),
		string(CallRecordsServiceUserAgentRolecustomBot),
		string(CallRecordsServiceUserAgentRoleexchangeUnifiedMessagingService),
		string(CallRecordsServiceUserAgentRolegateway),
		string(CallRecordsServiceUserAgentRolemediaController),
		string(CallRecordsServiceUserAgentRolemediationServer),
		string(CallRecordsServiceUserAgentRolemediationServerCloudConnectorEdition),
		string(CallRecordsServiceUserAgentRoleresponseGroupService),
		string(CallRecordsServiceUserAgentRoleresponseGroupServiceAnnouncementService),
		string(CallRecordsServiceUserAgentRoleskypeForBusinessApplicationSharingMcu),
		string(CallRecordsServiceUserAgentRoleskypeForBusinessAttendant),
		string(CallRecordsServiceUserAgentRoleskypeForBusinessAudioVideoMcu),
		string(CallRecordsServiceUserAgentRoleskypeForBusinessAutoAttendant),
		string(CallRecordsServiceUserAgentRoleskypeForBusinessCallQueues),
		string(CallRecordsServiceUserAgentRoleskypeForBusinessMicrosoftTeamsGateway),
		string(CallRecordsServiceUserAgentRoleskypeForBusinessUnifiedCommunicationApplicationPlatform),
		string(CallRecordsServiceUserAgentRoleskypeTranslator),
		string(CallRecordsServiceUserAgentRoleunknown),
		string(CallRecordsServiceUserAgentRolevoicemail),
	}
}

func parseCallRecordsServiceUserAgentRole(input string) (*CallRecordsServiceUserAgentRole, error) {
	vals := map[string]CallRecordsServiceUserAgentRole{
		"audioteleconferencercontroller":                          CallRecordsServiceUserAgentRoleaudioTeleconferencerController,
		"conferencingannouncementservice":                         CallRecordsServiceUserAgentRoleconferencingAnnouncementService,
		"conferencingattendant":                                   CallRecordsServiceUserAgentRoleconferencingAttendant,
		"custombot":                                               CallRecordsServiceUserAgentRolecustomBot,
		"exchangeunifiedmessagingservice":                         CallRecordsServiceUserAgentRoleexchangeUnifiedMessagingService,
		"gateway":                                                 CallRecordsServiceUserAgentRolegateway,
		"mediacontroller":                                         CallRecordsServiceUserAgentRolemediaController,
		"mediationserver":                                         CallRecordsServiceUserAgentRolemediationServer,
		"mediationservercloudconnectoredition":                    CallRecordsServiceUserAgentRolemediationServerCloudConnectorEdition,
		"responsegroupservice":                                    CallRecordsServiceUserAgentRoleresponseGroupService,
		"responsegroupserviceannouncementservice":                 CallRecordsServiceUserAgentRoleresponseGroupServiceAnnouncementService,
		"skypeforbusinessapplicationsharingmcu":                   CallRecordsServiceUserAgentRoleskypeForBusinessApplicationSharingMcu,
		"skypeforbusinessattendant":                               CallRecordsServiceUserAgentRoleskypeForBusinessAttendant,
		"skypeforbusinessaudiovideomcu":                           CallRecordsServiceUserAgentRoleskypeForBusinessAudioVideoMcu,
		"skypeforbusinessautoattendant":                           CallRecordsServiceUserAgentRoleskypeForBusinessAutoAttendant,
		"skypeforbusinesscallqueues":                              CallRecordsServiceUserAgentRoleskypeForBusinessCallQueues,
		"skypeforbusinessmicrosoftteamsgateway":                   CallRecordsServiceUserAgentRoleskypeForBusinessMicrosoftTeamsGateway,
		"skypeforbusinessunifiedcommunicationapplicationplatform": CallRecordsServiceUserAgentRoleskypeForBusinessUnifiedCommunicationApplicationPlatform,
		"skypetranslator":                                         CallRecordsServiceUserAgentRoleskypeTranslator,
		"unknown":                                                 CallRecordsServiceUserAgentRoleunknown,
		"voicemail":                                               CallRecordsServiceUserAgentRolevoicemail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsServiceUserAgentRole(input)
	return &out, nil
}
