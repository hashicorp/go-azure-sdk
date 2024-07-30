package manageddevicedevicehealthscriptstate

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceHealthScriptPolicyStateOperationPredicate struct {
}

func (p DeviceHealthScriptPolicyStateOperationPredicate) Matches(input beta.DeviceHealthScriptPolicyState) bool {

	return true
}
