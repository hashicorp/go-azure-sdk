package grouppolicydefinitionpresentationdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionPresentationDefinitionClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionPresentationDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionPresentationDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionpresentationdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionPresentationDefinitionClient: %+v", err)
	}

	return &GroupPolicyDefinitionPresentationDefinitionClient{
		Client: client,
	}, nil
}
