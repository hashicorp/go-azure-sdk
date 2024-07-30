package b2xuserflowuserattributeassignment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetIdentityB2xUserFlowUserAttributeAssignmentsOrderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// SetIdentityB2xUserFlowUserAttributeAssignmentsOrder ...
func (c B2xUserFlowUserAttributeAssignmentClient) SetIdentityB2xUserFlowUserAttributeAssignmentsOrder(ctx context.Context, id IdentityB2xUserFlowId, input SetIdentityB2xUserFlowUserAttributeAssignmentsOrderRequest) (result SetIdentityB2xUserFlowUserAttributeAssignmentsOrderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/userAttributeAssignments/setOrder", id.ID()),
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
