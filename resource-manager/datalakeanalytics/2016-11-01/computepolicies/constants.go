package computepolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AADObjectType string

const (
	AADObjectTypeGroup            AADObjectType = "Group"
	AADObjectTypeServicePrincipal AADObjectType = "ServicePrincipal"
	AADObjectTypeUser             AADObjectType = "User"
)

func PossibleValuesForAADObjectType() []string {
	return []string{
		string(AADObjectTypeGroup),
		string(AADObjectTypeServicePrincipal),
		string(AADObjectTypeUser),
	}
}
