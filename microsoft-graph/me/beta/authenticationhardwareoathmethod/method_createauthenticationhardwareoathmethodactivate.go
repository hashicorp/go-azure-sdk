package authenticationhardwareoathmethod

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAuthenticationHardwareOathMethodActivateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateAuthenticationHardwareOathMethodActivateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAuthenticationHardwareOathMethodActivateOperationOptions() CreateAuthenticationHardwareOathMethodActivateOperationOptions {
	return CreateAuthenticationHardwareOathMethodActivateOperationOptions{}
}

func (o CreateAuthenticationHardwareOathMethodActivateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAuthenticationHardwareOathMethodActivateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAuthenticationHardwareOathMethodActivateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAuthenticationHardwareOathMethodActivate - Invoke action activate. Activate a hardware OATH token that is
// already assigned to a user. A user can self-activate their token or an admin can activate for a user.
func (c AuthenticationHardwareOathMethodClient) CreateAuthenticationHardwareOathMethodActivate(ctx context.Context, id beta.MeAuthenticationHardwareOathMethodId, input CreateAuthenticationHardwareOathMethodActivateRequest, options CreateAuthenticationHardwareOathMethodActivateOperationOptions) (result CreateAuthenticationHardwareOathMethodActivateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/activate", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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
