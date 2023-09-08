package groupsitelistlastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListLastModifiedByUserClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListLastModifiedByUserClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistlastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListLastModifiedByUserClient: %+v", err)
	}

	return &GroupSiteListLastModifiedByUserClient{
		Client: client,
	}, nil
}
