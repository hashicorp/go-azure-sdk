package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationOAuth2ClientCredentialsConnectionSettings struct {
	ClientId     *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	Scope        *string `json:"scope,omitempty"`
	TokenUrl     *string `json:"tokenUrl,omitempty"`
}
