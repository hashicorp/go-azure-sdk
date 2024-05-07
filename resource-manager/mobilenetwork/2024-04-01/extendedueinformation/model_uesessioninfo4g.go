package extendedueinformation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UeSessionInfo4G struct {
	Apn         string      `json:"apn"`
	Ebi         int64       `json:"ebi"`
	PdnType     PdnType     `json:"pdnType"`
	UeIPAddress UeIPAddress `json:"ueIpAddress"`
}
