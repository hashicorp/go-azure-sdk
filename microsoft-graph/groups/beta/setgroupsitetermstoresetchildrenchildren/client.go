package setgroupsitetermstoresetchildrenchildren

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreSetChildrenChildrenClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreSetChildrenChildrenClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreSetChildrenChildrenClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoresetchildrenchildren", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreSetChildrenChildrenClient: %+v", err)
	}

	return &SetGroupSiteTermStoreSetChildrenChildrenClient{
		Client: client,
	}, nil
}
