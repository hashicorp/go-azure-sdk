package webtestsapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebTestKind string

const (
	WebTestKindMultistep WebTestKind = "multistep"
	WebTestKindPing      WebTestKind = "ping"
	WebTestKindStandard  WebTestKind = "standard"
)

func PossibleValuesForWebTestKind() []string {
	return []string{
		string(WebTestKindMultistep),
		string(WebTestKindPing),
		string(WebTestKindStandard),
	}
}
