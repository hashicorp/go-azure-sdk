package devicecompliancepolicyuserstatus

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type DeviceComplianceUserStatusOperationPredicate struct {
}

func (p DeviceComplianceUserStatusOperationPredicate) Matches(input stable.DeviceComplianceUserStatus) bool {

	return true
}
