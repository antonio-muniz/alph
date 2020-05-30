package memory

import (
	"context"

	"github.com/antonio-muniz/alph/cmd/alph/internal/model/auth"
	"github.com/antonio-muniz/alph/cmd/alph/internal/storage"
)

type database struct {
	subjects map[string]auth.Subject
}

func NewDatabase() storage.Database {
	return &database{
		subjects: make(map[string]auth.Subject),
	}
}

func (d *database) CreateSubject(ctx context.Context, subject auth.Subject) error {
	d.subjects[subject.ID] = subject
	return nil
}

func (d *database) GetSubject(ctx context.Context, subjectID string) (auth.Subject, error) {
	subject, found := d.subjects[subjectID]
	if !found {
		return auth.Subject{}, storage.ErrSubjectNotFound
	}
	return subject, nil
}
