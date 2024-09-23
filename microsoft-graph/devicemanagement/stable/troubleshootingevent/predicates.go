package troubleshootingevent

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type DeviceManagementTroubleshootingEventOperationPredicate struct {
}

func (p DeviceManagementTroubleshootingEventOperationPredicate) Matches(input stable.DeviceManagementTroubleshootingEvent) bool {

	return true
}
