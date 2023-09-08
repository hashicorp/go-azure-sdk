package setgroupsitetermstoresettermrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreSetTermRelationClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreSetTermRelationClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreSetTermRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoresettermrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreSetTermRelationClient: %+v", err)
	}

	return &SetGroupSiteTermStoreSetTermRelationClient{
		Client: client,
	}, nil
}
