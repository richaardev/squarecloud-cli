package api

import (
	"errors"
	"net/http"
	"os"
)

type Application struct {
	ID       string `json:"id"`
	Tag      string `json:"tag"`
	RAM      int    `json:"ram"`
	Lang     string `json:"lang"`
	Type     string `json:"type"`
	Cluster  string `json:"cluster"`
	IsWesite bool   `json:"isWebsite"`
	Avatar   string `json:"avatar"`
}

type ResponseApplicationStatus struct {
	CPU     string `json:"cpu"`
	RAM     string `json:"ram"`
	Status  string `json:"status"`
	Running bool   `json:"running"`
	Storage string `json:"storage"`
	Network struct {
		Total string `json:"total"`
		Now   string `json:"now"`
	} `json:"network"`
	Requests int   `json:"requests"`
	Uptime   int64 `json:"uptime"`
	Time     int64 `json:"time"`
}

type ResponseApplicationBackup struct {
	DownloadURL string `json:"downloadURL"`
}

type ResponseApplicationLogs struct {
	Logs string `json:"logs"`
}

func ApplicationStatus(appId string) (*ResponseApplicationStatus, error) {
	res, err := request[ResponseApplicationStatus](http.MethodGet, EndpointApplicationStatus(appId), nil)
	if err != nil {
		return nil, err
	}
	if res.RawResponse.StatusCode != 200 {
		return nil, errors.New(ParseError(res))
	}

	return &res.Response, nil
}

func ApplicationBackup(appId string) (*ResponseApplicationBackup, error) {
	res, err := request[ResponseApplicationBackup](http.MethodGet, EndpointApplicationBackup(appId), nil)
	if err != nil {
		return nil, err
	}
	if res.RawResponse.StatusCode != 200 {
		return nil, errors.New(ParseError(res))
	}

	return &res.Response, nil
}

func ApplicationLogs(appId string, full bool) (*ResponseApplicationLogs, error) {
	res, err := request[ResponseApplicationLogs](http.MethodGet, EndpointApplicationLogs(appId, full), nil)
	if err != nil {
		return nil, err
	}
	if res.RawResponse.StatusCode != 200 {
		return nil, errors.New(ParseError(res))
	}

	return &res.Response, nil
}

func ApplicationStart(appId string) (bool, error) {
	res, err := request[any](http.MethodPost, EndpointApplicationStart(appId), nil)
	if err != nil {
		return false, err
	}
	if res.RawResponse.StatusCode != 200 {
		return false, errors.New(ParseError(res))
	}

	return res.Code == "ACTION_SENT", err
}

func ApplicationStop(appId string) (bool, error) {
	res, err := request[any](http.MethodPost, EndpointApplicationStop(appId), nil)
	if err != nil {
		return false, err
	}
	if res.RawResponse.StatusCode != 200 {
		return false, errors.New(ParseError(res))
	}

	return res.Code == "ACTION_SENT", err
}

func ApplicationRestart(appId string) (bool, error) {
	res, err := request[any](http.MethodPost, EndpointApplicationRestart(appId), nil)
	if err != nil {
		return false, err
	}
	if res.RawResponse.StatusCode != 200 {
		return false, errors.New(ParseError(res))
	}

	return res.Code == "ACTION_SENT", nil
}

func ApplicationUpload(file *os.File) (bool, error) {
	res, err := sendFile(EndpointApplicationUpload, file.Name())
	if err != nil {
		return false, err
	}
	t, _ := decodeResponse[any](res)
	if res.StatusCode != 200 {
		return false, errors.New(ParseError(t))
	}

	return t.Code == "SUCCESS", nil
}

func ApplicationCommit(appId string, file *os.File) (bool, error) {
	res, err := sendFile(EndpointApplicationCommit(appId), file.Name())
	if err != nil {
		return false, err
	}
	t, _ := decodeResponse[any](res)
	if res.StatusCode != 200 {
		return false, errors.New(ParseError(t))
	}

	return t.Code == "SUCCESS", nil
}

func ApplicationDelete(appId string) (bool, error) {
	res, err := request[any](http.MethodPost, EndpointApplicationDelete(appId), nil)
	if err != nil {
		return false, err
	}
	if res.RawResponse.StatusCode != 200 {
		return false, errors.New(ParseError(res))
	}

	return res.Code == "APP_DELETED", nil
}
