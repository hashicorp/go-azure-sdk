package setgroupsitetermstoresetchildren

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreSetChildrenClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreSetChildrenClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreSetChildrenClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoresetchildren", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreSetChildrenClient: %+v", err)
	}

	return &SetGroupSiteTermStoreSetChildrenClient{
		Client: client,
	}, nil
}
