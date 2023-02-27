// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymlanalyticssettings

type SecurityMLAnalyticsSettingOperationPredicate struct {
}

func (p SecurityMLAnalyticsSettingOperationPredicate) Matches(input SecurityMLAnalyticsSetting) bool {

	return true
}
