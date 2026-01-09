package outlookmastercategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookMasterCategoryClient struct {
	Client *msgraph.Client
}

func NewOutlookMasterCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*OutlookMasterCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "outlookmastercategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OutlookMasterCategoryClient: %+v", err)
	}

	return &OutlookMasterCategoryClient{
		Client: client,
	}, nil
}
