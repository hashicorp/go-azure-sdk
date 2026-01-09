package datasharingconsent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSharingConsentClient struct {
	Client *msgraph.Client
}

func NewDataSharingConsentClientWithBaseURI(sdkApi sdkEnv.Api) (*DataSharingConsentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "datasharingconsent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataSharingConsentClient: %+v", err)
	}

	return &DataSharingConsentClient{
		Client: client,
	}, nil
}
