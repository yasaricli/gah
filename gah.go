package gah

type (
	GinAuth struct {
		AuthBackEnd Backend
	}
)

func NewGinAuth(a Backend) *GinAuth {
	return &GinAuth{a}
}
