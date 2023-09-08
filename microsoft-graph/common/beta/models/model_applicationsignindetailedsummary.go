package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationSignInDetailedSummary struct {
	AggregatedEventDateTime *string       `json:"aggregatedEventDateTime,omitempty"`
	AppDisplayName          *string       `json:"appDisplayName,omitempty"`
	AppId                   *string       `json:"appId,omitempty"`
	Id                      *string       `json:"id,omitempty"`
	ODataType               *string       `json:"@odata.type,omitempty"`
	SignInCount             *int64        `json:"signInCount,omitempty"`
	Status                  *SignInStatus `json:"status,omitempty"`
}
