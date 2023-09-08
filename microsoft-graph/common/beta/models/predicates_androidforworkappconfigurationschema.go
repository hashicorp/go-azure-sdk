package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkAppConfigurationSchemaOperationPredicate struct {
	ExampleJson *string
	Id          *string
	ODataType   *string
}

func (p AndroidForWorkAppConfigurationSchemaOperationPredicate) Matches(input AndroidForWorkAppConfigurationSchema) bool {

	if p.ExampleJson != nil && (input.ExampleJson == nil || *p.ExampleJson != *input.ExampleJson) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
