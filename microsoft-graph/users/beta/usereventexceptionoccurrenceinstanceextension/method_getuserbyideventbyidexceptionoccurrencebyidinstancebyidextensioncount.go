package usereventexceptionoccurrenceinstanceextension

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdExtensionCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// GetUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdExtensionCount ...
func (c UserEventExceptionOccurrenceInstanceExtensionClient) GetUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdExtensionCount(ctx context.Context, id UserEventExceptionOccurrenceInstanceId) (result GetUserByIdEventByIdExceptionOccurrenceByIdInstanceByIdExtensionCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/extensions/$count", id.ID()),
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
