package key

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Serial int64

const (
	SerialOne Serial = 1
	SerialTwo Serial = 2
)

func PossibleValuesForSerial() []int64 {
	return []int64{
		int64(SerialOne),
		int64(SerialTwo),
	}
}

func parseSerial(input int64) (*Serial, error) {
	vals := map[int64]Serial{
		1: SerialOne,
		2: SerialTwo,
	}
	if v, ok := vals[input]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Serial(input)
	return &out, nil
}
