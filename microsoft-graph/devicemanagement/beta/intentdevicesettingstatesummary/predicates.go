package intentdevicesettingstatesummary

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceManagementIntentDeviceSettingStateSummaryOperationPredicate struct {
}

func (p DeviceManagementIntentDeviceSettingStateSummaryOperationPredicate) Matches(input beta.DeviceManagementIntentDeviceSettingStateSummary) bool {

	return true
}
