package songusecase

import "github.com/cockroachdb/errors/domains"

var (
	SongNotFoundMark = domains.New("song_not_found")
	DefaultErrorMark = domains.New("default_error")
)
