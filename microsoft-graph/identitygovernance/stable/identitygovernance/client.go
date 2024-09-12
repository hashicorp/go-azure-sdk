package identitygovernance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceClient struct {
	Client *msgraph.Client
}

func NewIdentityGovernanceClientWithBaseURI(sdkApi sdkEnv.Api) (*IdentityGovernanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "identitygovernance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityGovernanceClient: %+v", err)
	}

	return &IdentityGovernanceClient{
		Client: client,
	}, nil
}
