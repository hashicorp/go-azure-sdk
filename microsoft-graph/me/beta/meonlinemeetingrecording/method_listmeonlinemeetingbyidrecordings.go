package meonlinemeetingrecording

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

type ListMeOnlineMeetingByIdRecordingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.CallRecordingCollectionResponse
}

type ListMeOnlineMeetingByIdRecordingsCompleteResult struct {
	Items []models.CallRecordingCollectionResponse
}

// ListMeOnlineMeetingByIdRecordings ...
func (c MeOnlineMeetingRecordingClient) ListMeOnlineMeetingByIdRecordings(ctx context.Context, id MeOnlineMeetingId) (result ListMeOnlineMeetingByIdRecordingsOperationResponse, err error) {
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

// ListMeOnlineMeetingByIdRecordingsComplete retrieves all the results into a single object
func (c MeOnlineMeetingRecordingClient) ListMeOnlineMeetingByIdRecordingsComplete(ctx context.Context, id MeOnlineMeetingId) (ListMeOnlineMeetingByIdRecordingsCompleteResult, error) {
	return c.ListMeOnlineMeetingByIdRecordingsCompleteMatchingPredicate(ctx, id, models.CallRecordingCollectionResponseOperationPredicate{})
}

// ListMeOnlineMeetingByIdRecordingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnlineMeetingRecordingClient) ListMeOnlineMeetingByIdRecordingsCompleteMatchingPredicate(ctx context.Context, id MeOnlineMeetingId, predicate models.CallRecordingCollectionResponseOperationPredicate) (result ListMeOnlineMeetingByIdRecordingsCompleteResult, err error) {
	items := make([]models.CallRecordingCollectionResponse, 0)

	resp, err := c.ListMeOnlineMeetingByIdRecordings(ctx, id)
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

	result = ListMeOnlineMeetingByIdRecordingsCompleteResult{
		Items: items,
	}
	return
}
