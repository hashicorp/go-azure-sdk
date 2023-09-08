package groupsitelastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewGroupSiteLastModifiedByUserClientWithBaseURI(api sdkEnv.Api) (*GroupSiteLastModifiedByUserClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteLastModifiedByUserClient: %+v", err)
	}

	return &GroupSiteLastModifiedByUserClient{
		Client: client,
	}, nil
}
