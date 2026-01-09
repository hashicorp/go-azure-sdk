package deviceconfiguration

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type DeviceConfigurationOperationPredicate struct {
}

func (p DeviceConfigurationOperationPredicate) Matches(input stable.DeviceConfiguration) bool {

	return true
}

type DeviceConfigurationAssignmentOperationPredicate struct {
}

func (p DeviceConfigurationAssignmentOperationPredicate) Matches(input stable.DeviceConfigurationAssignment) bool {

	return true
}
