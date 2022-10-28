package elasticsanskus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkuLocationInfo struct {
	Location *string `json:"location,omitempty"`
	Zones    *Zones  `json:"zones,omitempty"`
}
