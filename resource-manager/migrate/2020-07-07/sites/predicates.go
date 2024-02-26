package sites

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteHealthSummaryCollectionOperationPredicate struct {
	NextLink *string
}

func (p SiteHealthSummaryCollectionOperationPredicate) Matches(input SiteHealthSummaryCollection) bool {

	if p.NextLink != nil && (input.NextLink == nil || *p.NextLink != *input.NextLink) {
		return false
	}

	return true
}
