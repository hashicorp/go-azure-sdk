package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationIdentityMatchingOptions struct {
	AppliesTo          *EducationIdentityMatchingOptionsAppliesTo `json:"appliesTo,omitempty"`
	ODataType          *string                                    `json:"@odata.type,omitempty"`
	SourcePropertyName *string                                    `json:"sourcePropertyName,omitempty"`
	TargetDomain       *string                                    `json:"targetDomain,omitempty"`
	TargetPropertyName *string                                    `json:"targetPropertyName,omitempty"`
}
