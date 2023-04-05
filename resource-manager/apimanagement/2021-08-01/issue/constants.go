package issue

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type State string

const (
	StateClosed   State = "closed"
	StateOpen     State = "open"
	StateProposed State = "proposed"
	StateRemoved  State = "removed"
	StateResolved State = "resolved"
)

func PossibleValuesForState() []string {
	return []string{
		string(StateClosed),
		string(StateOpen),
		string(StateProposed),
		string(StateRemoved),
		string(StateResolved),
	}
}
