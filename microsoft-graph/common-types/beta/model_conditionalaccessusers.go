package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessUsers struct {
	ExcludeGroups                *[]string                               `json:"excludeGroups,omitempty"`
	ExcludeGuestsOrExternalUsers *ConditionalAccessGuestsOrExternalUsers `json:"excludeGuestsOrExternalUsers,omitempty"`
	ExcludeRoles                 *[]string                               `json:"excludeRoles,omitempty"`
	ExcludeUsers                 *[]string                               `json:"excludeUsers,omitempty"`
	IncludeGroups                *[]string                               `json:"includeGroups,omitempty"`
	IncludeGuestsOrExternalUsers *ConditionalAccessGuestsOrExternalUsers `json:"includeGuestsOrExternalUsers,omitempty"`
	IncludeRoles                 *[]string                               `json:"includeRoles,omitempty"`
	IncludeUsers                 *[]string                               `json:"includeUsers,omitempty"`
	ODataType                    *string                                 `json:"@odata.type,omitempty"`
}
