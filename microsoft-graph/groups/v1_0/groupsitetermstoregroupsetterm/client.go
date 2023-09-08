package groupsitetermstoregroupsetterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsetterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetTermClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetTermClient{
		Client: client,
	}, nil
}
