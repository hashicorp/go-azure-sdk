package grouponenotesectiongroupsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteSectionGroupSectionClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteSectionGroupSectionClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteSectionGroupSectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotesectiongroupsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteSectionGroupSectionClient: %+v", err)
	}

	return &GroupOnenoteSectionGroupSectionClient{
		Client: client,
	}, nil
}
