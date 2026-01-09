package chromeosonboardingsetting

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type ChromeOSOnboardingSettingsOperationPredicate struct {
}

func (p ChromeOSOnboardingSettingsOperationPredicate) Matches(input beta.ChromeOSOnboardingSettings) bool {

	return true
}
