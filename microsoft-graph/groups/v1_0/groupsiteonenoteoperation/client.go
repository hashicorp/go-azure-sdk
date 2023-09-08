package groupsiteonenoteoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenoteOperationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenoteOperationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenoteOperationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenoteoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenoteOperationClient: %+v", err)
	}

	return &GroupSiteOnenoteOperationClient{
		Client: client,
	}, nil
}
