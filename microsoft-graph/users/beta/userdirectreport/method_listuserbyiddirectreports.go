package userdirectreport

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

type ListUserByIdDirectReportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DirectoryObjectCollectionResponse
}

type ListUserByIdDirectReportsCompleteResult struct {
	Items []models.DirectoryObjectCollectionResponse
}

// ListUserByIdDirectReports ...
func (c UserDirectReportClient) ListUserByIdDirectReports(ctx context.Context, id UserId) (result ListUserByIdDirectReportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/directReports", id.ID()),
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
		Values *[]models.DirectoryObjectCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdDirectReportsComplete retrieves all the results into a single object
func (c UserDirectReportClient) ListUserByIdDirectReportsComplete(ctx context.Context, id UserId) (ListUserByIdDirectReportsCompleteResult, error) {
	return c.ListUserByIdDirectReportsCompleteMatchingPredicate(ctx, id, models.DirectoryObjectCollectionResponseOperationPredicate{})
}

// ListUserByIdDirectReportsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserDirectReportClient) ListUserByIdDirectReportsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.DirectoryObjectCollectionResponseOperationPredicate) (result ListUserByIdDirectReportsCompleteResult, err error) {
	items := make([]models.DirectoryObjectCollectionResponse, 0)

	resp, err := c.ListUserByIdDirectReports(ctx, id)
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

	result = ListUserByIdDirectReportsCompleteResult{
		Items: items,
	}
	return
}
