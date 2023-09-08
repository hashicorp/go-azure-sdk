package groupsiteoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOperationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOperationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOperationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOperationClient: %+v", err)
	}

	return &GroupSiteOperationClient{
		Client: client,
	}, nil
}
