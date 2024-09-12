package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactMatchClassificationRequest struct {
	ContentClassifications *[]ContentClassification `json:"contentClassifications,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	SensitiveTypeIds *[]string             `json:"sensitiveTypeIds,omitempty"`
	Text             nullable.Type[string] `json:"text,omitempty"`
	TimeoutInMs      nullable.Type[int64]  `json:"timeoutInMs,omitempty"`
}
