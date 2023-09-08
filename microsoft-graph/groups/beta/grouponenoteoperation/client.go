package grouponenoteoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteOperationClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteOperationClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteOperationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenoteoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteOperationClient: %+v", err)
	}

	return &GroupOnenoteOperationClient{
		Client: client,
	}, nil
}
