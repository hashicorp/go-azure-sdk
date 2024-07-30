package virtualendpointsupportedregion

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type CloudPCSupportedRegionOperationPredicate struct {
}

func (p CloudPCSupportedRegionOperationPredicate) Matches(input beta.CloudPCSupportedRegion) bool {

	return true
}
