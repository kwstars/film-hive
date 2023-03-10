package data

// Movie defines the movie movie.
type Movie struct {
	RecordID   string `json:"recordId"`
	RecordType string `json:"recordType"`
	UserID     string `json:"userId"`
	Value      uint32 `json:"value"`
}
