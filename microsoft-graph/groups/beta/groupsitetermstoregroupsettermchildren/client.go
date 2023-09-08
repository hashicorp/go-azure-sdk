package groupsitetermstoregroupsettermchildren

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetTermChildrenClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetTermChildrenClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetTermChildrenClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsettermchildren", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetTermChildrenClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetTermChildrenClient{
		Client: client,
	}, nil
}
