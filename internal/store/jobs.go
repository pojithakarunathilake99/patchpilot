package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/pojithakarunathilake99/patchpilot/internal/jobs"
)

var ErrNotFound = errors.New("not found")

const jobColumns = `j.id::text, r.owner, r.name, j.status, j.attempts, j.max_attempts,
	j.last_error, j.created_at, j.updated_at`

func scanJob(row pgx.Row) (jobs.Job, error) {
	var j jobs.Job
	err := row.Scan(&j.ID, &j.RepoOwner, &j.RepoName, &j.Status, &j.Attempts,
		&j.MaxAttempts, &j.LastError, &j.CreatedAt, &j.UpdatedAt)
	return j, err
}

func (s *Store) CreateJob(ctx context.Context, owner, name string) (jobs.Job, error) {
	var id string
	err := s.pool.QueryRow(ctx, `
		WITH r AS (
			INSERT INTO repos (owner, name) VALUES ($1, $2)
			ON CONFLICT (owner, name) DO UPDATE SET owner = EXCLUDED.owner
			RETURNING id
		)
		INSERT INTO jobs (repo_id) SELECT id FROM r
		RETURNING id::text`, owner, name).Scan(&id)
	if err != nil {
		return jobs.Job{}, fmt.Errorf("insert job: %w", err)
	}
	return s.GetJob(ctx, id)
}

func (s *Store) GetJob(ctx context.Context, id string) (jobs.Job, error) {
	j, err := scanJob(s.pool.QueryRow(ctx,
		`SELECT `+jobColumns+` FROM jobs j JOIN repos r ON r.id = j.repo_id WHERE j.id = $1`, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return jobs.Job{}, ErrNotFound
	}
	if err != nil {
		return jobs.Job{}, fmt.Errorf("get job: %w", err)
	}
	return j, nil
}

func (s *Store) ListJobs(ctx context.Context, limit int) ([]jobs.Job, error) {
	rows, err := s.pool.Query(ctx,
		`SELECT `+jobColumns+` FROM jobs j JOIN repos r ON r.id = j.repo_id
		 ORDER BY j.created_at DESC LIMIT $1`, limit)
	if err != nil {
		return nil, fmt.Errorf("list jobs: %w", err)
	}
	defer rows.Close()

	out := []jobs.Job{}
	for rows.Next() {
		j, err := scanJob(rows)
		if err != nil {
			return nil, fmt.Errorf("scan job: %w", err)
		}
		out = append(out, j)
	}
	return out, rows.Err()
}
