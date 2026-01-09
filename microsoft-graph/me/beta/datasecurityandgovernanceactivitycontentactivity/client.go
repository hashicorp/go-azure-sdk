package datasecurityandgovernanceactivitycontentactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSecurityAndGovernanceActivityContentActivityClient struct {
	Client *msgraph.Client
}

func NewDataSecurityAndGovernanceActivityContentActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*DataSecurityAndGovernanceActivityContentActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "datasecurityandgovernanceactivitycontentactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataSecurityAndGovernanceActivityContentActivityClient: %+v", err)
	}

	return &DataSecurityAndGovernanceActivityContentActivityClient{
		Client: client,
	}, nil
}
