package nasne

type Hdd struct {
	ID              int
	InternalFlag    int
	MountStatus     int
	RegisterFlag    int
	Format          string
	Name            string
	VendorID        string
	ProductID       string
	SerialNumber    string
	UsedVolumeSize  int
	FreeVolumeSize  int
	TotalVolumeSize int
}

type HDDList struct {
	Errorcode int
	Number    int
	Hdd       []Hdd
}

type HDDInfo struct {
	Errorcode int
	Hdd       Hdd
}

type Genre struct {
	ID   int
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
	ID               string
	Title            string
	Description      string
	StartDateTime    string // 日時
	Duration         int
	ConditionID      string
	Quality          int
	ChannelName      string
	ChannelNumber    int
	BroadcastingType int
	EerviceID        int
	EventID          int
	Genre            []Genre
	AudioInfo        []Audio
	CaptionInfo      int
	ComponentType    int
	ProtectFlag      int
	NewFlag          int
	PlayCount        int
	CreatorID        int
	StorageID        int
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

type DesiredMatchingID struct {
	NetworkID         int
	TransportStreamID int
	ServiceID         int
	EventID           int
}

type RecNgItem struct {
	ID                     string
	ContentID              string
	Title                  string
	ScheduledStartDateTime string
	ScheduledDuration      int
	ScheduledConditionID   string
	ScheduledChannelID     int
	BroadcastingType       int
	DesiredMatchingID      DesiredMatchingID
	DesiredQualityMode     int
	GenreIDNum             int
	GenreID                []Genre
	ReservationCreatorID   int
	RecordDestinationID    int
	NgID                   int
}

type RecNgList struct {
	Errorcode int
	Number    int
	Item      []RecNgItem
}
