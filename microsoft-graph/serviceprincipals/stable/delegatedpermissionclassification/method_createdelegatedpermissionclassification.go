package delegatedpermissionclassification

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDelegatedPermissionClassificationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DelegatedPermissionClassification
}

type CreateDelegatedPermissionClassificationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateDelegatedPermissionClassificationOperationOptions() CreateDelegatedPermissionClassificationOperationOptions {
	return CreateDelegatedPermissionClassificationOperationOptions{}
}

func (o CreateDelegatedPermissionClassificationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDelegatedPermissionClassificationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDelegatedPermissionClassificationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDelegatedPermissionClassification - Create delegatedPermissionClassification. Classify a delegated permission
// by adding a delegatedPermissionClassification to the servicePrincipal representing the API.
func (c DelegatedPermissionClassificationClient) CreateDelegatedPermissionClassification(ctx context.Context, id stable.ServicePrincipalId, input stable.DelegatedPermissionClassification, options CreateDelegatedPermissionClassificationOperationOptions) (result CreateDelegatedPermissionClassificationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/delegatedPermissionClassifications", id.ID()),
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

	var model stable.DelegatedPermissionClassification
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
