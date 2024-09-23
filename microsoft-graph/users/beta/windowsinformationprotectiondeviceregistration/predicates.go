package windowsinformationprotectiondeviceregistration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type WindowsInformationProtectionDeviceRegistrationOperationPredicate struct {
}

func (p WindowsInformationProtectionDeviceRegistrationOperationPredicate) Matches(input beta.WindowsInformationProtectionDeviceRegistration) bool {

	return true
}
