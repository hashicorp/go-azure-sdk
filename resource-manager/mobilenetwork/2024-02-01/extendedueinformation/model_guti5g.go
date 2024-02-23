package extendedueinformation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Guti5G struct {
	AmfId     AmfId  `json:"amfId"`
	FivegTmsi int64  `json:"fivegTmsi"`
	Plmn      PlmnId `json:"plmn"`
}
