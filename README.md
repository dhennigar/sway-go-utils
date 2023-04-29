# sway-go-utils
Extended functionality for the Sway Wayland compositor written in Go.

## `launch`
focuses the given program if it exists. Else, launches a new instance.
`launch -b` brings the target program to the current workspace

## `new-workspace`
creates a new free workspace.
`new-workspace -m` moves currently focused window to a free workspace.

## `mark-recapture`
marks window for easy access. subsequent calls focus that window.
`mark-recapture -n` marks a new window and `mark-recapture -b` brings marked window to current workspace.
