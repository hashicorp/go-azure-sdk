package groupsiteonenotepage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteOnenotePageClient struct {
	Client *msgraph.Client
}

func NewGroupSiteOnenotePageClientWithBaseURI(api sdkEnv.Api) (*GroupSiteOnenotePageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteonenotepage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteOnenotePageClient: %+v", err)
	}

	return &GroupSiteOnenotePageClient{
		Client: client,
	}, nil
}
