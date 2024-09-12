package appconsentappconsentrequestuserconsentrequestapprovalstep

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentAppConsentRequestUserConsentRequestApprovalStepClient struct {
	Client *msgraph.Client
}

func NewAppConsentAppConsentRequestUserConsentRequestApprovalStepClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentAppConsentRequestUserConsentRequestApprovalStepClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentappconsentrequestuserconsentrequestapprovalstep", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentAppConsentRequestUserConsentRequestApprovalStepClient: %+v", err)
	}

	return &AppConsentAppConsentRequestUserConsentRequestApprovalStepClient{
		Client: client,
	}, nil
}
