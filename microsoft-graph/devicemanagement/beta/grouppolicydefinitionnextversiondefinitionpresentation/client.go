package grouppolicydefinitionnextversiondefinitionpresentation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionNextVersionDefinitionPresentationClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionNextVersionDefinitionPresentationClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionNextVersionDefinitionPresentationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionnextversiondefinitionpresentation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionNextVersionDefinitionPresentationClient: %+v", err)
	}

	return &GroupPolicyDefinitionNextVersionDefinitionPresentationClient{
		Client: client,
	}, nil
}
