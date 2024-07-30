package chromeosonboardingsetting

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type ChromeOSOnboardingSettingsOperationPredicate struct {
}

func (p ChromeOSOnboardingSettingsOperationPredicate) Matches(input beta.ChromeOSOnboardingSettings) bool {

	return true
}
