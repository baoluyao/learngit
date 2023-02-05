package model

type EventData struct {
	Events []*Event `json:"events"`
}

type Event struct {
	Code         int64   `json:"code"`
	Type         int64   `json:"type"`
	StartTime    int64   `json:"startTime"`
	StartTime_tz int64   `json:"startTime_tz"`
	Level        int64   `json:"level"`
	Paras        []*Para `json:"paras"`
}

type Para struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
