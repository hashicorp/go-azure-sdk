package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkPeripheralsHealth struct {
	CommunicationSpeakerHealth *TeamworkPeripheralHealth   `json:"communicationSpeakerHealth,omitempty"`
	ContentCameraHealth        *TeamworkPeripheralHealth   `json:"contentCameraHealth,omitempty"`
	DisplayHealthCollection    *[]TeamworkPeripheralHealth `json:"displayHealthCollection,omitempty"`
	MicrophoneHealth           *TeamworkPeripheralHealth   `json:"microphoneHealth,omitempty"`
	ODataType                  *string                     `json:"@odata.type,omitempty"`
	RoomCameraHealth           *TeamworkPeripheralHealth   `json:"roomCameraHealth,omitempty"`
	SpeakerHealth              *TeamworkPeripheralHealth   `json:"speakerHealth,omitempty"`
}
