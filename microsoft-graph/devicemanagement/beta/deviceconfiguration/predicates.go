package deviceconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceConfigurationOperationPredicate struct {
}

func (p DeviceConfigurationOperationPredicate) Matches(input beta.DeviceConfiguration) bool {

	return true
}

type DeviceConfigurationAssignmentOperationPredicate struct {
}

func (p DeviceConfigurationAssignmentOperationPredicate) Matches(input beta.DeviceConfigurationAssignment) bool {

	return true
}

type DeviceConfigurationTargetedUserAndDeviceOperationPredicate struct {
}

func (p DeviceConfigurationTargetedUserAndDeviceOperationPredicate) Matches(input beta.DeviceConfigurationTargetedUserAndDevice) bool {

	return true
}

type HasPayloadLinkResultItemOperationPredicate struct {
}

func (p HasPayloadLinkResultItemOperationPredicate) Matches(input beta.HasPayloadLinkResultItem) bool {

	return true
}
