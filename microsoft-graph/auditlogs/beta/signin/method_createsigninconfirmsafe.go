package signin

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSignInConfirmSafeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateSignInConfirmSafeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateSignInConfirmSafeOperationOptions() CreateSignInConfirmSafeOperationOptions {
	return CreateSignInConfirmSafeOperationOptions{}
}

func (o CreateSignInConfirmSafeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSignInConfirmSafeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSignInConfirmSafeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSignInConfirmSafe - Invoke action confirmSafe. Allow admins to mark an event in Microsoft Entra sign-in logs as
// safe. Admins can either mark the events flagged as risky by Microsoft Entra ID Protection as safe, or they can mark
// unflagged events as safe. For details about investigating Identity Protection risks, see How to investigate risk.
func (c SignInClient) CreateSignInConfirmSafe(ctx context.Context, input CreateSignInConfirmSafeRequest, options CreateSignInConfirmSafeOperationOptions) (result CreateSignInConfirmSafeOperationResponse, err error) {
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
		Path:          "/auditLogs/signIns/confirmSafe",
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
