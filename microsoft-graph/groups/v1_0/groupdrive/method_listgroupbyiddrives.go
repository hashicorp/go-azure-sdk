package groupdrive

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdDrivesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DriveCollectionResponse
}

type ListGroupByIdDrivesCompleteResult struct {
	Items []models.DriveCollectionResponse
}

// ListGroupByIdDrives ...
func (c GroupDriveClient) ListGroupByIdDrives(ctx context.Context, id GroupId) (result ListGroupByIdDrivesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/drives", id.ID()),
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

// ListGroupByIdDrivesComplete retrieves all the results into a single object
func (c GroupDriveClient) ListGroupByIdDrivesComplete(ctx context.Context, id GroupId) (ListGroupByIdDrivesCompleteResult, error) {
	return c.ListGroupByIdDrivesCompleteMatchingPredicate(ctx, id, models.DriveCollectionResponseOperationPredicate{})
}

// ListGroupByIdDrivesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupDriveClient) ListGroupByIdDrivesCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate models.DriveCollectionResponseOperationPredicate) (result ListGroupByIdDrivesCompleteResult, err error) {
	items := make([]models.DriveCollectionResponse, 0)

	resp, err := c.ListGroupByIdDrives(ctx, id)
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

	result = ListGroupByIdDrivesCompleteResult{
		Items: items,
	}
	return
}
