package grouppolicydefinitionpreviousversiondefinitionnextversiondefinitiondefinitionfile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionDefinitionFileClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionDefinitionFileClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionDefinitionFileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionpreviousversiondefinitionnextversiondefinitiondefinitionfile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionDefinitionFileClient: %+v", err)
	}

	return &GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionDefinitionFileClient{
		Client: client,
	}, nil
}
