package useronlinemeetingregistrationregistrant

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

type ListUserByIdOnlineMeetingByIdRegistrationRegistrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MeetingRegistrantBaseCollectionResponse
}

type ListUserByIdOnlineMeetingByIdRegistrationRegistrantsCompleteResult struct {
	Items []models.MeetingRegistrantBaseCollectionResponse
}

// ListUserByIdOnlineMeetingByIdRegistrationRegistrants ...
func (c UserOnlineMeetingRegistrationRegistrantClient) ListUserByIdOnlineMeetingByIdRegistrationRegistrants(ctx context.Context, id UserOnlineMeetingId) (result ListUserByIdOnlineMeetingByIdRegistrationRegistrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/registration/registrants", id.ID()),
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
		Values *[]models.MeetingRegistrantBaseCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdOnlineMeetingByIdRegistrationRegistrantsComplete retrieves all the results into a single object
func (c UserOnlineMeetingRegistrationRegistrantClient) ListUserByIdOnlineMeetingByIdRegistrationRegistrantsComplete(ctx context.Context, id UserOnlineMeetingId) (ListUserByIdOnlineMeetingByIdRegistrationRegistrantsCompleteResult, error) {
	return c.ListUserByIdOnlineMeetingByIdRegistrationRegistrantsCompleteMatchingPredicate(ctx, id, models.MeetingRegistrantBaseCollectionResponseOperationPredicate{})
}

// ListUserByIdOnlineMeetingByIdRegistrationRegistrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserOnlineMeetingRegistrationRegistrantClient) ListUserByIdOnlineMeetingByIdRegistrationRegistrantsCompleteMatchingPredicate(ctx context.Context, id UserOnlineMeetingId, predicate models.MeetingRegistrantBaseCollectionResponseOperationPredicate) (result ListUserByIdOnlineMeetingByIdRegistrationRegistrantsCompleteResult, err error) {
	items := make([]models.MeetingRegistrantBaseCollectionResponse, 0)

	resp, err := c.ListUserByIdOnlineMeetingByIdRegistrationRegistrants(ctx, id)
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

	result = ListUserByIdOnlineMeetingByIdRegistrationRegistrantsCompleteResult{
		Items: items,
	}
	return
}
