package member

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

type AddMemberRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddMemberRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddMemberRefOperationOptions() AddMemberRefOperationOptions {
	return AddMemberRefOperationOptions{}
}

func (o AddMemberRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddMemberRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddMemberRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddMemberRef - Add a member. Use this API to add a member (user, group, or device) to an administrative unit or to
// create a new group within an administrative unit. All group types can be created within an administrative unit. Note:
// Currently, it's only possible to add one member at a time to an administrative unit.`
func (c MemberClient) AddMemberRef(ctx context.Context, id beta.AdministrativeUnitId, input beta.ReferenceCreate, options AddMemberRefOperationOptions) (result AddMemberRefOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/members/$ref", id.ID()),
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
