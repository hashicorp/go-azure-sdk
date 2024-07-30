package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationConditionsApplications struct {
	IncludeAllApplications *bool                                 `json:"includeAllApplications,omitempty"`
	IncludeApplications    *[]AuthenticationConditionApplication `json:"includeApplications,omitempty"`
	ODataType              *string                               `json:"@odata.type,omitempty"`
}
