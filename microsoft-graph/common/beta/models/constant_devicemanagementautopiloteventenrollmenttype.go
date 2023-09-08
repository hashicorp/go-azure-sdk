package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotEventEnrollmentType string

const (
	DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedUsingDeviceAuthWithAutopilotProfile    DeviceManagementAutopilotEventEnrollmentType = "AzureADJoinedUsingDeviceAuthWithAutopilotProfile"
	DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedUsingDeviceAuthWithoutAutopilotProfile DeviceManagementAutopilotEventEnrollmentType = "AzureADJoinedUsingDeviceAuthWithoutAutopilotProfile"
	DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedWithAutopilotProfile                   DeviceManagementAutopilotEventEnrollmentType = "AzureADJoinedWithAutopilotProfile"
	DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedWithOfflineAutopilotProfile            DeviceManagementAutopilotEventEnrollmentType = "AzureADJoinedWithOfflineAutopilotProfile"
	DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedWithWhiteGlove                         DeviceManagementAutopilotEventEnrollmentType = "AzureADJoinedWithWhiteGlove"
	DeviceManagementAutopilotEventEnrollmentTypeofflineDomainJoined                                 DeviceManagementAutopilotEventEnrollmentType = "OfflineDomainJoined"
	DeviceManagementAutopilotEventEnrollmentTypeofflineDomainJoinedWithOfflineAutopilotProfile      DeviceManagementAutopilotEventEnrollmentType = "OfflineDomainJoinedWithOfflineAutopilotProfile"
	DeviceManagementAutopilotEventEnrollmentTypeofflineDomainJoinedWithWhiteGlove                   DeviceManagementAutopilotEventEnrollmentType = "OfflineDomainJoinedWithWhiteGlove"
	DeviceManagementAutopilotEventEnrollmentTypeunknown                                             DeviceManagementAutopilotEventEnrollmentType = "Unknown"
)

func PossibleValuesForDeviceManagementAutopilotEventEnrollmentType() []string {
	return []string{
		string(DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedUsingDeviceAuthWithAutopilotProfile),
		string(DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedUsingDeviceAuthWithoutAutopilotProfile),
		string(DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedWithAutopilotProfile),
		string(DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedWithOfflineAutopilotProfile),
		string(DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedWithWhiteGlove),
		string(DeviceManagementAutopilotEventEnrollmentTypeofflineDomainJoined),
		string(DeviceManagementAutopilotEventEnrollmentTypeofflineDomainJoinedWithOfflineAutopilotProfile),
		string(DeviceManagementAutopilotEventEnrollmentTypeofflineDomainJoinedWithWhiteGlove),
		string(DeviceManagementAutopilotEventEnrollmentTypeunknown),
	}
}

func parseDeviceManagementAutopilotEventEnrollmentType(input string) (*DeviceManagementAutopilotEventEnrollmentType, error) {
	vals := map[string]DeviceManagementAutopilotEventEnrollmentType{
		"azureadjoinedusingdeviceauthwithautopilotprofile":    DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedUsingDeviceAuthWithAutopilotProfile,
		"azureadjoinedusingdeviceauthwithoutautopilotprofile": DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedUsingDeviceAuthWithoutAutopilotProfile,
		"azureadjoinedwithautopilotprofile":                   DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedWithAutopilotProfile,
		"azureadjoinedwithofflineautopilotprofile":            DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedWithOfflineAutopilotProfile,
		"azureadjoinedwithwhiteglove":                         DeviceManagementAutopilotEventEnrollmentTypeazureADJoinedWithWhiteGlove,
		"offlinedomainjoined":                                 DeviceManagementAutopilotEventEnrollmentTypeofflineDomainJoined,
		"offlinedomainjoinedwithofflineautopilotprofile":      DeviceManagementAutopilotEventEnrollmentTypeofflineDomainJoinedWithOfflineAutopilotProfile,
		"offlinedomainjoinedwithwhiteglove":                   DeviceManagementAutopilotEventEnrollmentTypeofflineDomainJoinedWithWhiteGlove,
		"unknown":                                             DeviceManagementAutopilotEventEnrollmentTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotEventEnrollmentType(input)
	return &out, nil
}
