package userapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserApprovalClient struct {
	Client *msgraph.Client
}

func NewUserApprovalClientWithBaseURI(api sdkEnv.Api) (*UserApprovalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserApprovalClient: %+v", err)
	}

	return &UserApprovalClient{
		Client: client,
	}, nil
}
