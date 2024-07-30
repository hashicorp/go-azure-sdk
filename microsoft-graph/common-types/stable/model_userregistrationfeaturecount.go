package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationFeatureCount struct {
	Feature   *UserRegistrationFeatureCountFeature `json:"feature,omitempty"`
	ODataType *string                              `json:"@odata.type,omitempty"`
	UserCount *int64                               `json:"userCount,omitempty"`
}
