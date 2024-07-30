package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkMicrophoneConfiguration struct {
	DefaultMicrophone    *TeamworkPeripheral   `json:"defaultMicrophone,omitempty"`
	IsMicrophoneOptional *bool                 `json:"isMicrophoneOptional,omitempty"`
	Microphones          *[]TeamworkPeripheral `json:"microphones,omitempty"`
	ODataType            *string               `json:"@odata.type,omitempty"`
}
