package azureadadministrators

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdministratorType string

const (
	AdministratorTypeActiveDirectory AdministratorType = "ActiveDirectory"
)

func PossibleValuesForAdministratorType() []string {
	return []string{
		string(AdministratorTypeActiveDirectory),
	}
}
