package extendedueinformation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UeQOSFlow struct {
	Fiveqi int64 `json:"fiveqi"`
	Gbr    *Ambr `json:"gbr,omitempty"`
	Mbr    *Ambr `json:"mbr,omitempty"`
	Qfi    int64 `json:"qfi"`
}
