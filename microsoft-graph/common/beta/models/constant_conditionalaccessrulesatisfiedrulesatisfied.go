package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessRuleSatisfiedRuleSatisfied string

const (
	ConditionalAccessRuleSatisfiedRuleSatisfiedacr                               ConditionalAccessRuleSatisfiedRuleSatisfied = "Acr"
	ConditionalAccessRuleSatisfiedRuleSatisfiedallApps                           ConditionalAccessRuleSatisfiedRuleSatisfied = "AllApps"
	ConditionalAccessRuleSatisfiedRuleSatisfiedallDevicePlatforms                ConditionalAccessRuleSatisfiedRuleSatisfied = "AllDevicePlatforms"
	ConditionalAccessRuleSatisfiedRuleSatisfiedallDeviceStates                   ConditionalAccessRuleSatisfiedRuleSatisfied = "AllDeviceStates"
	ConditionalAccessRuleSatisfiedRuleSatisfiedallDevices                        ConditionalAccessRuleSatisfiedRuleSatisfied = "AllDevices"
	ConditionalAccessRuleSatisfiedRuleSatisfiedallLocations                      ConditionalAccessRuleSatisfiedRuleSatisfied = "AllLocations"
	ConditionalAccessRuleSatisfiedRuleSatisfiedallTrustedLocations               ConditionalAccessRuleSatisfiedRuleSatisfied = "AllTrustedLocations"
	ConditionalAccessRuleSatisfiedRuleSatisfiedallUsers                          ConditionalAccessRuleSatisfiedRuleSatisfied = "AllUsers"
	ConditionalAccessRuleSatisfiedRuleSatisfiedanonymizedIPAddress               ConditionalAccessRuleSatisfiedRuleSatisfied = "AnonymizedIPAddress"
	ConditionalAccessRuleSatisfiedRuleSatisfiedappFilter                         ConditionalAccessRuleSatisfiedRuleSatisfied = "AppFilter"
	ConditionalAccessRuleSatisfiedRuleSatisfiedappId                             ConditionalAccessRuleSatisfiedRuleSatisfied = "AppId"
	ConditionalAccessRuleSatisfiedRuleSatisfiedauthenticationTransfer            ConditionalAccessRuleSatisfiedRuleSatisfied = "AuthenticationTransfer"
	ConditionalAccessRuleSatisfiedRuleSatisfiedb2bCollaborationGuest             ConditionalAccessRuleSatisfiedRuleSatisfied = "B2bCollaborationGuest"
	ConditionalAccessRuleSatisfiedRuleSatisfiedb2bCollaborationMember            ConditionalAccessRuleSatisfiedRuleSatisfied = "B2bCollaborationMember"
	ConditionalAccessRuleSatisfiedRuleSatisfiedb2bDirectConnectUser              ConditionalAccessRuleSatisfiedRuleSatisfied = "B2bDirectConnectUser"
	ConditionalAccessRuleSatisfiedRuleSatisfieddeviceCodeFlow                    ConditionalAccessRuleSatisfiedRuleSatisfied = "DeviceCodeFlow"
	ConditionalAccessRuleSatisfiedRuleSatisfieddeviceFilter                      ConditionalAccessRuleSatisfiedRuleSatisfied = "DeviceFilter"
	ConditionalAccessRuleSatisfiedRuleSatisfieddeviceFilterIncludeRuleNotMatched ConditionalAccessRuleSatisfiedRuleSatisfied = "DeviceFilterIncludeRuleNotMatched"
	ConditionalAccessRuleSatisfiedRuleSatisfieddevicePlatform                    ConditionalAccessRuleSatisfiedRuleSatisfied = "DevicePlatform"
	ConditionalAccessRuleSatisfiedRuleSatisfieddeviceState                       ConditionalAccessRuleSatisfiedRuleSatisfied = "DeviceState"
	ConditionalAccessRuleSatisfiedRuleSatisfiedfirstPartyApps                    ConditionalAccessRuleSatisfiedRuleSatisfied = "FirstPartyApps"
	ConditionalAccessRuleSatisfiedRuleSatisfiedgroupId                           ConditionalAccessRuleSatisfiedRuleSatisfied = "GroupId"
	ConditionalAccessRuleSatisfiedRuleSatisfiedguest                             ConditionalAccessRuleSatisfiedRuleSatisfied = "Guest"
	ConditionalAccessRuleSatisfiedRuleSatisfiedinsideCorpnet                     ConditionalAccessRuleSatisfiedRuleSatisfied = "InsideCorpnet"
	ConditionalAccessRuleSatisfiedRuleSatisfiedinsiderRisk                       ConditionalAccessRuleSatisfiedRuleSatisfied = "InsiderRisk"
	ConditionalAccessRuleSatisfiedRuleSatisfiedinternalGuest                     ConditionalAccessRuleSatisfiedRuleSatisfied = "InternalGuest"
	ConditionalAccessRuleSatisfiedRuleSatisfiedlocationId                        ConditionalAccessRuleSatisfiedRuleSatisfied = "LocationId"
	ConditionalAccessRuleSatisfiedRuleSatisfiedmicrosoftAdminPortals             ConditionalAccessRuleSatisfiedRuleSatisfied = "MicrosoftAdminPortals"
	ConditionalAccessRuleSatisfiedRuleSatisfiednationStateIPAddress              ConditionalAccessRuleSatisfiedRuleSatisfied = "NationStateIPAddress"
	ConditionalAccessRuleSatisfiedRuleSatisfiedoffice365                         ConditionalAccessRuleSatisfiedRuleSatisfied = "Office365"
	ConditionalAccessRuleSatisfiedRuleSatisfiedotherExternalUser                 ConditionalAccessRuleSatisfiedRuleSatisfied = "OtherExternalUser"
	ConditionalAccessRuleSatisfiedRuleSatisfiedrealTimeThreatIntelligence        ConditionalAccessRuleSatisfiedRuleSatisfied = "RealTimeThreatIntelligence"
	ConditionalAccessRuleSatisfiedRuleSatisfiedroleId                            ConditionalAccessRuleSatisfiedRuleSatisfied = "RoleId"
	ConditionalAccessRuleSatisfiedRuleSatisfiedserviceProvider                   ConditionalAccessRuleSatisfiedRuleSatisfied = "ServiceProvider"
	ConditionalAccessRuleSatisfiedRuleSatisfiedunfamiliarFeatures                ConditionalAccessRuleSatisfiedRuleSatisfied = "UnfamiliarFeatures"
	ConditionalAccessRuleSatisfiedRuleSatisfieduserId                            ConditionalAccessRuleSatisfiedRuleSatisfied = "UserId"
)

