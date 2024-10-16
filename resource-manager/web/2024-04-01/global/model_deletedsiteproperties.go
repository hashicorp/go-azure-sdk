package global

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedSiteProperties struct {
	DeletedSiteId    *int64  `json:"deletedSiteId,omitempty"`
	DeletedSiteName  *string `json:"deletedSiteName,omitempty"`
	DeletedTimestamp *string `json:"deletedTimestamp,omitempty"`
	GeoRegionName    *string `json:"geoRegionName,omitempty"`
	Kind             *string `json:"kind,omitempty"`
	ResourceGroup    *string `json:"resourceGroup,omitempty"`
	Slot             *string `json:"slot,omitempty"`
	Subscription     *string `json:"subscription,omitempty"`
}
