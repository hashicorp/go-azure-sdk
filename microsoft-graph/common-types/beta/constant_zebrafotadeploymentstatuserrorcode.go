package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentStatusErrorCode string

const (
	ZebraFotaDeploymentStatusErrorCode_NoDevicesFoundInSelectedAadGroups                  ZebraFotaDeploymentStatusErrorCode = "noDevicesFoundInSelectedAadGroups"
	ZebraFotaDeploymentStatusErrorCode_NoIntuneDevicesFoundInSelectedAadGroups            ZebraFotaDeploymentStatusErrorCode = "noIntuneDevicesFoundInSelectedAadGroups"
	ZebraFotaDeploymentStatusErrorCode_NoZebraFotaDevicesFoundForSelectedDeviceModel      ZebraFotaDeploymentStatusErrorCode = "noZebraFotaDevicesFoundForSelectedDeviceModel"
	ZebraFotaDeploymentStatusErrorCode_NoZebraFotaEnrolledDevicesFoundForCurrentTenant    ZebraFotaDeploymentStatusErrorCode = "noZebraFotaEnrolledDevicesFoundForCurrentTenant"
	ZebraFotaDeploymentStatusErrorCode_NoZebraFotaEnrolledDevicesFoundInSelectedAadGroups ZebraFotaDeploymentStatusErrorCode = "noZebraFotaEnrolledDevicesFoundInSelectedAadGroups"
	ZebraFotaDeploymentStatusErrorCode_Success                                            ZebraFotaDeploymentStatusErrorCode = "success"
	ZebraFotaDeploymentStatusErrorCode_ZebraFotaCreateDeploymentRequestFailure            ZebraFotaDeploymentStatusErrorCode = "zebraFotaCreateDeploymentRequestFailure"
)

func PossibleValuesForZebraFotaDeploymentStatusErrorCode() []string {
	return []string{
		string(ZebraFotaDeploymentStatusErrorCode_NoDevicesFoundInSelectedAadGroups),
		string(ZebraFotaDeploymentStatusErrorCode_NoIntuneDevicesFoundInSelectedAadGroups),
		string(ZebraFotaDeploymentStatusErrorCode_NoZebraFotaDevicesFoundForSelectedDeviceModel),
		string(ZebraFotaDeploymentStatusErrorCode_NoZebraFotaEnrolledDevicesFoundForCurrentTenant),
		string(ZebraFotaDeploymentStatusErrorCode_NoZebraFotaEnrolledDevicesFoundInSelectedAadGroups),
		string(ZebraFotaDeploymentStatusErrorCode_Success),
		string(ZebraFotaDeploymentStatusErrorCode_ZebraFotaCreateDeploymentRequestFailure),
	}
}

func (s *ZebraFotaDeploymentStatusErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZebraFotaDeploymentStatusErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZebraFotaDeploymentStatusErrorCode(input string) (*ZebraFotaDeploymentStatusErrorCode, error) {
	vals := map[string]ZebraFotaDeploymentStatusErrorCode{
		"nodevicesfoundinselectedaadgroups":                  ZebraFotaDeploymentStatusErrorCode_NoDevicesFoundInSelectedAadGroups,
		"nointunedevicesfoundinselectedaadgroups":            ZebraFotaDeploymentStatusErrorCode_NoIntuneDevicesFoundInSelectedAadGroups,
		"nozebrafotadevicesfoundforselecteddevicemodel":      ZebraFotaDeploymentStatusErrorCode_NoZebraFotaDevicesFoundForSelectedDeviceModel,
		"nozebrafotaenrolleddevicesfoundforcurrenttenant":    ZebraFotaDeploymentStatusErrorCode_NoZebraFotaEnrolledDevicesFoundForCurrentTenant,
		"nozebrafotaenrolleddevicesfoundinselectedaadgroups": ZebraFotaDeploymentStatusErrorCode_NoZebraFotaEnrolledDevicesFoundInSelectedAadGroups,
		"success": ZebraFotaDeploymentStatusErrorCode_Success,
		"zebrafotacreatedeploymentrequestfailure": ZebraFotaDeploymentStatusErrorCode_ZebraFotaCreateDeploymentRequestFailure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaDeploymentStatusErrorCode(input)
	return &out, nil
}
