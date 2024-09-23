package androiddeviceownerenrollmentprofile

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type AndroidDeviceOwnerEnrollmentProfileOperationPredicate struct {
}

func (p AndroidDeviceOwnerEnrollmentProfileOperationPredicate) Matches(input beta.AndroidDeviceOwnerEnrollmentProfile) bool {

	return true
}
