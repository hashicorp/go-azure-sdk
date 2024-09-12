package member

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

type RemoveMemberRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveMemberRefsOperationOptions struct {
	Id      *string
	IfMatch *string
}

func DefaultRemoveMemberRefsOperationOptions() RemoveMemberRefsOperationOptions {
	return RemoveMemberRefsOperationOptions{}
}

func (o RemoveMemberRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveMemberRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RemoveMemberRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Id != nil {
		out.Append("@id", fmt.Sprintf("%v", *o.Id))
	}
	return &out
}

// RemoveMemberRefs - Remove directory role member. Remove a member from a directoryRole. You can use both the object ID
// and template ID of the directoryRole with this API. The template ID of a built-in role is immutable and can be seen
// in the role description on the Microsoft Entra admin center. For details, see Role template IDs.
func (c MemberClient) RemoveMemberRefs(ctx context.Context, id stable.DirectoryRoleId, options RemoveMemberRefsOperationOptions) (result RemoveMemberRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/members/$ref", id.ID()),
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
