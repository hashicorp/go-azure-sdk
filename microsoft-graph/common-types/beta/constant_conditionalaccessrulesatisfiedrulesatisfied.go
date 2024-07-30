package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessRuleSatisfiedRuleSatisfied string

const (
	ConditionalAccessRuleSatisfiedRuleSatisfied_Acr                               ConditionalAccessRuleSatisfiedRuleSatisfied = "acr"
	ConditionalAccessRuleSatisfiedRuleSatisfied_AllApps                           ConditionalAccessRuleSatisfiedRuleSatisfied = "allApps"
	ConditionalAccessRuleSatisfiedRuleSatisfied_AllDevicePlatforms                ConditionalAccessRuleSatisfiedRuleSatisfied = "allDevicePlatforms"
	ConditionalAccessRuleSatisfiedRuleSatisfied_AllDeviceStates                   ConditionalAccessRuleSatisfiedRuleSatisfied = "allDeviceStates"
	ConditionalAccessRuleSatisfiedRuleSatisfied_AllDevices                        ConditionalAccessRuleSatisfiedRuleSatisfied = "allDevices"
	ConditionalAccessRuleSatisfiedRuleSatisfied_AllLocations                      ConditionalAccessRuleSatisfiedRuleSatisfied = "allLocations"
	ConditionalAccessRuleSatisfiedRuleSatisfied_AllTrustedLocations               ConditionalAccessRuleSatisfiedRuleSatisfied = "allTrustedLocations"
	ConditionalAccessRuleSatisfiedRuleSatisfied_AllUsers                          ConditionalAccessRuleSatisfiedRuleSatisfied = "allUsers"
	ConditionalAccessRuleSatisfiedRuleSatisfied_AnonymizedIPAddress               ConditionalAccessRuleSatisfiedRuleSatisfied = "anonymizedIPAddress"
	ConditionalAccessRuleSatisfiedRuleSatisfied_AppFilter                         ConditionalAccessRuleSatisfiedRuleSatisfied = "appFilter"
	ConditionalAccessRuleSatisfiedRuleSatisfied_AppId                             ConditionalAccessRuleSatisfiedRuleSatisfied = "appId"
	ConditionalAccessRuleSatisfiedRuleSatisfied_AuthenticationTransfer            ConditionalAccessRuleSatisfiedRuleSatisfied = "authenticationTransfer"
	ConditionalAccessRuleSatisfiedRuleSatisfied_B2bCollaborationGuest             ConditionalAccessRuleSatisfiedRuleSatisfied = "b2bCollaborationGuest"
	ConditionalAccessRuleSatisfiedRuleSatisfied_B2bCollaborationMember            ConditionalAccessRuleSatisfiedRuleSatisfied = "b2bCollaborationMember"
	ConditionalAccessRuleSatisfiedRuleSatisfied_B2bDirectConnectUser              ConditionalAccessRuleSatisfiedRuleSatisfied = "b2bDirectConnectUser"
	ConditionalAccessRuleSatisfiedRuleSatisfied_DeviceCodeFlow                    ConditionalAccessRuleSatisfiedRuleSatisfied = "deviceCodeFlow"
	ConditionalAccessRuleSatisfiedRuleSatisfied_DeviceFilter                      ConditionalAccessRuleSatisfiedRuleSatisfied = "deviceFilter"
	ConditionalAccessRuleSatisfiedRuleSatisfied_DeviceFilterIncludeRuleNotMatched ConditionalAccessRuleSatisfiedRuleSatisfied = "deviceFilterIncludeRuleNotMatched"
	ConditionalAccessRuleSatisfiedRuleSatisfied_DevicePlatform                    ConditionalAccessRuleSatisfiedRuleSatisfied = "devicePlatform"
	ConditionalAccessRuleSatisfiedRuleSatisfied_DeviceState                       ConditionalAccessRuleSatisfiedRuleSatisfied = "deviceState"
	ConditionalAccessRuleSatisfiedRuleSatisfied_FirstPartyApps                    ConditionalAccessRuleSatisfiedRuleSatisfied = "firstPartyApps"
	ConditionalAccessRuleSatisfiedRuleSatisfied_GroupId                           ConditionalAccessRuleSatisfiedRuleSatisfied = "groupId"
	ConditionalAccessRuleSatisfiedRuleSatisfied_Guest                             ConditionalAccessRuleSatisfiedRuleSatisfied = "guest"
	ConditionalAccessRuleSatisfiedRuleSatisfied_InsideCorpnet                     ConditionalAccessRuleSatisfiedRuleSatisfied = "insideCorpnet"
	ConditionalAccessRuleSatisfiedRuleSatisfied_InsiderRisk                       ConditionalAccessRuleSatisfiedRuleSatisfied = "insiderRisk"
	ConditionalAccessRuleSatisfiedRuleSatisfied_InternalGuest                     ConditionalAccessRuleSatisfiedRuleSatisfied = "internalGuest"
	ConditionalAccessRuleSatisfiedRuleSatisfied_LocationId                        ConditionalAccessRuleSatisfiedRuleSatisfied = "locationId"
	ConditionalAccessRuleSatisfiedRuleSatisfied_MicrosoftAdminPortals             ConditionalAccessRuleSatisfiedRuleSatisfied = "microsoftAdminPortals"
	ConditionalAccessRuleSatisfiedRuleSatisfied_NationStateIPAddress              ConditionalAccessRuleSatisfiedRuleSatisfied = "nationStateIPAddress"
	ConditionalAccessRuleSatisfiedRuleSatisfied_Office365                         ConditionalAccessRuleSatisfiedRuleSatisfied = "office365"
	ConditionalAccessRuleSatisfiedRuleSatisfied_OtherExternalUser                 ConditionalAccessRuleSatisfiedRuleSatisfied = "otherExternalUser"
	ConditionalAccessRuleSatisfiedRuleSatisfied_RealTimeThreatIntelligence        ConditionalAccessRuleSatisfiedRuleSatisfied = "realTimeThreatIntelligence"
	ConditionalAccessRuleSatisfiedRuleSatisfied_RoleId                            ConditionalAccessRuleSatisfiedRuleSatisfied = "roleId"
	ConditionalAccessRuleSatisfiedRuleSatisfied_ServiceProvider                   ConditionalAccessRuleSatisfiedRuleSatisfied = "serviceProvider"
	ConditionalAccessRuleSatisfiedRuleSatisfied_UnfamiliarFeatures                ConditionalAccessRuleSatisfiedRuleSatisfied = "unfamiliarFeatures"
	ConditionalAccessRuleSatisfiedRuleSatisfied_UserId                            ConditionalAccessRuleSatisfiedRuleSatisfied = "userId"
)

