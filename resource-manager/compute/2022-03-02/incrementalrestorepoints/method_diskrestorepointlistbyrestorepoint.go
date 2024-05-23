package incrementalrestorepoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskRestorePointListByRestorePointOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DiskRestorePoint
}

type DiskRestorePointListByRestorePointCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DiskRestorePoint
}

// DiskRestorePointListByRestorePoint ...
func (c IncrementalRestorePointsClient) DiskRestorePointListByRestorePoint(ctx context.Context, id RestorePointId) (result DiskRestorePointListByRestorePointOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/diskRestorePoints", id.ID()),
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
		Values *[]DiskRestorePoint `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DiskRestorePointListByRestorePointComplete retrieves all the results into a single object
func (c IncrementalRestorePointsClient) DiskRestorePointListByRestorePointComplete(ctx context.Context, id RestorePointId) (DiskRestorePointListByRestorePointCompleteResult, error) {
	return c.DiskRestorePointListByRestorePointCompleteMatchingPredicate(ctx, id, DiskRestorePointOperationPredicate{})
}

// DiskRestorePointListByRestorePointCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IncrementalRestorePointsClient) DiskRestorePointListByRestorePointCompleteMatchingPredicate(ctx context.Context, id RestorePointId, predicate DiskRestorePointOperationPredicate) (result DiskRestorePointListByRestorePointCompleteResult, err error) {
	items := make([]DiskRestorePoint, 0)

	resp, err := c.DiskRestorePointListByRestorePoint(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = DiskRestorePointListByRestorePointCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
