package hardwareconfiguration

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type HardwareConfigurationOperationPredicate struct {
}

func (p HardwareConfigurationOperationPredicate) Matches(input beta.HardwareConfiguration) bool {

	return true
}

type HardwareConfigurationAssignmentOperationPredicate struct {
}

func (p HardwareConfigurationAssignmentOperationPredicate) Matches(input beta.HardwareConfigurationAssignment) bool {

	return true
}