func PossibleValuesForConditionalAccessRuleSatisfiedRuleSatisfied() []string {
	return []string{
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_Acr),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_AllApps),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_AllDevicePlatforms),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_AllDeviceStates),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_AllDevices),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_AllLocations),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_AllTrustedLocations),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_AllUsers),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_AnonymizedIPAddress),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_AppFilter),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_AppId),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_AuthenticationTransfer),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_B2bCollaborationGuest),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_B2bCollaborationMember),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_B2bDirectConnectUser),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_DeviceCodeFlow),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_DeviceFilter),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_DeviceFilterIncludeRuleNotMatched),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_DevicePlatform),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_DeviceState),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_FirstPartyApps),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_GroupId),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_Guest),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_InsideCorpnet),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_InsiderRisk),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_InternalGuest),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_LocationId),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_MicrosoftAdminPortals),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_NationStateIPAddress),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_Office365),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_OtherExternalUser),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_RealTimeThreatIntelligence),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_RoleId),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_ServiceProvider),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_UnfamiliarFeatures),
		string(ConditionalAccessRuleSatisfiedRuleSatisfied_UserId),
	}
}

func (s *ConditionalAccessRuleSatisfiedRuleSatisfied) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessRuleSatisfiedRuleSatisfied(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessRuleSatisfiedRuleSatisfied(input string) (*ConditionalAccessRuleSatisfiedRuleSatisfied, error) {
	vals := map[string]ConditionalAccessRuleSatisfiedRuleSatisfied{
		"acr":                               ConditionalAccessRuleSatisfiedRuleSatisfied_Acr,
		"allapps":                           ConditionalAccessRuleSatisfiedRuleSatisfied_AllApps,
		"alldeviceplatforms":                ConditionalAccessRuleSatisfiedRuleSatisfied_AllDevicePlatforms,
		"alldevicestates":                   ConditionalAccessRuleSatisfiedRuleSatisfied_AllDeviceStates,
		"alldevices":                        ConditionalAccessRuleSatisfiedRuleSatisfied_AllDevices,
		"alllocations":                      ConditionalAccessRuleSatisfiedRuleSatisfied_AllLocations,
		"alltrustedlocations":               ConditionalAccessRuleSatisfiedRuleSatisfied_AllTrustedLocations,
		"allusers":                          ConditionalAccessRuleSatisfiedRuleSatisfied_AllUsers,
		"anonymizedipaddress":               ConditionalAccessRuleSatisfiedRuleSatisfied_AnonymizedIPAddress,
		"appfilter":                         ConditionalAccessRuleSatisfiedRuleSatisfied_AppFilter,
		"appid":                             ConditionalAccessRuleSatisfiedRuleSatisfied_AppId,
		"authenticationtransfer":            ConditionalAccessRuleSatisfiedRuleSatisfied_AuthenticationTransfer,
		"b2bcollaborationguest":             ConditionalAccessRuleSatisfiedRuleSatisfied_B2bCollaborationGuest,
		"b2bcollaborationmember":            ConditionalAccessRuleSatisfiedRuleSatisfied_B2bCollaborationMember,
		"b2bdirectconnectuser":              ConditionalAccessRuleSatisfiedRuleSatisfied_B2bDirectConnectUser,
		"devicecodeflow":                    ConditionalAccessRuleSatisfiedRuleSatisfied_DeviceCodeFlow,
		"devicefilter":                      ConditionalAccessRuleSatisfiedRuleSatisfied_DeviceFilter,
		"devicefilterincluderulenotmatched": ConditionalAccessRuleSatisfiedRuleSatisfied_DeviceFilterIncludeRuleNotMatched,
		"deviceplatform":                    ConditionalAccessRuleSatisfiedRuleSatisfied_DevicePlatform,
		"devicestate":                       ConditionalAccessRuleSatisfiedRuleSatisfied_DeviceState,
		"firstpartyapps":                    ConditionalAccessRuleSatisfiedRuleSatisfied_FirstPartyApps,
		"groupid":                           ConditionalAccessRuleSatisfiedRuleSatisfied_GroupId,
		"guest":                             ConditionalAccessRuleSatisfiedRuleSatisfied_Guest,
		"insidecorpnet":                     ConditionalAccessRuleSatisfiedRuleSatisfied_InsideCorpnet,
		"insiderrisk":                       ConditionalAccessRuleSatisfiedRuleSatisfied_InsiderRisk,
		"internalguest":                     ConditionalAccessRuleSatisfiedRuleSatisfied_InternalGuest,
		"locationid":                        ConditionalAccessRuleSatisfiedRuleSatisfied_LocationId,
		"microsoftadminportals":             ConditionalAccessRuleSatisfiedRuleSatisfied_MicrosoftAdminPortals,
		"nationstateipaddress":              ConditionalAccessRuleSatisfiedRuleSatisfied_NationStateIPAddress,
		"office365":                         ConditionalAccessRuleSatisfiedRuleSatisfied_Office365,
		"otherexternaluser":                 ConditionalAccessRuleSatisfiedRuleSatisfied_OtherExternalUser,
		"realtimethreatintelligence":        ConditionalAccessRuleSatisfiedRuleSatisfied_RealTimeThreatIntelligence,
		"roleid":                            ConditionalAccessRuleSatisfiedRuleSatisfied_RoleId,
		"serviceprovider":                   ConditionalAccessRuleSatisfiedRuleSatisfied_ServiceProvider,
		"unfamiliarfeatures":                ConditionalAccessRuleSatisfiedRuleSatisfied_UnfamiliarFeatures,
		"userid":                            ConditionalAccessRuleSatisfiedRuleSatisfied_UserId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessRuleSatisfiedRuleSatisfied(input)
	return &out, nil
}
