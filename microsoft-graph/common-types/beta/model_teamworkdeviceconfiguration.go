package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceConfiguration struct {
	CameraConfiguration      *TeamworkCameraConfiguration      `json:"cameraConfiguration,omitempty"`
	CreatedBy                *IdentitySet                      `json:"createdBy,omitempty"`
	CreatedDateTime          *string                           `json:"createdDateTime,omitempty"`
	DisplayConfiguration     *TeamworkDisplayConfiguration     `json:"displayConfiguration,omitempty"`
	HardwareConfiguration    *TeamworkHardwareConfiguration    `json:"hardwareConfiguration,omitempty"`
	Id                       *string                           `json:"id,omitempty"`
	LastModifiedBy           *IdentitySet                      `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime     *string                           `json:"lastModifiedDateTime,omitempty"`
	MicrophoneConfiguration  *TeamworkMicrophoneConfiguration  `json:"microphoneConfiguration,omitempty"`
	ODataType                *string                           `json:"@odata.type,omitempty"`
	SoftwareVersions         *TeamworkDeviceSoftwareVersions   `json:"softwareVersions,omitempty"`
	SpeakerConfiguration     *TeamworkSpeakerConfiguration     `json:"speakerConfiguration,omitempty"`
	SystemConfiguration      *TeamworkSystemConfiguration      `json:"systemConfiguration,omitempty"`
	TeamsClientConfiguration *TeamworkTeamsClientConfiguration `json:"teamsClientConfiguration,omitempty"`
}
