package drivespecial

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type DriveItemOperationPredicate struct {
}

func (p DriveItemOperationPredicate) Matches(input stable.DriveItem) bool {

	return true
}
