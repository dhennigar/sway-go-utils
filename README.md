# sway-go-utils
Extended functionality for the Sway Wayland compositor written in Go.

## `launch <app_id>`
Focus the given `app_id` if it exists. Else, launch a new instance.
`launch -b` brings `app_id` to the current workspace if it exists elsewhere.

## `new-workspace`
Create a new empty workspace.
`new-workspace -m` moves focused window to an empty workspace.

## `mark-recapture`
Mark a window for easy recall. Subsequent calls to `mark-recapture` refocus that window.
`mark-recapture -n` marks a new window and `mark-recapture -b` brings marked window to current workspace.
Only one window may be marked at a time.
