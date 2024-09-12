package devicecompliancepolicysettingstatesummary

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type DeviceCompliancePolicySettingStateSummaryOperationPredicate struct {
}

func (p DeviceCompliancePolicySettingStateSummaryOperationPredicate) Matches(input stable.DeviceCompliancePolicySettingStateSummary) bool {

	return true
}
