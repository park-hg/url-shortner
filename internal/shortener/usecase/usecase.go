package usecase

import (
	"context"
)

type urlMapper interface {
	Shorten(ctx context.Context, original string) (shortened string, err error)
	RetrieveOriginal(ctx context.Context, shortened string) (original string, err error)
}

type ShortenURLUseCase struct {
	repo urlMapper
}

func NewShortenURLUseCase(repo urlMapper) *ShortenURLUseCase {
	return &ShortenURLUseCase{repo: repo}
}

func (uc *ShortenURLUseCase) Shorten(ctx context.Context, original string) (string, error) {
	return uc.repo.Shorten(ctx, original)
}

func (uc *ShortenURLUseCase) RetrieveOriginal(ctx context.Context, shortened string) (string, error) {
	return uc.repo.RetrieveOriginal(ctx, shortened)
}
