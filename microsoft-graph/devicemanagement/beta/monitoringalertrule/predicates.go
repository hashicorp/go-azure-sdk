package monitoringalertrule

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceManagementAlertRuleOperationPredicate struct {
}

func (p DeviceManagementAlertRuleOperationPredicate) Matches(input beta.DeviceManagementAlertRule) bool {

	return true
}
