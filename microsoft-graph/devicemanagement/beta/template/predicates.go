package template

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceManagementIntentOperationPredicate struct {
}

func (p DeviceManagementIntentOperationPredicate) Matches(input beta.DeviceManagementIntent) bool {

	return true
}

type DeviceManagementTemplateOperationPredicate struct {
}

func (p DeviceManagementTemplateOperationPredicate) Matches(input beta.DeviceManagementTemplate) bool {

	return true
}