func PossibleValuesForConditionalAccessRuleSatisfiedRuleSatisfied() []string {
	return []string{
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedacr),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedallApps),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedallDevicePlatforms),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedallDeviceStates),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedallDevices),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedallLocations),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedallTrustedLocations),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedallUsers),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedanonymizedIPAddress),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedappFilter),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedappId),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedauthenticationTransfer),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedb2bCollaborationGuest),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedb2bCollaborationMember),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedb2bDirectConnectUser),
		string(ConditionalAccessRuleSatisfiedRuleSatisfieddeviceCodeFlow),
		string(ConditionalAccessRuleSatisfiedRuleSatisfieddeviceFilter),
		string(ConditionalAccessRuleSatisfiedRuleSatisfieddeviceFilterIncludeRuleNotMatched),
		string(ConditionalAccessRuleSatisfiedRuleSatisfieddevicePlatform),
		string(ConditionalAccessRuleSatisfiedRuleSatisfieddeviceState),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedfirstPartyApps),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedgroupId),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedguest),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedinsideCorpnet),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedinsiderRisk),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedinternalGuest),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedlocationId),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedmicrosoftAdminPortals),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiednationStateIPAddress),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedoffice365),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedotherExternalUser),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedrealTimeThreatIntelligence),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedroleId),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedserviceProvider),
		string(ConditionalAccessRuleSatisfiedRuleSatisfiedunfamiliarFeatures),
		string(ConditionalAccessRuleSatisfiedRuleSatisfieduserId),
	}
}

