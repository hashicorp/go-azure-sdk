package functions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UdfType string

const (
	UdfTypeScalar UdfType = "Scalar"
)

func PossibleValuesForUdfType() []string {
	return []string{
		string(UdfTypeScalar),
	}
}

type UpdateMode string

const (
	UpdateModeRefreshable UpdateMode = "Refreshable"
	UpdateModeStatic      UpdateMode = "Static"
)

func PossibleValuesForUpdateMode() []string {
	return []string{
		string(UpdateModeRefreshable),
		string(UpdateModeStatic),
	}
}
