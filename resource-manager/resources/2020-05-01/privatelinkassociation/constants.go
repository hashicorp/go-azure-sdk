package privatelinkassociation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublicNetworkAccessOptions string

const (
	PublicNetworkAccessOptionsDisabled PublicNetworkAccessOptions = "Disabled"
	PublicNetworkAccessOptionsEnabled  PublicNetworkAccessOptions = "Enabled"
)

func PossibleValuesForPublicNetworkAccessOptions() []string {
	return []string{
		string(PublicNetworkAccessOptionsDisabled),
		string(PublicNetworkAccessOptionsEnabled),
	}
}
