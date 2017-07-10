package nasne

import (
	"fmt"
	"net/url"
	"strconv"
)

type StatusService service

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

func (s *StatusService) HDDListGet() (*HDDList, error) {
	u := url.URL{
		Scheme: s.client.baseURL.Scheme,
		Host:   fmt.Sprintf("%s:64210", s.client.baseURL.Host),
		Path:   "/status/HDDListGet",
	}
	fmt.Println(u.String())
	res, err := s.client.Get(nil, u.String())
	if err != nil {
		return nil, err
	}

	var hddList HDDList
	err = decodeBody(res, &hddList)
	if err != nil {
		return nil, err
	}
	return &hddList, nil
}

func (s *StatusService) HDDInfoGet(id int) (*HDDInfo, error) {
	values := url.Values{}
	values.Set("id", strconv.Itoa(id))
	u := url.URL{
		Scheme:   s.client.baseURL.Scheme,
		Host:     fmt.Sprintf("%s:64210", s.client.baseURL.Host),
		Path:     "/status/HDDInfoGet",
		RawQuery: values.Encode(),
	}
	fmt.Println(u.String())
	res, err := s.client.Get(nil, u.String())
	if err != nil {
		return nil, err
	}

	var hddInfo HDDInfo
	err = decodeBody(res, &hddInfo)
	if err != nil {
		return nil, err
	}
	return &hddInfo, nil
}