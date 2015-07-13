package shortener

type URLDataType uint

const (
	URLDataTypeUndefiend URLDataType = iota
	URLDataTypeURI
	URLDataTypeSubdomain
)

type URLData struct {
	url        string
	urlType    URLDataType
	slug       string
	uniqID     int64
	visitCount int64
}
