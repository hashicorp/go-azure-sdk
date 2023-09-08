package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioProperties struct {
	ExternalBucketId      *string `json:"externalBucketId,omitempty"`
	ExternalContextId     *string `json:"externalContextId,omitempty"`
	ExternalObjectId      *string `json:"externalObjectId,omitempty"`
	ExternalObjectVersion *string `json:"externalObjectVersion,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	WebUrl                *string `json:"webUrl,omitempty"`
}
