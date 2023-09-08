package setgroupsitetermstoregroupsettermrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreGroupSetTermRelationClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreGroupSetTermRelationClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreGroupSetTermRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoregroupsettermrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreGroupSetTermRelationClient: %+v", err)
	}

	return &SetGroupSiteTermStoreGroupSetTermRelationClient{
		Client: client,
	}, nil
}
