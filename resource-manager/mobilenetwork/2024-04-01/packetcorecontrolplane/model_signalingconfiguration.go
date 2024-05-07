package packetcorecontrolplane

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignalingConfiguration struct {
	NasEncryption *[]NasEncryptionType     `json:"nasEncryption,omitempty"`
	NasReroute    *NASRerouteConfiguration `json:"nasReroute,omitempty"`
}
