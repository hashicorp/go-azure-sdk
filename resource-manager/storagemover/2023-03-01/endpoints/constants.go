package endpoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointType string

const (
	EndpointTypeAzureStorageBlobContainer EndpointType = "AzureStorageBlobContainer"
	EndpointTypeNfsMount                  EndpointType = "NfsMount"
)

func PossibleValuesForEndpointType() []string {
	return []string{
		string(EndpointTypeAzureStorageBlobContainer),
		string(EndpointTypeNfsMount),
	}
}

func (s *EndpointType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForEndpointType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = EndpointType(decoded)
	return nil
}

type NfsVersion string

const (
	NfsVersionNFSauto   NfsVersion = "NFSauto"
	NfsVersionNFSvFour  NfsVersion = "NFSv4"
	NfsVersionNFSvThree NfsVersion = "NFSv3"
)

func PossibleValuesForNfsVersion() []string {
	return []string{
		string(NfsVersionNFSauto),
		string(NfsVersionNFSvFour),
		string(NfsVersionNFSvThree),
	}
}

func (s *NfsVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForNfsVersion() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = NfsVersion(decoded)
	return nil
}

type ProvisioningState string

const (
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateSucceeded),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ProvisioningState(decoded)
	return nil
}
