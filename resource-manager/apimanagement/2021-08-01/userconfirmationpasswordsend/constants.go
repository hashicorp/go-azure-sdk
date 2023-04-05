package userconfirmationpasswordsend

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppType string

const (
	AppTypeDeveloperPortal AppType = "developerPortal"
	AppTypePortal          AppType = "portal"
)

func PossibleValuesForAppType() []string {
	return []string{
		string(AppTypeDeveloperPortal),
		string(AppTypePortal),
	}
}
