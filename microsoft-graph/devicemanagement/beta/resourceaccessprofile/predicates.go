package resourceaccessprofile

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceManagementResourceAccessProfileAssignmentOperationPredicate struct {
}

func (p DeviceManagementResourceAccessProfileAssignmentOperationPredicate) Matches(input beta.DeviceManagementResourceAccessProfileAssignment) bool {

	return true
}

type DeviceManagementResourceAccessProfileBaseOperationPredicate struct {
}

func (p DeviceManagementResourceAccessProfileBaseOperationPredicate) Matches(input beta.DeviceManagementResourceAccessProfileBase) bool {

	return true
}
