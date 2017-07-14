# go-nasne
nasne API Client for golang.

## Usage ###

```go
import "github.com/ayumu83s/go-nasne/nasne"

apiClient, err := nasne.NewClient("192.168.11.3", nil)
```

### HDDの一覧取得 ###
/status/HDDListGet
```go
hddList, err := apiClient.Status.HDDListGet(nil)
```

### HDDの詳細取得 ###
/status/HDDInfoGet
```go
hddInfo, err := apiClient.Status.HDDInfoGet(nil, hddList[0].ID)
```
* hddInfo.HDD.UsedVolumeSize: 使用量
* hddInfo.HDD.FreeVolumeSize: 空き容量
* hddInfo.HDD.TotalVolumeSize: 総量

### 録画失敗の一覧取得 ###
/status/RecNgListGet
```go
recNgList, err := apiClient.Status.RecNgListGet(nil)
```
* recNgList.Number: 録画失敗件数


### 録画済みタイトル一覧取得 ###
/recorded/TitleListGet
```go
args := &nasne.RecordedTitleListArgs{
	RequestedCount: 1, // 取得したい件数
}
titleList, err := apiClient.Recorded.TitleListGet(nil, args)
```
* titleList.TotalMatches: 録画件数
