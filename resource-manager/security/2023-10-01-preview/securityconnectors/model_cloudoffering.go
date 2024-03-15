package securityconnectors

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudOffering interface {
}

// RawCloudOfferingImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCloudOfferingImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalCloudOfferingImplementation(input []byte) (CloudOffering, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudOffering into map[string]interface: %+v", err)
	}

	value, ok := temp["offeringType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "CspmMonitorAws") {
		var out CspmMonitorAwsOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CspmMonitorAwsOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "CspmMonitorAzureDevOps") {
		var out CspmMonitorAzureDevOpsOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CspmMonitorAzureDevOpsOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "CspmMonitorGcp") {
		var out CspmMonitorGcpOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CspmMonitorGcpOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "CspmMonitorGitLab") {
		var out CspmMonitorGitLabOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CspmMonitorGitLabOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "CspmMonitorGithub") {
		var out CspmMonitorGithubOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CspmMonitorGithubOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DefenderCspmAws") {
		var out DefenderCspmAwsOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefenderCspmAwsOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DefenderCspmGcp") {
		var out DefenderCspmGcpOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefenderCspmGcpOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DefenderForDatabasesAws") {
		var out DefenderFoDatabasesAwsOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefenderFoDatabasesAwsOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DefenderForContainersAws") {
		var out DefenderForContainersAwsOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefenderForContainersAwsOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DefenderForContainersGcp") {
		var out DefenderForContainersGcpOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefenderForContainersGcpOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DefenderForDatabasesGcp") {
		var out DefenderForDatabasesGcpOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefenderForDatabasesGcpOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DefenderForDevOpsAzureDevOps") {
		var out DefenderForDevOpsAzureDevOpsOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefenderForDevOpsAzureDevOpsOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DefenderForDevOpsGitLab") {
		var out DefenderForDevOpsGitLabOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefenderForDevOpsGitLabOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DefenderForDevOpsGithub") {
		var out DefenderForDevOpsGithubOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefenderForDevOpsGithubOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DefenderForServersAws") {
		var out DefenderForServersAwsOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefenderForServersAwsOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DefenderForServersGcp") {
		var out DefenderForServersGcpOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefenderForServersGcpOffering: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "InformationProtectionAws") {
		var out InformationProtectionAwsOffering
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InformationProtectionAwsOffering: %+v", err)
		}
		return out, nil
	}

	out := RawCloudOfferingImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
