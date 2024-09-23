package managedappregistration

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetManagedAppRegistrationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.ManagedAppRegistration
}

type GetManagedAppRegistrationOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetManagedAppRegistrationOperationOptions() GetManagedAppRegistrationOperationOptions {
	return GetManagedAppRegistrationOperationOptions{}
}

func (o GetManagedAppRegistrationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetManagedAppRegistrationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetManagedAppRegistrationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetManagedAppRegistration - Get managedAppRegistrations from me. Zero or more managed app registrations that belong
// to the user.
func (c ManagedAppRegistrationClient) GetManagedAppRegistration(ctx context.Context, id beta.MeManagedAppRegistrationId, options GetManagedAppRegistrationOperationOptions) (result GetManagedAppRegistrationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalManagedAppRegistrationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
