package groupsitetermstoresetchildrenchildren

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetChildrenChildrenClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetChildrenChildrenClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetChildrenChildrenClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetchildrenchildren", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetChildrenChildrenClient: %+v", err)
	}

	return &GroupSiteTermStoreSetChildrenChildrenClient{
		Client: client,
	}, nil
}
