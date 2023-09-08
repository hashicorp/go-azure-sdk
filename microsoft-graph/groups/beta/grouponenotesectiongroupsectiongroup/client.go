package grouponenotesectiongroupsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteSectionGroupSectionGroupClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteSectionGroupSectionGroupClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteSectionGroupSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotesectiongroupsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteSectionGroupSectionGroupClient: %+v", err)
	}

	return &GroupOnenoteSectionGroupSectionGroupClient{
		Client: client,
	}, nil
}
