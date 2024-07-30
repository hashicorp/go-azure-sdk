package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterSupportedProperty struct {
	DataType                *string                                              `json:"dataType,omitempty"`
	IsCollection            *bool                                                `json:"isCollection,omitempty"`
	Name                    *string                                              `json:"name,omitempty"`
	ODataType               *string                                              `json:"@odata.type,omitempty"`
	PropertyRegexConstraint *string                                              `json:"propertyRegexConstraint,omitempty"`
	SupportedOperators      *AssignmentFilterSupportedPropertySupportedOperators `json:"supportedOperators,omitempty"`
	SupportedValues         *[]string                                            `json:"supportedValues,omitempty"`
}
