package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CorsConfigurationOperationPredicate struct {
	MaxAgeInSeconds *int64
	ODataType       *string
	Resource        *string
}

func (p CorsConfigurationOperationPredicate) Matches(input CorsConfiguration) bool {

	if p.MaxAgeInSeconds != nil && (input.MaxAgeInSeconds == nil || *p.MaxAgeInSeconds != *input.MaxAgeInSeconds) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Resource != nil && (input.Resource == nil || *p.Resource != *input.Resource) {
		return false
	}

	return true
}
