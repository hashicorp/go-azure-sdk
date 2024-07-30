package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotEventEnrollmentType string

const (
	DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedUsingDeviceAuthWithAutopilotProfile    DeviceManagementAutopilotEventEnrollmentType = "azureADJoinedUsingDeviceAuthWithAutopilotProfile"
	DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedUsingDeviceAuthWithoutAutopilotProfile DeviceManagementAutopilotEventEnrollmentType = "azureADJoinedUsingDeviceAuthWithoutAutopilotProfile"
	DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedWithAutopilotProfile                   DeviceManagementAutopilotEventEnrollmentType = "azureADJoinedWithAutopilotProfile"
	DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedWithOfflineAutopilotProfile            DeviceManagementAutopilotEventEnrollmentType = "azureADJoinedWithOfflineAutopilotProfile"
	DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedWithWhiteGlove                         DeviceManagementAutopilotEventEnrollmentType = "azureADJoinedWithWhiteGlove"
	DeviceManagementAutopilotEventEnrollmentType_OfflineDomainJoined                                 DeviceManagementAutopilotEventEnrollmentType = "offlineDomainJoined"
	DeviceManagementAutopilotEventEnrollmentType_OfflineDomainJoinedWithOfflineAutopilotProfile      DeviceManagementAutopilotEventEnrollmentType = "offlineDomainJoinedWithOfflineAutopilotProfile"
	DeviceManagementAutopilotEventEnrollmentType_OfflineDomainJoinedWithWhiteGlove                   DeviceManagementAutopilotEventEnrollmentType = "offlineDomainJoinedWithWhiteGlove"
	DeviceManagementAutopilotEventEnrollmentType_Unknown                                             DeviceManagementAutopilotEventEnrollmentType = "unknown"
)

func PossibleValuesForDeviceManagementAutopilotEventEnrollmentType() []string {
	return []string{
		string(DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedUsingDeviceAuthWithAutopilotProfile),
		string(DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedUsingDeviceAuthWithoutAutopilotProfile),
		string(DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedWithAutopilotProfile),
		string(DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedWithOfflineAutopilotProfile),
		string(DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedWithWhiteGlove),
		string(DeviceManagementAutopilotEventEnrollmentType_OfflineDomainJoined),
		string(DeviceManagementAutopilotEventEnrollmentType_OfflineDomainJoinedWithOfflineAutopilotProfile),
		string(DeviceManagementAutopilotEventEnrollmentType_OfflineDomainJoinedWithWhiteGlove),
		string(DeviceManagementAutopilotEventEnrollmentType_Unknown),
	}
}

func (s *DeviceManagementAutopilotEventEnrollmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAutopilotEventEnrollmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAutopilotEventEnrollmentType(input string) (*DeviceManagementAutopilotEventEnrollmentType, error) {
	vals := map[string]DeviceManagementAutopilotEventEnrollmentType{
		"azureadjoinedusingdeviceauthwithautopilotprofile":    DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedUsingDeviceAuthWithAutopilotProfile,
		"azureadjoinedusingdeviceauthwithoutautopilotprofile": DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedUsingDeviceAuthWithoutAutopilotProfile,
		"azureadjoinedwithautopilotprofile":                   DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedWithAutopilotProfile,
		"azureadjoinedwithofflineautopilotprofile":            DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedWithOfflineAutopilotProfile,
		"azureadjoinedwithwhiteglove":                         DeviceManagementAutopilotEventEnrollmentType_AzureADJoinedWithWhiteGlove,
		"offlinedomainjoined":                                 DeviceManagementAutopilotEventEnrollmentType_OfflineDomainJoined,
		"offlinedomainjoinedwithofflineautopilotprofile":      DeviceManagementAutopilotEventEnrollmentType_OfflineDomainJoinedWithOfflineAutopilotProfile,
		"offlinedomainjoinedwithwhiteglove":                   DeviceManagementAutopilotEventEnrollmentType_OfflineDomainJoinedWithWhiteGlove,
		"unknown":                                             DeviceManagementAutopilotEventEnrollmentType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotEventEnrollmentType(input)
	return &out, nil
}
