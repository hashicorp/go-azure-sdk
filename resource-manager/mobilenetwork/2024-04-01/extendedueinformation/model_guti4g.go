package extendedueinformation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Guti4G struct {
	MTmsi int64  `json:"mTmsi"`
	MmeId MmeId  `json:"mmeId"`
	Plmn  PlmnId `json:"plmn"`
}
