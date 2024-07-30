package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationConditions struct {
	Applications *AuthenticationConditionsApplications `json:"applications,omitempty"`
	ODataType    *string                               `json:"@odata.type,omitempty"`
}
