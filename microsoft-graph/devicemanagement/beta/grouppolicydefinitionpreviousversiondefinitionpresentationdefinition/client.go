package grouppolicydefinitionpreviousversiondefinitionpresentationdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionPreviousVersionDefinitionPresentationDefinitionClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionPreviousVersionDefinitionPresentationDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionPreviousVersionDefinitionPresentationDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionpreviousversiondefinitionpresentationdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionPreviousVersionDefinitionPresentationDefinitionClient: %+v", err)
	}

	return &GroupPolicyDefinitionPreviousVersionDefinitionPresentationDefinitionClient{
		Client: client,
	}, nil
}