func parseConditionalAccessRuleSatisfiedRuleSatisfied(input string) (*ConditionalAccessRuleSatisfiedRuleSatisfied, error) {
	vals := map[string]ConditionalAccessRuleSatisfiedRuleSatisfied{
		"acr":                               ConditionalAccessRuleSatisfiedRuleSatisfiedacr,
		"allapps":                           ConditionalAccessRuleSatisfiedRuleSatisfiedallApps,
		"alldeviceplatforms":                ConditionalAccessRuleSatisfiedRuleSatisfiedallDevicePlatforms,
		"alldevicestates":                   ConditionalAccessRuleSatisfiedRuleSatisfiedallDeviceStates,
		"alldevices":                        ConditionalAccessRuleSatisfiedRuleSatisfiedallDevices,
		"alllocations":                      ConditionalAccessRuleSatisfiedRuleSatisfiedallLocations,
		"alltrustedlocations":               ConditionalAccessRuleSatisfiedRuleSatisfiedallTrustedLocations,
		"allusers":                          ConditionalAccessRuleSatisfiedRuleSatisfiedallUsers,
		"anonymizedipaddress":               ConditionalAccessRuleSatisfiedRuleSatisfiedanonymizedIPAddress,
		"appfilter":                         ConditionalAccessRuleSatisfiedRuleSatisfiedappFilter,
		"appid":                             ConditionalAccessRuleSatisfiedRuleSatisfiedappId,
		"authenticationtransfer":            ConditionalAccessRuleSatisfiedRuleSatisfiedauthenticationTransfer,
		"b2bcollaborationguest":             ConditionalAccessRuleSatisfiedRuleSatisfiedb2bCollaborationGuest,
		"b2bcollaborationmember":            ConditionalAccessRuleSatisfiedRuleSatisfiedb2bCollaborationMember,
		"b2bdirectconnectuser":              ConditionalAccessRuleSatisfiedRuleSatisfiedb2bDirectConnectUser,
		"devicecodeflow":                    ConditionalAccessRuleSatisfiedRuleSatisfieddeviceCodeFlow,
		"devicefilter":                      ConditionalAccessRuleSatisfiedRuleSatisfieddeviceFilter,
		"devicefilterincluderulenotmatched": ConditionalAccessRuleSatisfiedRuleSatisfieddeviceFilterIncludeRuleNotMatched,
		"deviceplatform":                    ConditionalAccessRuleSatisfiedRuleSatisfieddevicePlatform,
		"devicestate":                       ConditionalAccessRuleSatisfiedRuleSatisfieddeviceState,
		"firstpartyapps":                    ConditionalAccessRuleSatisfiedRuleSatisfiedfirstPartyApps,
		"groupid":                           ConditionalAccessRuleSatisfiedRuleSatisfiedgroupId,
		"guest":                             ConditionalAccessRuleSatisfiedRuleSatisfiedguest,
		"insidecorpnet":                     ConditionalAccessRuleSatisfiedRuleSatisfiedinsideCorpnet,
		"insiderrisk":                       ConditionalAccessRuleSatisfiedRuleSatisfiedinsiderRisk,
		"internalguest":                     ConditionalAccessRuleSatisfiedRuleSatisfiedinternalGuest,
		"locationid":                        ConditionalAccessRuleSatisfiedRuleSatisfiedlocationId,
		"microsoftadminportals":             ConditionalAccessRuleSatisfiedRuleSatisfiedmicrosoftAdminPortals,
		"nationstateipaddress":              ConditionalAccessRuleSatisfiedRuleSatisfiednationStateIPAddress,
		"office365":                         ConditionalAccessRuleSatisfiedRuleSatisfiedoffice365,
		"otherexternaluser":                 ConditionalAccessRuleSatisfiedRuleSatisfiedotherExternalUser,
		"realtimethreatintelligence":        ConditionalAccessRuleSatisfiedRuleSatisfiedrealTimeThreatIntelligence,
		"roleid":                            ConditionalAccessRuleSatisfiedRuleSatisfiedroleId,
		"serviceprovider":                   ConditionalAccessRuleSatisfiedRuleSatisfiedserviceProvider,
		"unfamiliarfeatures":                ConditionalAccessRuleSatisfiedRuleSatisfiedunfamiliarFeatures,
		"userid":                            ConditionalAccessRuleSatisfiedRuleSatisfieduserId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessRuleSatisfiedRuleSatisfied(input)
	return &out, nil
}
