package jobs

import "time"

type Status string

const (
	StatusQueued         Status = "queued"
	StatusScanning       Status = "scanning"
	StatusFixing         Status = "fixing"
	StatusAwaitingReview Status = "awaiting_review"
	StatusVerifying      Status = "verifying"
	StatusDone           Status = "done"
	StatusFailed         Status = "failed"
	StatusCancelled      Status = "cancelled"
)

type Job struct {
	ID          string    `json:"id"`
	RepoOwner   string    `json:"repo_owner"`
	RepoName    string    `json:"repo_name"`
	Status      Status    `json:"status"`
	Attempts    int       `json:"attempts"`
	MaxAttempts int       `json:"max_attempts"`
	LastError   *string   `json:"last_error,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
