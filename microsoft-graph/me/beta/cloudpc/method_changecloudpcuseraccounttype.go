package cloudpc

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

type ChangeCloudPCUserAccountTypeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ChangeCloudPCUserAccountTypeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultChangeCloudPCUserAccountTypeOperationOptions() ChangeCloudPCUserAccountTypeOperationOptions {
	return ChangeCloudPCUserAccountTypeOperationOptions{}
}

func (o ChangeCloudPCUserAccountTypeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ChangeCloudPCUserAccountTypeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ChangeCloudPCUserAccountTypeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ChangeCloudPCUserAccountType - Invoke action changeUserAccountType. Change the account type of the user on a specific
// Cloud PC.
func (c CloudPCClient) ChangeCloudPCUserAccountType(ctx context.Context, id beta.MeCloudPCId, input ChangeCloudPCUserAccountTypeRequest, options ChangeCloudPCUserAccountTypeOperationOptions) (result ChangeCloudPCUserAccountTypeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/changeUserAccountType", id.ID()),
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
