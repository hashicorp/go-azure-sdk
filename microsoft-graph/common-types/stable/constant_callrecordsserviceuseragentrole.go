package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsServiceUserAgentRole string

const (
	CallRecordsServiceUserAgentRole_AudioTeleconferencerController                          CallRecordsServiceUserAgentRole = "audioTeleconferencerController"
	CallRecordsServiceUserAgentRole_ConferencingAnnouncementService                         CallRecordsServiceUserAgentRole = "conferencingAnnouncementService"
	CallRecordsServiceUserAgentRole_ConferencingAttendant                                   CallRecordsServiceUserAgentRole = "conferencingAttendant"
	CallRecordsServiceUserAgentRole_CustomBot                                               CallRecordsServiceUserAgentRole = "customBot"
	CallRecordsServiceUserAgentRole_ExchangeUnifiedMessagingService                         CallRecordsServiceUserAgentRole = "exchangeUnifiedMessagingService"
	CallRecordsServiceUserAgentRole_Gateway                                                 CallRecordsServiceUserAgentRole = "gateway"
	CallRecordsServiceUserAgentRole_MediaController                                         CallRecordsServiceUserAgentRole = "mediaController"
	CallRecordsServiceUserAgentRole_MediationServer                                         CallRecordsServiceUserAgentRole = "mediationServer"
	CallRecordsServiceUserAgentRole_MediationServerCloudConnectorEdition                    CallRecordsServiceUserAgentRole = "mediationServerCloudConnectorEdition"
	CallRecordsServiceUserAgentRole_ResponseGroupService                                    CallRecordsServiceUserAgentRole = "responseGroupService"
	CallRecordsServiceUserAgentRole_ResponseGroupServiceAnnouncementService                 CallRecordsServiceUserAgentRole = "responseGroupServiceAnnouncementService"
	CallRecordsServiceUserAgentRole_SkypeForBusinessApplicationSharingMcu                   CallRecordsServiceUserAgentRole = "skypeForBusinessApplicationSharingMcu"
	CallRecordsServiceUserAgentRole_SkypeForBusinessAttendant                               CallRecordsServiceUserAgentRole = "skypeForBusinessAttendant"
	CallRecordsServiceUserAgentRole_SkypeForBusinessAudioVideoMcu                           CallRecordsServiceUserAgentRole = "skypeForBusinessAudioVideoMcu"
	CallRecordsServiceUserAgentRole_SkypeForBusinessAutoAttendant                           CallRecordsServiceUserAgentRole = "skypeForBusinessAutoAttendant"
	CallRecordsServiceUserAgentRole_SkypeForBusinessCallQueues                              CallRecordsServiceUserAgentRole = "skypeForBusinessCallQueues"
	CallRecordsServiceUserAgentRole_SkypeForBusinessMicrosoftTeamsGateway                   CallRecordsServiceUserAgentRole = "skypeForBusinessMicrosoftTeamsGateway"
	CallRecordsServiceUserAgentRole_SkypeForBusinessUnifiedCommunicationApplicationPlatform CallRecordsServiceUserAgentRole = "skypeForBusinessUnifiedCommunicationApplicationPlatform"
	CallRecordsServiceUserAgentRole_SkypeTranslator                                         CallRecordsServiceUserAgentRole = "skypeTranslator"
	CallRecordsServiceUserAgentRole_Unknown                                                 CallRecordsServiceUserAgentRole = "unknown"
	CallRecordsServiceUserAgentRole_Voicemail                                               CallRecordsServiceUserAgentRole = "voicemail"
)

