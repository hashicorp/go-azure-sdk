package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsAccessKeyDetails struct {
	CreatedDateTime  *string `json:"createdDateTime,omitempty"`
	LastUsedDateTime *string `json:"lastUsedDateTime,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
}
