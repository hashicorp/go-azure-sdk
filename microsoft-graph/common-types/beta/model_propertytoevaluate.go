package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PropertyToEvaluate struct {
	ODataType     *string `json:"@odata.type,omitempty"`
	PropertyName  *string `json:"propertyName,omitempty"`
	PropertyValue *string `json:"propertyValue,omitempty"`
}
