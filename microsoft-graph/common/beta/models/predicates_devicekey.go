package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceKeyOperationPredicate struct {
	DeviceId    *string
	KeyMaterial *string
	KeyType     *string
	ODataType   *string
}

func (p DeviceKeyOperationPredicate) Matches(input DeviceKey) bool {

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.KeyMaterial != nil && (input.KeyMaterial == nil || *p.KeyMaterial != *input.KeyMaterial) {
		return false
	}

	if p.KeyType != nil && (input.KeyType == nil || *p.KeyType != *input.KeyType) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
