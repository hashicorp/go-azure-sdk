package groupsitelistoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListOperationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListOperationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListOperationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListOperationClient: %+v", err)
	}

	return &GroupSiteListOperationClient{
		Client: client,
	}, nil
}
