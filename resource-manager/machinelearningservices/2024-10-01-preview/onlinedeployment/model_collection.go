package onlinedeployment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Collection struct {
	ClientId           *string             `json:"clientId,omitempty"`
	DataCollectionMode *DataCollectionMode `json:"dataCollectionMode,omitempty"`
	DataId             *string             `json:"dataId,omitempty"`
	SamplingRate       *float64            `json:"samplingRate,omitempty"`
}
