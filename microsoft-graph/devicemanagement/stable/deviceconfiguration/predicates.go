package deviceconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type DeviceConfigurationOperationPredicate struct {
}

func (p DeviceConfigurationOperationPredicate) Matches(input stable.DeviceConfiguration) bool {

	return true
}

type DeviceConfigurationAssignmentOperationPredicate struct {
	Id        *string
	ODataType *string
}

func (p DeviceConfigurationAssignmentOperationPredicate) Matches(input stable.DeviceConfigurationAssignment) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
