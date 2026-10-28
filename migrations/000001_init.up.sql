CREATE TABLE repos (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner       TEXT NOT NULL,
    name        TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (owner, name)
);

CREATE TABLE jobs (
    id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    repo_id          UUID NOT NULL REFERENCES repos(id) ON DELETE CASCADE,
    status           TEXT NOT NULL DEFAULT 'queued'
                     CHECK (status IN ('queued','scanning','fixing','awaiting_review','verifying','done','failed','cancelled')),
    attempts         INT  NOT NULL DEFAULT 0,
    max_attempts     INT  NOT NULL DEFAULT 3,
    run_after        TIMESTAMPTZ NOT NULL DEFAULT now(),
    locked_by        TEXT,
    heartbeat_at     TIMESTAMPTZ,
    cancel_requested BOOLEAN NOT NULL DEFAULT false,
    last_error       TEXT,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX jobs_claim_idx ON jobs (status, run_after);

CREATE TABLE audit_log (
    id          BIGSERIAL PRIMARY KEY,
    job_id      UUID NOT NULL REFERENCES jobs(id) ON DELETE CASCADE,
    actor       TEXT NOT NULL,
    action      TEXT NOT NULL,
    detail      JSONB NOT NULL DEFAULT '{}',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX audit_log_job_idx ON audit_log (job_id, created_at);
