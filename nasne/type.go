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
