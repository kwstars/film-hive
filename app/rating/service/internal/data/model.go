package data

// Rating defines the movie rating.
type Rating struct {
	RecordID   string `json:"recordId"`
	RecordType string `json:"recordType"`
	UserID     string `json:"userId"`
	Value      uint32 `json:"value"`
}
