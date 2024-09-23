package grouppolicydefinitionnextversiondefinitionpresentationdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionNextVersionDefinitionPresentationDefinitionClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionNextVersionDefinitionPresentationDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionNextVersionDefinitionPresentationDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionnextversiondefinitionpresentationdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionNextVersionDefinitionPresentationDefinitionClient: %+v", err)
	}

	return &GroupPolicyDefinitionNextVersionDefinitionPresentationDefinitionClient{
		Client: client,
	}, nil
}
