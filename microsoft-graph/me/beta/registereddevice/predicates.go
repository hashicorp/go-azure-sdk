package registereddevice

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type AppRoleAssignmentOperationPredicate struct {
}

func (p AppRoleAssignmentOperationPredicate) Matches(input beta.AppRoleAssignment) bool {

	return true
}

type DeviceOperationPredicate struct {
}

func (p DeviceOperationPredicate) Matches(input beta.Device) bool {

	return true
}

type DirectoryObjectOperationPredicate struct {
}

func (p DirectoryObjectOperationPredicate) Matches(input beta.DirectoryObject) bool {

	return true
}

type EndpointOperationPredicate struct {
}

func (p EndpointOperationPredicate) Matches(input beta.Endpoint) bool {

	return true
}
