package grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentationdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationDefinitionClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentationdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationDefinitionClient: %+v", err)
	}

	return &GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationDefinitionClient{
		Client: client,
	}, nil
}
