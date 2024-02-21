package extendedueinformation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UeInfo5GProperties struct {
	ConnectionInfo *UeConnectionInfo5G `json:"connectionInfo,omitempty"`
	FivegGuti      Guti5G              `json:"fivegGuti"`
	Pei            *string             `json:"pei,omitempty"`
	SessionInfo    *[]UeSessionInfo5G  `json:"sessionInfo,omitempty"`
	Supi           string              `json:"supi"`
}
