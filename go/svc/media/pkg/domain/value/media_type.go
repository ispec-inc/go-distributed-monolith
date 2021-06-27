package value

const (
	MediaTypeImage = "images"
)

type MediaType string

func (m MediaType) String() string {
	return string(m)
}
