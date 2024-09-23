package credentialuserregistrationdetail

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCredentialUserRegistrationDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CredentialUserRegistrationDetails
}

type GetCredentialUserRegistrationDetailOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetCredentialUserRegistrationDetailOperationOptions() GetCredentialUserRegistrationDetailOperationOptions {
	return GetCredentialUserRegistrationDetailOperationOptions{}
}

func (o GetCredentialUserRegistrationDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCredentialUserRegistrationDetailOperationOptions) ToOData() *odata.Query {
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

func (o GetCredentialUserRegistrationDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetCredentialUserRegistrationDetail - Get credentialUserRegistrationDetails from reports. Details of the usage of
// self-service password reset and multifactor authentication (MFA) for all registered users.
func (c CredentialUserRegistrationDetailClient) GetCredentialUserRegistrationDetail(ctx context.Context, id beta.ReportCredentialUserRegistrationDetailId, options GetCredentialUserRegistrationDetailOperationOptions) (result GetCredentialUserRegistrationDetailOperationResponse, err error) {
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

	var model beta.CredentialUserRegistrationDetails
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
