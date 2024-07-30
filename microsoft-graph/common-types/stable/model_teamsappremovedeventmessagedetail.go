package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppRemovedEventMessageDetail struct {
	Initiator           *IdentitySet `json:"initiator,omitempty"`
	ODataType           *string      `json:"@odata.type,omitempty"`
	TeamsAppDisplayName *string      `json:"teamsAppDisplayName,omitempty"`
	TeamsAppId          *string      `json:"teamsAppId,omitempty"`
}
