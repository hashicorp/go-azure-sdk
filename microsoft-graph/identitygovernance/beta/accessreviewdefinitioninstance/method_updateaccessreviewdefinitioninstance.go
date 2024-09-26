package accessreviewdefinitioninstance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAccessReviewDefinitionInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAccessReviewDefinitionInstanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateAccessReviewDefinitionInstanceOperationOptions() UpdateAccessReviewDefinitionInstanceOperationOptions {
	return UpdateAccessReviewDefinitionInstanceOperationOptions{}
}

func (o UpdateAccessReviewDefinitionInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAccessReviewDefinitionInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAccessReviewDefinitionInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAccessReviewDefinitionInstance - Update accessReviewInstance. Update the properties of an accessReviewInstance
// object. Only the reviewers and fallbackReviewers properties can be updated but the scope property is also required in
// the request body. You can only add reviewers to the fallbackReviewers property but can't remove existing
// fallbackReviewers. To update an accessReviewInstance, its status must be InProgress.
func (c AccessReviewDefinitionInstanceClient) UpdateAccessReviewDefinitionInstance(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceId, input beta.AccessReviewInstance, options UpdateAccessReviewDefinitionInstanceOperationOptions) (result UpdateAccessReviewDefinitionInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
