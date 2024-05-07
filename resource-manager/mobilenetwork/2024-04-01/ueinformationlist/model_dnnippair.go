package ueinformationlist

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DnnIPPair struct {
	Dnn         *string      `json:"dnn,omitempty"`
	UeIPAddress *UeIPAddress `json:"ueIpAddress,omitempty"`
}
