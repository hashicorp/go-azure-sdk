package signin

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSignInConfirmCompromisedOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateSignInConfirmCompromisedOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateSignInConfirmCompromisedOperationOptions() CreateSignInConfirmCompromisedOperationOptions {
	return CreateSignInConfirmCompromisedOperationOptions{}
}

func (o CreateSignInConfirmCompromisedOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSignInConfirmCompromisedOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSignInConfirmCompromisedOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSignInConfirmCompromised - Invoke action confirmCompromised. Allow admins to mark an event in the Microsoft
// Entra sign-in logs as risky. Events marked as risky by an admin are immediately flagged as high risk in Microsoft
// Entra ID Protection, overriding previous risk states. Admins can confirm that events flagged as risky by Microsoft
// Entra ID Protection are in fact risky. For details about investigating Identity Protection risks, see How to
// investigate risk.
func (c SignInClient) CreateSignInConfirmCompromised(ctx context.Context, input CreateSignInConfirmCompromisedRequest, options CreateSignInConfirmCompromisedOperationOptions) (result CreateSignInConfirmCompromisedOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/auditLogs/signIns/confirmCompromised",
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
