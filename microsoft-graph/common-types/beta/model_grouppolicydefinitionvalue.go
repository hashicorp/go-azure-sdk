package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionValue struct {
	ConfigurationType    *GroupPolicyDefinitionValueConfigurationType `json:"configurationType,omitempty"`
	CreatedDateTime      *string                                      `json:"createdDateTime,omitempty"`
	Definition           *GroupPolicyDefinition                       `json:"definition,omitempty"`
	Enabled              *bool                                        `json:"enabled,omitempty"`
	Id                   *string                                      `json:"id,omitempty"`
	LastModifiedDateTime *string                                      `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                      `json:"@odata.type,omitempty"`
	PresentationValues   *[]GroupPolicyPresentationValue              `json:"presentationValues,omitempty"`
}
