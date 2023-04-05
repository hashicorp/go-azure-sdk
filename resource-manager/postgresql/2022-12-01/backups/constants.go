package backups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Origin string

const (
	OriginFull Origin = "Full"
)

func PossibleValuesForOrigin() []string {
	return []string{
		string(OriginFull),
	}
}
