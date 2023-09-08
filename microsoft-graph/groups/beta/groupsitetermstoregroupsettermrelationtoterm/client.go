package groupsitetermstoregroupsettermrelationtoterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetTermRelationToTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetTermRelationToTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetTermRelationToTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsettermrelationtoterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetTermRelationToTermClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetTermRelationToTermClient{
		Client: client,
	}, nil
}
