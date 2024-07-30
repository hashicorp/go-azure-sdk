package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExactMatchClassificationResult struct {
	Classification *[]ExactMatchDetectedSensitiveContent `json:"classification,omitempty"`
	Errors         *[]ClassificationError                `json:"errors,omitempty"`
	ODataType      *string                               `json:"@odata.type,omitempty"`
}
