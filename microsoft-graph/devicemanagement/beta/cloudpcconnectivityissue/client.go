package cloudpcconnectivityissue

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityIssueClient struct {
	Client *msgraph.Client
}

func NewCloudPCConnectivityIssueClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudPCConnectivityIssueClient, error) {
	client, err := msgraph.NewClient(sdkApi, "cloudpcconnectivityissue", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudPCConnectivityIssueClient: %+v", err)
	}

	return &CloudPCConnectivityIssueClient{
		Client: client,
	}, nil
}
