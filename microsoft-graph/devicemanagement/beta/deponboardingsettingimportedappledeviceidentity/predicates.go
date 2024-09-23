package deponboardingsettingimportedappledeviceidentity

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type ImportedAppleDeviceIdentityOperationPredicate struct {
}

func (p ImportedAppleDeviceIdentityOperationPredicate) Matches(input beta.ImportedAppleDeviceIdentity) bool {

	return true
}

type ImportedAppleDeviceIdentityResultOperationPredicate struct {
}

func (p ImportedAppleDeviceIdentityResultOperationPredicate) Matches(input beta.ImportedAppleDeviceIdentityResult) bool {

	return true
}
