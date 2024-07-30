package virtualendpointusersettingassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointUserSettingAssignmentClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointUserSettingAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointUserSettingAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointusersettingassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointUserSettingAssignmentClient: %+v", err)
	}

	return &VirtualEndpointUserSettingAssignmentClient{
		Client: client,
	}, nil
}
