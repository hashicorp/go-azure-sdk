package group

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdateGroup - Upsert group. Create a new group object if it doesn't exist, or update the properties of an existing
// group object. You can create or update the following types of group: By default, this operation returns only a subset
// of the properties for each group. For a list of properties that are returned by default, see the Properties section
// of the group resource. To get properties that are not returned by default, do a GET operation and specify the
// properties in a $select OData query option.
func (c GroupClient) UpdateGroup(ctx context.Context, id beta.GroupId, input beta.Group) (result UpdateGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPatch,
		Path:       id.ID(),
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
