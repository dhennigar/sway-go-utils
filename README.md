# sway-go-utils
Extended functionality for the Sway Wayland compositor written in Go.

## sg-launch
focuses the given program if it exists. Else, launches a new instance.
### options
`-b` brings the target program to the current workspace

## sg-free
creates a new free workspace.
### options
`-m` moves currently focused window to a free workspace.

## sg-mark
marks window for easy access. subsequent calls focus that window.
### options
`-n` marks a new window.
`-b` brings marked window to current workspace.
