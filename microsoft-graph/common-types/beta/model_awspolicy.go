package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsPolicy struct {
	AwsPolicyType *AwsPolicyAwsPolicyType `json:"awsPolicyType,omitempty"`
	DisplayName   *string                 `json:"displayName,omitempty"`
	ExternalId    *string                 `json:"externalId,omitempty"`
	Id            *string                 `json:"id,omitempty"`
	ODataType     *string                 `json:"@odata.type,omitempty"`
}
