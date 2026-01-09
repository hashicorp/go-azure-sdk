package servicenowconnection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceNowConnectionClient struct {
	Client *msgraph.Client
}

func NewServiceNowConnectionClientWithBaseURI(sdkApi sdkEnv.Api) (*ServiceNowConnectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "servicenowconnection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServiceNowConnectionClient: %+v", err)
	}

	return &ServiceNowConnectionClient{
		Client: client,
	}, nil
}
