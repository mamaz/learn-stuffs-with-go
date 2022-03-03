package tokenbucket

type Bucket struct {
	TokenNumber int
	RefillUnit  string
	Cache       interface{}
}

func (b Bucket) HandleRequest(payload interface{}) {

}
