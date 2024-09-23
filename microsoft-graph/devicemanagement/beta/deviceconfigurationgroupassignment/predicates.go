package deviceconfigurationgroupassignment

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceConfigurationGroupAssignmentOperationPredicate struct {
}

func (p DeviceConfigurationGroupAssignmentOperationPredicate) Matches(input beta.DeviceConfigurationGroupAssignment) bool {

	return true
}
