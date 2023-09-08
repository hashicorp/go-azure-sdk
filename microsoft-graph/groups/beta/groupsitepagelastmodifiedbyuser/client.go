package groupsitepagelastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSitePageLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewGroupSitePageLastModifiedByUserClientWithBaseURI(api sdkEnv.Api) (*GroupSitePageLastModifiedByUserClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitepagelastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSitePageLastModifiedByUserClient: %+v", err)
	}

	return &GroupSitePageLastModifiedByUserClient{
		Client: client,
	}, nil
}
