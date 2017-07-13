package nasne

import (
	"fmt"
	"net/url"
	"strconv"
)

type RecordedService service

type Genre struct {
	Id   int
	Type int
}

type Audio struct {
	ComponentTag  int
	ComponentType int
}

type ContainerSize struct {
	Main   int
	Mobile int
	Thumb  int
}

// Title is response
type Title struct {
	Id               string
	Title            string
	Description      string
	StartDateTime    string // 日時
	Duration         int
	ConditionId      string
	Quality          int
	ChannelName      string
	ChannelNumber    int
	BroadcastingType int
	EerviceId        int
	EventId          int
	Genre            []Genre
	AudioInfo        []Audio
	CaptionInfo      int
	ComponentType    int
	ProtectFlag      int
	NewFlag          int
	PlayCount        int
	CreatorId        int
	StorageId        int
	RecordingFlag    int
	CopyControl      int
	CopyCount        int
	ParentalRating   int
	ResumePosition   int
	ContainerSize    ContainerSize
	BitrateMobile    int
}

type TitleList struct {
	Errorcode      int
	Item           []Title
	TotalMatches   int
	NumberReturned int
}

type RecordedTitleListArgs struct {
	SearchCriteria      int
	Filter              int
	StartingIndex       int
	RequestedCount      int
	SortCriteria        int
	WithDescriptionLong int
	WithUserData        int
}

func (s *RecordedService) TitleListGet(args *RecordedTitleListArgs) (*TitleList, error) {
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
	fmt.Println(u.String())
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
