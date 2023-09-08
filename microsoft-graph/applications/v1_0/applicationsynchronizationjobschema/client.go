package applicationsynchronizationjobschema

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationSynchronizationJobSchemaClient struct {
	Client *msgraph.Client
}

func NewApplicationSynchronizationJobSchemaClientWithBaseURI(api sdkEnv.Api) (*ApplicationSynchronizationJobSchemaClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "applicationsynchronizationjobschema", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationSynchronizationJobSchemaClient: %+v", err)
	}

	return &ApplicationSynchronizationJobSchemaClient{
		Client: client,
	}, nil
}
