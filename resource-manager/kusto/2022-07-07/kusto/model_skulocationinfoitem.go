package kusto

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkuLocationInfoItem struct {
	Location string    `json:"location"`
	Zones    *[]string `json:"zones,omitempty"`
}
