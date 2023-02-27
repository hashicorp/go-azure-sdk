// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synchronizationsetting

type SynchronizationSettingOperationPredicate struct {
}

func (p SynchronizationSettingOperationPredicate) Matches(input SynchronizationSetting) bool {

	return true
}
