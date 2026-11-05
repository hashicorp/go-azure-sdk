package managedenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckMigrationEligibilityResponse struct {
	CurrentMode     string                              `json:"currentMode"`
	EnvironmentName string                              `json:"environmentName"`
	FailureReasons  []MigrationEligibilityFailureReason `json:"failureReasons"`
	IsEligible      bool                                `json:"isEligible"`
	TargetMode      string                              `json:"targetMode"`
}
