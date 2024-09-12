package threadpostextension

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateThreadPostExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdateThreadPostExtension - Update openTypeExtension. Update an open extension (openTypeExtension object) on a
// supported resource type. - If a property in the request body matches the name of an existing property in the
// extension, the data in the extension is updated. - Otherwise, that property and its data are added to the extension.
// The data in an extension can be primitive types or arrays of primitive types. The operation behaves differently for
// resources that are directory objects vs other resources. See the table in the Permissions section for the list of
// resources that support open extensions.
func (c ThreadPostExtensionClient) UpdateThreadPostExtension(ctx context.Context, id beta.GroupIdThreadIdPostIdExtensionId, input beta.Extension) (result UpdateThreadPostExtensionOperationResponse, err error) {
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
