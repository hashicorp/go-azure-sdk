package extendedueinformation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GlobalRanNodeId struct {
	ENbId   *string `json:"eNbId,omitempty"`
	GNbId   *GNbId  `json:"gNbId,omitempty"`
	N3IwfId *string `json:"n3IwfId,omitempty"`
	NgeNbId *string `json:"ngeNbId,omitempty"`
	Nid     *string `json:"nid,omitempty"`
	PlmnId  PlmnId  `json:"plmnId"`
	TngfId  *string `json:"tngfId,omitempty"`
	WagfId  *string `json:"wagfId,omitempty"`
}
