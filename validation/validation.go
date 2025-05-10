package validation

type CreateUrlRequest struct {
	OriginalUrl string `json:"original_url" validate:"required,url"`
}
