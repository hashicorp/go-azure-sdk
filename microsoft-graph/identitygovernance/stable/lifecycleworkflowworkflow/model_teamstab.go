package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsTab struct {
	Configuration *TeamsTabConfiguration `json:"configuration,omitempty"`
	DisplayName   *string                `json:"displayName,omitempty"`
	Id            *string                `json:"id,omitempty"`
	ODataType     *string                `json:"@odata.type,omitempty"`
	TeamsApp      *TeamsApp              `json:"teamsApp,omitempty"`
	WebUrl        *string                `json:"webUrl,omitempty"`
}
