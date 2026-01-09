package datasecurityandgovernance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSecurityAndGovernanceClient struct {
	Client *msgraph.Client
}

func NewDataSecurityAndGovernanceClientWithBaseURI(sdkApi sdkEnv.Api) (*DataSecurityAndGovernanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "datasecurityandgovernance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataSecurityAndGovernanceClient: %+v", err)
	}

	return &DataSecurityAndGovernanceClient{
		Client: client,
	}, nil
}
