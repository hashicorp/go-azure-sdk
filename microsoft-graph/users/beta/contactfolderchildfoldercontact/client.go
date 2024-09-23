package contactfolderchildfoldercontact

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactFolderChildFolderContactClient struct {
	Client *msgraph.Client
}

func NewContactFolderChildFolderContactClientWithBaseURI(sdkApi sdkEnv.Api) (*ContactFolderChildFolderContactClient, error) {
	client, err := msgraph.NewClient(sdkApi, "contactfolderchildfoldercontact", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContactFolderChildFolderContactClient: %+v", err)
	}

	return &ContactFolderChildFolderContactClient{
		Client: client,
	}, nil
}
