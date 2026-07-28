# Frequently Asked Questions (FAQ)

### Q1: What CPU threshold is used to define an idle instance?
By default, any instance with less than 5% average CPU utilization is flagged as idle and recommended for termination or downsizing.

### Q2: Does auto-remediation execute live deletions by default?
No! Remediator defaults to `dryRun = true` mode, printing execution preview logs until explicitly toggled to live remediation.

### Q3: How do I run unit tests?
Run `go test -v ./...` in the root directory to execute all Go unit and integration tests.
