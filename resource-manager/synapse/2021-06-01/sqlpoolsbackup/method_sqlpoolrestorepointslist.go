package sqlpoolsbackup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolRestorePointsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RestorePoint
}

type SqlPoolRestorePointsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RestorePoint
}

type SqlPoolRestorePointsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SqlPoolRestorePointsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SqlPoolRestorePointsList ...
func (c SqlPoolsBackupClient) SqlPoolRestorePointsList(ctx context.Context, id SqlPoolId) (result SqlPoolRestorePointsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &SqlPoolRestorePointsListCustomPager{},
		Path:       fmt.Sprintf("%s/restorePoints", id.ID()),
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
		Values *[]RestorePoint `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SqlPoolRestorePointsListComplete retrieves all the results into a single object
func (c SqlPoolsBackupClient) SqlPoolRestorePointsListComplete(ctx context.Context, id SqlPoolId) (SqlPoolRestorePointsListCompleteResult, error) {
	return c.SqlPoolRestorePointsListCompleteMatchingPredicate(ctx, id, RestorePointOperationPredicate{})
}

// SqlPoolRestorePointsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlPoolsBackupClient) SqlPoolRestorePointsListCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate RestorePointOperationPredicate) (result SqlPoolRestorePointsListCompleteResult, err error) {
	items := make([]RestorePoint, 0)

	resp, err := c.SqlPoolRestorePointsList(ctx, id)
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

	result = SqlPoolRestorePointsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
