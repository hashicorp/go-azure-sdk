package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeleconferenceDeviceQuality struct {
	CallChainId                       *string                             `json:"callChainId,omitempty"`
	CloudServiceDeploymentEnvironment *string                             `json:"cloudServiceDeploymentEnvironment,omitempty"`
	CloudServiceDeploymentId          *string                             `json:"cloudServiceDeploymentId,omitempty"`
	CloudServiceInstanceName          *string                             `json:"cloudServiceInstanceName,omitempty"`
	CloudServiceName                  *string                             `json:"cloudServiceName,omitempty"`
	DeviceDescription                 *string                             `json:"deviceDescription,omitempty"`
	DeviceName                        *string                             `json:"deviceName,omitempty"`
	MediaLegId                        *string                             `json:"mediaLegId,omitempty"`
	MediaQualityList                  *[]TeleconferenceDeviceMediaQuality `json:"mediaQualityList,omitempty"`
	ODataType                         *string                             `json:"@odata.type,omitempty"`
	ParticipantId                     *string                             `json:"participantId,omitempty"`
}
