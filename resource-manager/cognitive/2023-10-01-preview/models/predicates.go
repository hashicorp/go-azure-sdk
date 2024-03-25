package models

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModelOperationPredicate struct {
	Kind    *string
	SkuName *string
}

func (p ModelOperationPredicate) Matches(input Model) bool {

	if p.Kind != nil && (input.Kind == nil || *p.Kind != *input.Kind) {
		return false
	}

	if p.SkuName != nil && (input.SkuName == nil || *p.SkuName != *input.SkuName) {
		return false
	}

	return true
}
