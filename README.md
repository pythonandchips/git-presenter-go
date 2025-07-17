# git-presenter (Go Implementation)

When presenting code live on stage you have a few choices:

* Change code live on stage and risk making a mistake and the code not working.
* Place code in a slide and not be able to run the code live.

Git-presenter solves this problem by giving a presentation style interface for your code.

## Pre-requisites

* Git
* Go 1.19 or later

## Installation

### From Source

```bash
git clone https://github.com/pythonandchips/git-presenter.git
cd git-presenter/go
go build -o bin/git-presenter ./cmd/git-presenter
```

### Build and Install

```bash
go build -o $HOME/bin/git-presenter ./cmd/git-presenter
```

Make sure `$HOME/bin` is in your PATH.

## Usage

* Commit to git as you develop your code.
* When the code is ready use the "git-presenter init" command to initialise
* Once it is initialised you can start the presentation with "git-presenter start"
* Make more commits if need be and use "git-presenter update"
* Then use the following command to navigate the presentation
* next/n: move to next slide
* back/b: move back a slide
* end/e:  move to end of presentation
* start/s: move to start of presentation
* list/l : list slides in presentation
* help/h: display this message

### Command mode

The default for git presenter is interactive mode however if you want to use git presenter from a text editor you can use command mode.

To start a presentation in command mode use "git-presenter start -c"

Once started you run "git-presenter {{command}}" e.g. to move to the next slide run "git-presenter next"

## Development

### Running Tests

```bash
go test ./internal/controller ./internal/presentation ./internal/presenter ./internal/slide
```

### Building

```bash
go build -o bin/git-presenter ./cmd/git-presenter
```

## Architecture

The Go implementation follows clean architecture principles with the following components:

- **GitPresenter** (`internal/presenter/`): Main entry point that coordinates the presentation system
- **Controller** (`internal/controller/`): Handles initialization, starting, and updating presentations
- **Presentation** (`internal/presentation/`): Manages the presentation state and slide navigation
- **Slide** (`internal/slide/`): Represents individual slides (git commits)

## Other resources

There are couple of videos showing git presenter and how to use it:
* [video 1](https://vimeo.com/38949496)
* [video 2](https://vimeo.com/39225144)

## Contributing to git-presenter

* Check out the latest master to make sure the feature hasn't been implemented or the bug hasn't been fixed yet
* Check out the issue tracker to make sure someone already hasn't requested it and/or contributed it
* Fork the project
* Start a feature/bugfix branch
* Commit and push until you are happy with your contribution
* Make sure to add tests for it. This is important so I don't break it in a future version unintentionally.
* Please try not to mess with the go.mod, version, or history. If you want to have your own version, or is otherwise necessary, that is fine, but please isolate to its own commit so I can cherry-pick around it.

### Development Methodology

This implementation follows:
- **Test-Driven Development (TDD)**: All features are implemented with tests first
- **Clean Architecture**: Separated concerns with clear boundaries
- **Tidy First**: Structural and behavioral changes are kept separate

## Contributors

* [Luís Ferreira - Zamith](https://github.com/zamith) on the original ruby version
* Go implementation created using Claude Code

## Copyright

Copyright (c) 2025 Colin Gemmell

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.