package useronlinemeetingrecording

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

type ListUserByIdOnlineMeetingByIdRecordingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CallRecordingCollectionResponse
}

type ListUserByIdOnlineMeetingByIdRecordingsCompleteResult struct {
	Items []models.CallRecordingCollectionResponse
}

// ListUserByIdOnlineMeetingByIdRecordings ...
func (c UserOnlineMeetingRecordingClient) ListUserByIdOnlineMeetingByIdRecordings(ctx context.Context, id UserOnlineMeetingId) (result ListUserByIdOnlineMeetingByIdRecordingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/recordings", id.ID()),
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
		Values *[]models.CallRecordingCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdOnlineMeetingByIdRecordingsComplete retrieves all the results into a single object
func (c UserOnlineMeetingRecordingClient) ListUserByIdOnlineMeetingByIdRecordingsComplete(ctx context.Context, id UserOnlineMeetingId) (ListUserByIdOnlineMeetingByIdRecordingsCompleteResult, error) {
	return c.ListUserByIdOnlineMeetingByIdRecordingsCompleteMatchingPredicate(ctx, id, models.CallRecordingCollectionResponseOperationPredicate{})
}

// ListUserByIdOnlineMeetingByIdRecordingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOnlineMeetingRecordingClient) ListUserByIdOnlineMeetingByIdRecordingsCompleteMatchingPredicate(ctx context.Context, id UserOnlineMeetingId, predicate models.CallRecordingCollectionResponseOperationPredicate) (result ListUserByIdOnlineMeetingByIdRecordingsCompleteResult, err error) {
	items := make([]models.CallRecordingCollectionResponse, 0)

	resp, err := c.ListUserByIdOnlineMeetingByIdRecordings(ctx, id)
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

	result = ListUserByIdOnlineMeetingByIdRecordingsCompleteResult{
		Items: items,
	}
	return
}
