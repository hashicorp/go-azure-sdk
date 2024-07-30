package b2xuserflowapiconnectorconfigurationpostfederationsignup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRef ...
func (c B2xUserFlowApiConnectorConfigurationPostFederationSignupClient) RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRef(ctx context.Context, id IdentityB2xUserFlowId) (result RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodDelete,
		Path:       fmt.Sprintf("%s/apiConnectorConfiguration/postFederationSignup/$ref", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
