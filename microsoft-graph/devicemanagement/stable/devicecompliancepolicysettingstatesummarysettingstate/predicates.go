package devicecompliancepolicysettingstatesummarysettingstate

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type DeviceComplianceSettingStateOperationPredicate struct {
}

func (p DeviceComplianceSettingStateOperationPredicate) Matches(input stable.DeviceComplianceSettingState) bool {

	return true
}
