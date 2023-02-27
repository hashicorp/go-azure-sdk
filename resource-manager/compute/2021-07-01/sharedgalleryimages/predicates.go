// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sharedgalleryimages

type SharedGalleryImageOperationPredicate struct {
	Location *string
	Name     *string
}

func (p SharedGalleryImageOperationPredicate) Matches(input SharedGalleryImage) bool {

	if p.Location != nil && (input.Location == nil && *p.Location != *input.Location) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	return true
}
