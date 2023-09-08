package grouponenotesectiongroupsectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteSectionGroupSectionPageClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteSectionGroupSectionPageClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteSectionGroupSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenotesectiongroupsectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteSectionGroupSectionPageClient: %+v", err)
	}

	return &GroupOnenoteSectionGroupSectionPageClient{
		Client: client,
	}, nil
}
