# PatchPilot – Build Plan

AI vulnerability-fix agent. Connect a GitHub repo → scan → AI proposes patches in a sandbox → human review → verify → open PR. Every agent action is audited.

Target roles: OpenAI London – Product Engineer (Cyber), Forward Deployed Engineer.

---

## Phase 1 – Setup ✅ (in progress)
- [x] Install Go 1.27
- [x] `go mod init github.com/pojithakarunathilake99/patchpilot`
- [x] Project layout (`cmd/`, `internal/`, `migrations/`, `web/`, `evals/`, `deploy/`)
- [x] Minimal HTTP server with `GET /healthz`
- [x] First commit
- [ ] Create GitHub repo and push
- [ ] CV bullet

## Phase 2 – Go API + Postgres job model
- [ ] Postgres via Docker Compose (local dev)
- [ ] Migrations (`golang-migrate`): `repos`, `jobs`, `findings`, `patches`, `audit_log`
- [ ] Job states: `queued → scanning → fixing → awaiting_review → verifying → done | failed`
- [ ] `internal/store`: DB access with `pgx`
- [ ] REST endpoints: create job, get job, list jobs, cancel job
- [ ] Worker pool: claim jobs with `SELECT … FOR UPDATE SKIP LOCKED`
- [ ] Retries with backoff, per-job timeout (`context.WithTimeout`), cancellation
- [ ] Crash recovery: reclaim jobs with stale heartbeats on startup
- [ ] Structured logging (`log/slog`), graceful shutdown
- [ ] Unit + integration tests
- [ ] CV bullet

## Phase 3 – Scanner integration
- [ ] `Scanner` interface in `internal/scanner`
- [ ] Semgrep adapter (code findings, JSON output)
- [ ] Trivy adapter (dependency CVEs, JSON output)
- [ ] Normalised `Finding` schema (id, tool, rule, severity, file, line range, message, CWE/CVE)
- [ ] Clone repo (shallow) per job; store findings in Postgres
- [ ] Dedupe + severity filtering
- [ ] Tests against a small vulnerable fixture repo
- [ ] CV bullet

## Phase 4 – Sandbox runner
- [ ] `internal/sandbox` using Docker Engine API (Go SDK)
- [ ] Per-job container: no network, read-only repo mount + writable scratch overlay, non-root user, dropped capabilities, CPU/memory/pids limits, seccomp default
- [ ] Hard timeout + forced cleanup
- [ ] Run tests inside sandbox, capture stdout/stderr/exit code
- [ ] Short-lived scoped GitHub token (GitHub App installation token), never mounted in sandbox
- [ ] Per-job audit log entries for every container action
- [ ] CV bullet

## Phase 5 – OpenAI agent + structured output
- [ ] `internal/agent` using OpenAI Responses API
- [ ] Tools: `read_file`, `propose_patch`, `run_tests` (executed in sandbox)
- [ ] Structured JSON output schema: `{finding_id, diff, explanation, confidence, tests_run}`
- [ ] Guardrails: max tool calls, max tokens, path allow-list, diff size limit, reject edits outside finding scope
- [ ] Token + cost telemetry per job
- [ ] Log every tool call/response to `audit_log`
- [ ] Prompt-injection defence: treat repo content as data, never as instructions
- [ ] CV bullet

## Phase 6 – React dashboard
- [ ] Vite + React + TypeScript in `web/`
- [ ] Pages: repos, jobs list (live status), job detail
- [ ] Findings table with severity
- [ ] Diff viewer (side-by-side) with approve / reject
- [ ] Audit log timeline per job
- [ ] On approve: re-scan → verify → open PR via GitHub API
- [ ] CV bullet

## Phase 7 – Evals
- [ ] Python harness in `evals/`
- [ ] Dataset: known-vulnerable repos (OWASP samples, small seeded repos)
- [ ] Metrics: fix rate, verified-fix rate (re-scan clean + tests pass), false-positive rate, time per fix, cost per fix
- [ ] Results table + chart committed to repo
- [ ] Compare at least two prompt/model configs
- [ ] CV bullet

## Phase 8 – Deploy, CI, docs
- [ ] Dockerfiles (multi-stage Go build, static frontend)
- [ ] Docker Compose for full stack
- [ ] Terraform: EC2, security group, IAM role, (optional) RDS
- [ ] GitHub Actions: `go test`, `go vet`, `staticcheck`, frontend build, image build
- [ ] README with architecture diagram
- [ ] Customer onboarding guide (FDE angle)
- [ ] 2-minute demo video
- [ ] Final CV bullets + project summary

---

## Working rules
- One step at a time; confirm before next step.
- Copy-paste commands, no scripts during live work.
- Diagnose errors before fixing.
- One honest CV bullet at the end of each phase.