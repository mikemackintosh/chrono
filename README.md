# chrono
<p align="center">
  <img width="120px" src="https://github.com/mikemackintosh/chrono/raw/main/.github/example.png">
</p>

### Installation
```sh
go install github.com/mikemackintosh/chrono`
```

Then, with `zsh` you can use it by adding the following to your `.zshrc`:

```sh
_preexec_chrono() {
  export START_TIME=$(chrono -m)
}
preexec_functions+=(_preexec_chrono)

_precmd_chrono() {
  CHRONO_DURATION=$(chrono)
  unset START_TIME
}
precmd_functions+=(_precmd_chrono)
```

And then call the `$CHRONO_DURATION` variable in one of your prompts, such as:
```sh
precmd() {
  RPROMPT="%B%F{215}${CHRONO_DURATION}%f%b"
}
```
