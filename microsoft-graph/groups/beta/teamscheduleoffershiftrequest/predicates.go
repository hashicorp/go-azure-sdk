package teamscheduleoffershiftrequest

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type OfferShiftRequestOperationPredicate struct {
}

func (p OfferShiftRequestOperationPredicate) Matches(input beta.OfferShiftRequest) bool {

	return true
}
