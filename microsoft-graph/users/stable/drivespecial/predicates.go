package drivespecial

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type DriveItemOperationPredicate struct {
}

func (p DriveItemOperationPredicate) Matches(input stable.DriveItem) bool {

	return true
}
