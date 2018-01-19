## Go Examples

A list of useful go snippets demonstrating a functionality, concept and scripts for common software development use cases in [go](https://golang.org/). All the code snippets are functional and a good reference to get started. [gb tool](https://getgb.io/) is used to structure the project.

Snippet list is not yet comprehensive and work in is progress. If you have any code snippet not present currently in the list, feel free to contribute. Follow the guidelines to add code snippet.


#### Project structure
The project is structured as shown below:

├── README.md

├── bin

│   ├── helloworld

├── pkg

│   ├── darwin-amd64

│   │   └── helloworld.a

├── src

│   ├── helloworld

│   │   └── main.go


 - `src/` contains source files
 - `bin/` contains binary executables
 - `pkg`  contains package objects


#### Usage

- Build a snippet: `gb build <mysnippet>`
- Run a snippet: `./bin/mysnippet`

#### Available snippets
- **helloworld**: "Hello world", the first program.
- **file-read-write**: Performs file read and write operations.
- **httpserver**: A simplest http server in go, Use of `HandlFunc`.
- **httphandler**: A http server, handler implemented by a user defined type. Use of `Handle` and `ServeHTTP`.
- **kafka-producer**:  Script to produce events to kafka topic.
- **kafka-consumer**: Script to consume events from kafka topic.
- **http-fileserver**: A static file server using built in handler `FileServer`. Usage of `StripPrefix` is also shown.


#### How to add a snippet
- Create a directory inside `src/` representing name of snippet i.e helloworld
- Add source code in `src/helloworld/main.go`
- Build source code:  `gb build helloworld`
- Test source code: `./bin/hellworld`
