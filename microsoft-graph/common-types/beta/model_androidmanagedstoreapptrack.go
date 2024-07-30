package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppTrack struct {
	ODataType  *string `json:"@odata.type,omitempty"`
	TrackAlias *string `json:"trackAlias,omitempty"`
	TrackId    *string `json:"trackId,omitempty"`
}
