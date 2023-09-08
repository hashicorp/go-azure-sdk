package grouponenotesectiongroupsectionpagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteSectionGroupSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteSectionGroupSectionPageContentClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteSectionGroupSectionPageContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotesectiongroupsectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteSectionGroupSectionPageContentClient: %+v", err)
	}

	return &GroupOnenoteSectionGroupSectionPageContentClient{
		Client: client,
	}, nil
}
