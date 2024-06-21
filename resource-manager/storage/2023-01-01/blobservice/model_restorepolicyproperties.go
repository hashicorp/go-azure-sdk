package blobservice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePolicyProperties struct {
	Days            *int64  `json:"days,omitempty"`
	Enabled         bool    `json:"enabled"`
	LastEnabledTime *string `json:"lastEnabledTime,omitempty"`
	MinRestoreTime  *string `json:"minRestoreTime,omitempty"`
}
