package healthbots

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkuName string

const (
	SkuNameCZero SkuName = "C0"
	SkuNameFZero SkuName = "F0"
	SkuNameSOne  SkuName = "S1"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameCZero),
		string(SkuNameFZero),
		string(SkuNameSOne),
	}
}