func PossibleValuesForCallRecordsServiceUserAgentRole() []string {
	return []string{
		string(CallRecordsServiceUserAgentRole_AudioTeleconferencerController),
		string(CallRecordsServiceUserAgentRole_ConferencingAnnouncementService),
		string(CallRecordsServiceUserAgentRole_ConferencingAttendant),
		string(CallRecordsServiceUserAgentRole_CustomBot),
		string(CallRecordsServiceUserAgentRole_ExchangeUnifiedMessagingService),
		string(CallRecordsServiceUserAgentRole_Gateway),
		string(CallRecordsServiceUserAgentRole_MediaController),
		string(CallRecordsServiceUserAgentRole_MediationServer),
		string(CallRecordsServiceUserAgentRole_MediationServerCloudConnectorEdition),
		string(CallRecordsServiceUserAgentRole_ResponseGroupService),
		string(CallRecordsServiceUserAgentRole_ResponseGroupServiceAnnouncementService),
		string(CallRecordsServiceUserAgentRole_SkypeForBusinessApplicationSharingMcu),
		string(CallRecordsServiceUserAgentRole_SkypeForBusinessAttendant),
		string(CallRecordsServiceUserAgentRole_SkypeForBusinessAudioVideoMcu),
		string(CallRecordsServiceUserAgentRole_SkypeForBusinessAutoAttendant),
		string(CallRecordsServiceUserAgentRole_SkypeForBusinessCallQueues),
		string(CallRecordsServiceUserAgentRole_SkypeForBusinessMicrosoftTeamsGateway),
		string(CallRecordsServiceUserAgentRole_SkypeForBusinessUnifiedCommunicationApplicationPlatform),
		string(CallRecordsServiceUserAgentRole_SkypeTranslator),
		string(CallRecordsServiceUserAgentRole_Unknown),
		string(CallRecordsServiceUserAgentRole_Voicemail),
	}
}

func (s *CallRecordsServiceUserAgentRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsServiceUserAgentRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsServiceUserAgentRole(input string) (*CallRecordsServiceUserAgentRole, error) {
	vals := map[string]CallRecordsServiceUserAgentRole{
		"audioteleconferencercontroller":                          CallRecordsServiceUserAgentRole_AudioTeleconferencerController,
		"conferencingannouncementservice":                         CallRecordsServiceUserAgentRole_ConferencingAnnouncementService,
		"conferencingattendant":                                   CallRecordsServiceUserAgentRole_ConferencingAttendant,
		"custombot":                                               CallRecordsServiceUserAgentRole_CustomBot,
		"exchangeunifiedmessagingservice":                         CallRecordsServiceUserAgentRole_ExchangeUnifiedMessagingService,
		"gateway":                                                 CallRecordsServiceUserAgentRole_Gateway,
		"mediacontroller":                                         CallRecordsServiceUserAgentRole_MediaController,
		"mediationserver":                                         CallRecordsServiceUserAgentRole_MediationServer,
		"mediationservercloudconnectoredition":                    CallRecordsServiceUserAgentRole_MediationServerCloudConnectorEdition,
		"responsegroupservice":                                    CallRecordsServiceUserAgentRole_ResponseGroupService,
		"responsegroupserviceannouncementservice":                 CallRecordsServiceUserAgentRole_ResponseGroupServiceAnnouncementService,
		"skypeforbusinessapplicationsharingmcu":                   CallRecordsServiceUserAgentRole_SkypeForBusinessApplicationSharingMcu,
		"skypeforbusinessattendant":                               CallRecordsServiceUserAgentRole_SkypeForBusinessAttendant,
		"skypeforbusinessaudiovideomcu":                           CallRecordsServiceUserAgentRole_SkypeForBusinessAudioVideoMcu,
		"skypeforbusinessautoattendant":                           CallRecordsServiceUserAgentRole_SkypeForBusinessAutoAttendant,
		"skypeforbusinesscallqueues":                              CallRecordsServiceUserAgentRole_SkypeForBusinessCallQueues,
		"skypeforbusinessmicrosoftteamsgateway":                   CallRecordsServiceUserAgentRole_SkypeForBusinessMicrosoftTeamsGateway,
		"skypeforbusinessunifiedcommunicationapplicationplatform": CallRecordsServiceUserAgentRole_SkypeForBusinessUnifiedCommunicationApplicationPlatform,
		"skypetranslator":                                         CallRecordsServiceUserAgentRole_SkypeTranslator,
		"unknown":                                                 CallRecordsServiceUserAgentRole_Unknown,
		"voicemail":                                               CallRecordsServiceUserAgentRole_Voicemail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsServiceUserAgentRole(input)
	return &out, nil
}
