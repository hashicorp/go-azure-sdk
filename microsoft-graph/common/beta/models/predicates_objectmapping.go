package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectMappingOperationPredicate struct {
	Enabled          *bool
	Name             *string
	ODataType        *string
	SourceObjectName *string
	TargetObjectName *string
}

func (p ObjectMappingOperationPredicate) Matches(input ObjectMapping) bool {

	if p.Enabled != nil && (input.Enabled == nil || *p.Enabled != *input.Enabled) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SourceObjectName != nil && (input.SourceObjectName == nil || *p.SourceObjectName != *input.SourceObjectName) {
		return false
	}

	if p.TargetObjectName != nil && (input.TargetObjectName == nil || *p.TargetObjectName != *input.TargetObjectName) {
		return false
	}

	return true
}
