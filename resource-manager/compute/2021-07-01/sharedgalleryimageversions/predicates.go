// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sharedgalleryimageversions

type SharedGalleryImageVersionOperationPredicate struct {
	Location *string
	Name     *string
}

func (p SharedGalleryImageVersionOperationPredicate) Matches(input SharedGalleryImageVersion) bool {

	if p.Location != nil && (input.Location == nil && *p.Location != *input.Location) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	return true
}
