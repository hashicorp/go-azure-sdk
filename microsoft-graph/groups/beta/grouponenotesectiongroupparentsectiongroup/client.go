package grouponenotesectiongroupparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteSectionGroupParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteSectionGroupParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteSectionGroupParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotesectiongroupparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteSectionGroupParentSectionGroupClient: %+v", err)
	}

	return &GroupOnenoteSectionGroupParentSectionGroupClient{
		Client: client,
	}, nil
}
