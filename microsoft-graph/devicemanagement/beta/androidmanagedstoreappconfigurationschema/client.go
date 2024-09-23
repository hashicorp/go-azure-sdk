package androidmanagedstoreappconfigurationschema

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppConfigurationSchemaClient struct {
	Client *msgraph.Client
}

func NewAndroidManagedStoreAppConfigurationSchemaClientWithBaseURI(sdkApi sdkEnv.Api) (*AndroidManagedStoreAppConfigurationSchemaClient, error) {
	client, err := msgraph.NewClient(sdkApi, "androidmanagedstoreappconfigurationschema", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AndroidManagedStoreAppConfigurationSchemaClient: %+v", err)
	}

	return &AndroidManagedStoreAppConfigurationSchemaClient{
		Client: client,
	}, nil
}
