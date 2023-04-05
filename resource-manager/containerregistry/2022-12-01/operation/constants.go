package operation

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerRegistryResourceType string

const (
	ContainerRegistryResourceTypeMicrosoftPointContainerRegistryRegistries ContainerRegistryResourceType = "Microsoft.ContainerRegistry/registries"
)

func PossibleValuesForContainerRegistryResourceType() []string {
	return []string{
		string(ContainerRegistryResourceTypeMicrosoftPointContainerRegistryRegistries),
	}
}

func (s *ContainerRegistryResourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForContainerRegistryResourceType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ContainerRegistryResourceType(decoded)
	return nil
}
