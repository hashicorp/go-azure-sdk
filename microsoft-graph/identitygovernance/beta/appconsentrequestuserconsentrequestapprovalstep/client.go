package appconsentrequestuserconsentrequestapprovalstep

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentRequestUserConsentRequestApprovalStepClient struct {
	Client *msgraph.Client
}

func NewAppConsentRequestUserConsentRequestApprovalStepClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentRequestUserConsentRequestApprovalStepClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentrequestuserconsentrequestapprovalstep", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentRequestUserConsentRequestApprovalStepClient: %+v", err)
	}

	return &AppConsentRequestUserConsentRequestApprovalStepClient{
		Client: client,
	}, nil
}
