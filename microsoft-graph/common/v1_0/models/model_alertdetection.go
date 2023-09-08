package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertDetection struct {
	DetectionType *string `json:"detectionType,omitempty"`
	Method        *string `json:"method,omitempty"`
	Name          *string `json:"name,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
}
