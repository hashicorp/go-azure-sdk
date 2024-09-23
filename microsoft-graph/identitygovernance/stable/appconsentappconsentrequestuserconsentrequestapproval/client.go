package appconsentappconsentrequestuserconsentrequestapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentAppConsentRequestUserConsentRequestApprovalClient struct {
	Client *msgraph.Client
}

func NewAppConsentAppConsentRequestUserConsentRequestApprovalClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentAppConsentRequestUserConsentRequestApprovalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentappconsentrequestuserconsentrequestapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentAppConsentRequestUserConsentRequestApprovalClient: %+v", err)
	}

	return &AppConsentAppConsentRequestUserConsentRequestApprovalClient{
		Client: client,
	}, nil
}
