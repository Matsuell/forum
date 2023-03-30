package forum

type Topic struct {
	TopicID       int
	TopicTitle    string
	TopicMessage  string
	TopicTime     string
	TopicAuthor   string
	TopicCategory string
	Likes         int
	Dislikes      int
	Comments      []Comment
}
