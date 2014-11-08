# Prompt

[GoDoc](http://godoc.org/github.com/Bowery/prompt)

Prompt is a cross platform line-editing prompting library. Read the GoDoc page
for more info and for API details.

## Features
- Keyboard shortcuts in prompts
- Secure password prompt
- Custom prompt support
- Fallback prompt for unsupported terminals
- ANSI conversion for Windows

## Todo
- Make refresh less jittery on Windows
- Add support for BSD systems
- Multi-byte character support on Windows
- `AnsiWriter` should execute the equivalent ANSI escape code functionality on Windows
- Support for more ANSI escape codes on Windows.
- More keyboard shortcuts from Readlines shortcut list

## Contributing

Make sure Go is setup and running the latest release version, and make sure your `GOPATH` is setup properly.

Run `git clone https://github.com/Bowery/prompt.git ${GOPATH}/src/github.com/Bowery/prompt`

Follow the guidelines [here](https://guides.github.com/activities/contributing-to-open-source/#contributing).

Please be sure to `gofmt` any code before doing commits. You can simply run `gofmt -w .` to format all the code in the directory.

## License

Prompt is MIT licensed, details can be found [here](https://raw.githubusercontent.com/Bowery/prompt/master/LICENSE).
