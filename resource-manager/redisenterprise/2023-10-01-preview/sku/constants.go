package sku

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkuName string

const (
	SkuNameEnterpriseEFive                SkuName = "Enterprise_E5"
	SkuNameEnterpriseEFiveZero            SkuName = "Enterprise_E50"
	SkuNameEnterpriseEFourHundred         SkuName = "Enterprise_E400"
	SkuNameEnterpriseEOneHundred          SkuName = "Enterprise_E100"
	SkuNameEnterpriseEOneZero             SkuName = "Enterprise_E10"
	SkuNameEnterpriseETwoHundred          SkuName = "Enterprise_E200"
	SkuNameEnterpriseETwoZero             SkuName = "Enterprise_E20"
	SkuNameEnterpriseFlashFOneFiveHundred SkuName = "EnterpriseFlash_F1500"
	SkuNameEnterpriseFlashFSevenHundred   SkuName = "EnterpriseFlash_F700"
	SkuNameEnterpriseFlashFThreeHundred   SkuName = "EnterpriseFlash_F300"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameEnterpriseEFive),
		string(SkuNameEnterpriseEFiveZero),
		string(SkuNameEnterpriseEFourHundred),
		string(SkuNameEnterpriseEOneHundred),
		string(SkuNameEnterpriseEOneZero),
		string(SkuNameEnterpriseETwoHundred),
		string(SkuNameEnterpriseETwoZero),
		string(SkuNameEnterpriseFlashFOneFiveHundred),
		string(SkuNameEnterpriseFlashFSevenHundred),
		string(SkuNameEnterpriseFlashFThreeHundred),
	}
}

func (s *SkuName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkuName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkuName(input string) (*SkuName, error) {
	vals := map[string]SkuName{
		"enterprise_e5":         SkuNameEnterpriseEFive,
		"enterprise_e50":        SkuNameEnterpriseEFiveZero,
		"enterprise_e400":       SkuNameEnterpriseEFourHundred,
		"enterprise_e100":       SkuNameEnterpriseEOneHundred,
		"enterprise_e10":        SkuNameEnterpriseEOneZero,
		"enterprise_e200":       SkuNameEnterpriseETwoHundred,
		"enterprise_e20":        SkuNameEnterpriseETwoZero,
		"enterpriseflash_f1500": SkuNameEnterpriseFlashFOneFiveHundred,
		"enterpriseflash_f700":  SkuNameEnterpriseFlashFSevenHundred,
		"enterpriseflash_f300":  SkuNameEnterpriseFlashFThreeHundred,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuName(input)
	return &out, nil
}
