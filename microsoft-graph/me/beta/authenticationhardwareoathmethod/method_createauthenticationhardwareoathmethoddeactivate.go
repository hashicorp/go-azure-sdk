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

type CreateAuthenticationHardwareOathMethodDeactivateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateAuthenticationHardwareOathMethodDeactivateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAuthenticationHardwareOathMethodDeactivateOperationOptions() CreateAuthenticationHardwareOathMethodDeactivateOperationOptions {
	return CreateAuthenticationHardwareOathMethodDeactivateOperationOptions{}
}

func (o CreateAuthenticationHardwareOathMethodDeactivateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAuthenticationHardwareOathMethodDeactivateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAuthenticationHardwareOathMethodDeactivateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAuthenticationHardwareOathMethodDeactivate - Invoke action deactivate. Deactive a hardware OATH token. It
// remains assigned to a user.
func (c AuthenticationHardwareOathMethodClient) CreateAuthenticationHardwareOathMethodDeactivate(ctx context.Context, id beta.MeAuthenticationHardwareOathMethodId, options CreateAuthenticationHardwareOathMethodDeactivateOperationOptions) (result CreateAuthenticationHardwareOathMethodDeactivateOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/deactivate", id.ID()),
		RetryFunc:     options.RetryFunc,
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
