package memobileapptroubleshootingeventapplogcollectionrequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeMobileAppTroubleshootingEventByIdAppLogCollectionRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AppLogCollectionRequestCollectionResponse
}

type ListMeMobileAppTroubleshootingEventByIdAppLogCollectionRequestsCompleteResult struct {
	Items []models.AppLogCollectionRequestCollectionResponse
}

// ListMeMobileAppTroubleshootingEventByIdAppLogCollectionRequests ...
func (c MeMobileAppTroubleshootingEventAppLogCollectionRequestClient) ListMeMobileAppTroubleshootingEventByIdAppLogCollectionRequests(ctx context.Context, id MeMobileAppTroubleshootingEventId) (result ListMeMobileAppTroubleshootingEventByIdAppLogCollectionRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/appLogCollectionRequests", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]models.AppLogCollectionRequestCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeMobileAppTroubleshootingEventByIdAppLogCollectionRequestsComplete retrieves all the results into a single object
func (c MeMobileAppTroubleshootingEventAppLogCollectionRequestClient) ListMeMobileAppTroubleshootingEventByIdAppLogCollectionRequestsComplete(ctx context.Context, id MeMobileAppTroubleshootingEventId) (ListMeMobileAppTroubleshootingEventByIdAppLogCollectionRequestsCompleteResult, error) {
	return c.ListMeMobileAppTroubleshootingEventByIdAppLogCollectionRequestsCompleteMatchingPredicate(ctx, id, models.AppLogCollectionRequestCollectionResponseOperationPredicate{})
}

// ListMeMobileAppTroubleshootingEventByIdAppLogCollectionRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeMobileAppTroubleshootingEventAppLogCollectionRequestClient) ListMeMobileAppTroubleshootingEventByIdAppLogCollectionRequestsCompleteMatchingPredicate(ctx context.Context, id MeMobileAppTroubleshootingEventId, predicate models.AppLogCollectionRequestCollectionResponseOperationPredicate) (result ListMeMobileAppTroubleshootingEventByIdAppLogCollectionRequestsCompleteResult, err error) {
	items := make([]models.AppLogCollectionRequestCollectionResponse, 0)

	resp, err := c.ListMeMobileAppTroubleshootingEventByIdAppLogCollectionRequests(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListMeMobileAppTroubleshootingEventByIdAppLogCollectionRequestsCompleteResult{
		Items: items,
	}
	return
}
