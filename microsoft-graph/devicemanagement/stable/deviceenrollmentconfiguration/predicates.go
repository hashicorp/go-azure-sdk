package deviceenrollmentconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type DeviceEnrollmentConfigurationOperationPredicate struct {
}

func (p DeviceEnrollmentConfigurationOperationPredicate) Matches(input stable.DeviceEnrollmentConfiguration) bool {

	return true
}
