package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkLoginStatus struct {
	ExchangeConnection *TeamworkConnection `json:"exchangeConnection,omitempty"`
	ODataType          *string             `json:"@odata.type,omitempty"`
	SkypeConnection    *TeamworkConnection `json:"skypeConnection,omitempty"`
	TeamsConnection    *TeamworkConnection `json:"teamsConnection,omitempty"`
}
