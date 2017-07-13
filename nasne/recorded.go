package nasne

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type RecordedService service

type RecordedTitleListArgs struct {
	SearchCriteria      int
	Filter              int
	StartingIndex       int
	RequestedCount      int
	SortCriteria        int
	WithDescriptionLong int
	WithUserData        int
}

func (s *RecordedService) TitleListGet(ctx context.Context, args *RecordedTitleListArgs) (*TitleList, error) {
	values := url.Values{}
	values.Set("searchCriteria", strconv.Itoa(args.SearchCriteria))
	values.Set("filter", strconv.Itoa(args.Filter))
	values.Set("startingIndex", strconv.Itoa(args.StartingIndex))
	values.Set("requestedCount", strconv.Itoa(args.RequestedCount))
	values.Set("sortCriteria", strconv.Itoa(args.SortCriteria))
	values.Set("withDescriptionLong", strconv.Itoa(args.WithDescriptionLong))
	values.Set("withUserData", strconv.Itoa(args.WithUserData))

	u := url.URL{
		Scheme:   s.client.baseURL.Scheme,
		Host:     fmt.Sprintf("%s:64220", s.client.baseURL.Host),
		Path:     "/recorded/titleListGet",
		RawQuery: values.Encode(),
	}
	res, err := s.client.Get(nil, u.String())
	if err != nil {
		return nil, err
	}

	var titleList TitleList
	err = decodeBody(res, &titleList)
	if err != nil {
		return nil, err
	}
	return &titleList, nil
}
