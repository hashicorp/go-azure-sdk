package deleteditem

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type AdministrativeUnitOperationPredicate struct {
}

func (p AdministrativeUnitOperationPredicate) Matches(input stable.AdministrativeUnit) bool {

	return true
}

type ApplicationOperationPredicate struct {
}

func (p ApplicationOperationPredicate) Matches(input stable.Application) bool {

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

type ExtensionPropertyOperationPredicate struct {
}

func (p ExtensionPropertyOperationPredicate) Matches(input stable.ExtensionProperty) bool {

	return true
}

type GroupOperationPredicate struct {
}

func (p GroupOperationPredicate) Matches(input stable.Group) bool {

	return true
}

type ServicePrincipalOperationPredicate struct {
}

func (p ServicePrincipalOperationPredicate) Matches(input stable.ServicePrincipal) bool {

	return true
}

type UserOperationPredicate struct {
}

func (p UserOperationPredicate) Matches(input stable.User) bool {

	return true
}
