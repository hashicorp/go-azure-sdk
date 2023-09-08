package meonlinemeetingregistrationregistrant

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

type ListMeOnlineMeetingByIdRegistrationRegistrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.MeetingRegistrantBaseCollectionResponse
}

type ListMeOnlineMeetingByIdRegistrationRegistrantsCompleteResult struct {
	Items []models.MeetingRegistrantBaseCollectionResponse
}

// ListMeOnlineMeetingByIdRegistrationRegistrants ...
func (c MeOnlineMeetingRegistrationRegistrantClient) ListMeOnlineMeetingByIdRegistrationRegistrants(ctx context.Context, id MeOnlineMeetingId) (result ListMeOnlineMeetingByIdRegistrationRegistrantsOperationResponse, err error) {
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

// ListMeOnlineMeetingByIdRegistrationRegistrantsComplete retrieves all the results into a single object
func (c MeOnlineMeetingRegistrationRegistrantClient) ListMeOnlineMeetingByIdRegistrationRegistrantsComplete(ctx context.Context, id MeOnlineMeetingId) (ListMeOnlineMeetingByIdRegistrationRegistrantsCompleteResult, error) {
	return c.ListMeOnlineMeetingByIdRegistrationRegistrantsCompleteMatchingPredicate(ctx, id, models.MeetingRegistrantBaseCollectionResponseOperationPredicate{})
}

// ListMeOnlineMeetingByIdRegistrationRegistrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnlineMeetingRegistrationRegistrantClient) ListMeOnlineMeetingByIdRegistrationRegistrantsCompleteMatchingPredicate(ctx context.Context, id MeOnlineMeetingId, predicate models.MeetingRegistrantBaseCollectionResponseOperationPredicate) (result ListMeOnlineMeetingByIdRegistrationRegistrantsCompleteResult, err error) {
	items := make([]models.MeetingRegistrantBaseCollectionResponse, 0)

	resp, err := c.ListMeOnlineMeetingByIdRegistrationRegistrants(ctx, id)
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

	result = ListMeOnlineMeetingByIdRegistrationRegistrantsCompleteResult{
		Items: items,
	}
	return
}
