package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsMinimumOperatingSystemOperationPredicate struct {
	ODataType *string
	V100      *bool
	V101607   *bool
	V101703   *bool
	V101709   *bool
	V101803   *bool
	V101809   *bool
	V101903   *bool
	V101909   *bool
	V102004   *bool
	V1021H1   *bool
	V102H20   *bool
	V80       *bool
	V81       *bool
}

func (p WindowsMinimumOperatingSystemOperationPredicate) Matches(input WindowsMinimumOperatingSystem) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.V100 != nil && (input.V100 == nil || *p.V100 != *input.V100) {
		return false
	}

	if p.V101607 != nil && (input.V101607 == nil || *p.V101607 != *input.V101607) {
		return false
	}

	if p.V101703 != nil && (input.V101703 == nil || *p.V101703 != *input.V101703) {
		return false
	}

	if p.V101709 != nil && (input.V101709 == nil || *p.V101709 != *input.V101709) {
		return false
	}

	if p.V101803 != nil && (input.V101803 == nil || *p.V101803 != *input.V101803) {
		return false
	}

	if p.V101809 != nil && (input.V101809 == nil || *p.V101809 != *input.V101809) {
		return false
	}

	if p.V101903 != nil && (input.V101903 == nil || *p.V101903 != *input.V101903) {
		return false
	}

	if p.V101909 != nil && (input.V101909 == nil || *p.V101909 != *input.V101909) {
		return false
	}

	if p.V102004 != nil && (input.V102004 == nil || *p.V102004 != *input.V102004) {
		return false
	}

	if p.V1021H1 != nil && (input.V1021H1 == nil || *p.V1021H1 != *input.V1021H1) {
		return false
	}

	if p.V102H20 != nil && (input.V102H20 == nil || *p.V102H20 != *input.V102H20) {
		return false
	}

	if p.V80 != nil && (input.V80 == nil || *p.V80 != *input.V80) {
		return false
	}

	if p.V81 != nil && (input.V81 == nil || *p.V81 != *input.V81) {
		return false
	}

	return true
}
