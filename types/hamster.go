package types

type Hamster struct {
	Id          uint64   `json:"id"`
	ImgUrl      string   `json:"imgUrl"`
	Content     string   `json:"content"`
	Visibility  string   `json:"visibility"`
	LikesCount  uint64   `json:"likes"`
	SharesCount uint64   `json:"shares"`
	Comments    []string `json:"comments"`
	AuthorID    uint64   `json:"authorId"`
	CreatedAt   string   `json:"createdAt"`
}
