package cloudpcroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewCloudPCRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudPCRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "cloudpcroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudPCRoleDefinitionClient: %+v", err)
	}

	return &CloudPCRoleDefinitionClient{
		Client: client,
	}, nil
}
