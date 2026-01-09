package androidmanagedstoreaccountenterprisesetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAccountEnterpriseSettingClient struct {
	Client *msgraph.Client
}

func NewAndroidManagedStoreAccountEnterpriseSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*AndroidManagedStoreAccountEnterpriseSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "androidmanagedstoreaccountenterprisesetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AndroidManagedStoreAccountEnterpriseSettingClient: %+v", err)
	}

	return &AndroidManagedStoreAccountEnterpriseSettingClient{
		Client: client,
	}, nil
}
