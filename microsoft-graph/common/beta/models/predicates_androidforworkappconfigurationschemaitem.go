package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkAppConfigurationSchemaItemOperationPredicate struct {
	DefaultBoolValue   *bool
	DefaultIntValue    *int64
	DefaultStringValue *string
	Description        *string
	DisplayName        *string
	ODataType          *string
	SchemaItemKey      *string
}

func (p AndroidForWorkAppConfigurationSchemaItemOperationPredicate) Matches(input AndroidForWorkAppConfigurationSchemaItem) bool {

	if p.DefaultBoolValue != nil && (input.DefaultBoolValue == nil || *p.DefaultBoolValue != *input.DefaultBoolValue) {
		return false
	}

	if p.DefaultIntValue != nil && (input.DefaultIntValue == nil || *p.DefaultIntValue != *input.DefaultIntValue) {
		return false
	}

	if p.DefaultStringValue != nil && (input.DefaultStringValue == nil || *p.DefaultStringValue != *input.DefaultStringValue) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SchemaItemKey != nil && (input.SchemaItemKey == nil || *p.SchemaItemKey != *input.SchemaItemKey) {
		return false
	}

	return true
}
