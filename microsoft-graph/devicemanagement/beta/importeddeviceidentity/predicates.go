package importeddeviceidentity

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type ImportedDeviceIdentityOperationPredicate struct {
}

func (p ImportedDeviceIdentityOperationPredicate) Matches(input beta.ImportedDeviceIdentity) bool {

	return true
}

type ImportedDeviceIdentityResultOperationPredicate struct {
}

func (p ImportedDeviceIdentityResultOperationPredicate) Matches(input beta.ImportedDeviceIdentityResult) bool {

	return true
}
