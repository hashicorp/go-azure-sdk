package grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentationdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationDefinitionClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentationdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationDefinitionClient: %+v", err)
	}

	return &GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationDefinitionClient{
		Client: client,
	}, nil
}
