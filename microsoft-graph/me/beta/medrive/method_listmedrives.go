package medrive

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

type ListMeDrivesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DriveCollectionResponse
}

type ListMeDrivesCompleteResult struct {
	Items []models.DriveCollectionResponse
}

// ListMeDrives ...
func (c MeDriveClient) ListMeDrives(ctx context.Context) (result ListMeDrivesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/drives",
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
		Values *[]models.DriveCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeDrivesComplete retrieves all the results into a single object
func (c MeDriveClient) ListMeDrivesComplete(ctx context.Context) (ListMeDrivesCompleteResult, error) {
	return c.ListMeDrivesCompleteMatchingPredicate(ctx, models.DriveCollectionResponseOperationPredicate{})
}

// ListMeDrivesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeDriveClient) ListMeDrivesCompleteMatchingPredicate(ctx context.Context, predicate models.DriveCollectionResponseOperationPredicate) (result ListMeDrivesCompleteResult, err error) {
	items := make([]models.DriveCollectionResponse, 0)

	resp, err := c.ListMeDrives(ctx)
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

	result = ListMeDrivesCompleteResult{
		Items: items,
	}
	return
}
