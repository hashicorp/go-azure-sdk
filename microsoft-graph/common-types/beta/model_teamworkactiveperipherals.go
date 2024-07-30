package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkActivePeripherals struct {
	CommunicationSpeaker *TeamworkPeripheral `json:"communicationSpeaker,omitempty"`
	ContentCamera        *TeamworkPeripheral `json:"contentCamera,omitempty"`
	Microphone           *TeamworkPeripheral `json:"microphone,omitempty"`
	ODataType            *string             `json:"@odata.type,omitempty"`
	RoomCamera           *TeamworkPeripheral `json:"roomCamera,omitempty"`
	Speaker              *TeamworkPeripheral `json:"speaker,omitempty"`
}
