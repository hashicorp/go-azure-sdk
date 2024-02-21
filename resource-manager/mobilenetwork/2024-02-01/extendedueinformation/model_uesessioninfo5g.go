package extendedueinformation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UeSessionInfo5G struct {
	Ambr         Ambr        `json:"ambr"`
	Dnn          string      `json:"dnn"`
	PdnType      PdnType     `json:"pdnType"`
	PduSessionId int64       `json:"pduSessionId"`
	QosFlow      []UeQOSFlow `json:"qosFlow"`
	Snssai       Snssai      `json:"snssai"`
	UeIPAddress  UeIPAddress `json:"ueIpAddress"`
}
