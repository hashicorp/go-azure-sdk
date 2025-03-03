package key

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Serial int64

const (
	SerialPrimary   Serial = 1
	SerialSecondary Serial = 2
)

func PossibleValuesForSerial() []int64 {
	return []int64{
		int64(SerialPrimary),
		int64(SerialSecondary),
	}
}
