package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkSpeakerConfiguration struct {
	DefaultCommunicationSpeaker    *TeamworkPeripheral   `json:"defaultCommunicationSpeaker,omitempty"`
	DefaultSpeaker                 *TeamworkPeripheral   `json:"defaultSpeaker,omitempty"`
	IsCommunicationSpeakerOptional *bool                 `json:"isCommunicationSpeakerOptional,omitempty"`
	IsSpeakerOptional              *bool                 `json:"isSpeakerOptional,omitempty"`
	ODataType                      *string               `json:"@odata.type,omitempty"`
	Speakers                       *[]TeamworkPeripheral `json:"speakers,omitempty"`
}
