package grouppolicydefinitionpreviousversiondefinitionpresentation

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateGroupPolicyDefinitionPreviousVersionDefinitionPresentationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdateGroupPolicyDefinitionPreviousVersionDefinitionPresentation - Update the navigation property presentations in
// deviceManagement
func (c GroupPolicyDefinitionPreviousVersionDefinitionPresentationClient) UpdateGroupPolicyDefinitionPreviousVersionDefinitionPresentation(ctx context.Context, id beta.DeviceManagementGroupPolicyDefinitionIdPreviousVersionDefinitionPresentationId, input beta.GroupPolicyPresentation) (result UpdateGroupPolicyDefinitionPreviousVersionDefinitionPresentationOperationResponse, err error) {
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
