package registereddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type AppRoleAssignmentOperationPredicate struct {
}

func (p AppRoleAssignmentOperationPredicate) Matches(input stable.AppRoleAssignment) bool {

	return true
}

type DeviceOperationPredicate struct {
}

func (p DeviceOperationPredicate) Matches(input stable.Device) bool {

	return true
}

type DirectoryObjectOperationPredicate struct {
}

func (p DirectoryObjectOperationPredicate) Matches(input stable.DirectoryObject) bool {

	return true
}

type EndpointOperationPredicate struct {
}

func (p EndpointOperationPredicate) Matches(input stable.Endpoint) bool {

	return true
}
