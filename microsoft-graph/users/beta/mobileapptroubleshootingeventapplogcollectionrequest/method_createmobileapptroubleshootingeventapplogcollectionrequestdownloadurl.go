package mobileapptroubleshootingeventapplogcollectionrequest

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

type CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrlOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AppLogCollectionDownloadDetails
}

// CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrl - Invoke action createDownloadUrl
func (c MobileAppTroubleshootingEventAppLogCollectionRequestClient) CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrl(ctx context.Context, id beta.UserIdMobileAppTroubleshootingEventIdAppLogCollectionRequestId) (result CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrlOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/createDownloadUrl", id.ID()),
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

	var model beta.AppLogCollectionDownloadDetails
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
