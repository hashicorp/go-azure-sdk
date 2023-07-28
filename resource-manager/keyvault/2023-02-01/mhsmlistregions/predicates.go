package mhsmlistregions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MHSMGeoReplicatedRegionOperationPredicate struct {
	IsPrimary *bool
	Name      *string
}

func (p MHSMGeoReplicatedRegionOperationPredicate) Matches(input MHSMGeoReplicatedRegion) bool {

	if p.IsPrimary != nil && (input.IsPrimary == nil || *p.IsPrimary != *input.IsPrimary) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	return true
}
