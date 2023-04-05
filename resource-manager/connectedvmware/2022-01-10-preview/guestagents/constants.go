package guestagents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningAction string

const (
	ProvisioningActionInstall   ProvisioningAction = "install"
	ProvisioningActionRepair    ProvisioningAction = "repair"
	ProvisioningActionUninstall ProvisioningAction = "uninstall"
)

func PossibleValuesForProvisioningAction() []string {
	return []string{
		string(ProvisioningActionInstall),
		string(ProvisioningActionRepair),
		string(ProvisioningActionUninstall),
	}
}

func (s *ProvisioningAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForProvisioningAction() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ProvisioningAction(decoded)
	return nil
}
