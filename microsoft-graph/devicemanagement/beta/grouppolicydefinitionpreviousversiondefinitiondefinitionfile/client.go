package grouppolicydefinitionpreviousversiondefinitiondefinitionfile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionPreviousVersionDefinitionDefinitionFileClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionPreviousVersionDefinitionDefinitionFileClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionPreviousVersionDefinitionDefinitionFileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionpreviousversiondefinitiondefinitionfile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionPreviousVersionDefinitionDefinitionFileClient: %+v", err)
	}

	return &GroupPolicyDefinitionPreviousVersionDefinitionDefinitionFileClient{
		Client: client,
	}, nil
}
