# goScheduler
Partial replacement for Windows Task Scheduler

## Next Steps

- Running a shell command doesn't work
- Errors are not trapped by `exec.Command`
- Want to use `slog`
- Want to have some indicator process ran
- Can we trap STDOUT/ERR from an invoked command like WinSW does?

I want to work on scheduling but the above fundamentals need to be understood and ironed out so I build upon them and not duct taping them.