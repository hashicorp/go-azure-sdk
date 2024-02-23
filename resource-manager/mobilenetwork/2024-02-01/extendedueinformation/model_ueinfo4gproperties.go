package extendedueinformation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UeInfo4GProperties struct {
	ConnectionInfo *UeConnectionInfo4G `json:"connectionInfo,omitempty"`
	Guti           Guti4G              `json:"guti"`
	Imei           *string             `json:"imei,omitempty"`
	Imeisv         *string             `json:"imeisv,omitempty"`
	Imsi           string              `json:"imsi"`
	SessionInfo    *[]UeSessionInfo4G  `json:"sessionInfo,omitempty"`
}
