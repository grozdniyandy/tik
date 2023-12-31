# Tik (Тык)

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Tik (Тык) is a command-line tool written in Go for concurrently checking security headers (Content-Security-Policy and X-Frame-Options) of multiple domains from a list. This tool features a custom text display, a progress bar, and the ability to specify the number of concurrent threads for checking. The purpose of this tool is to identify clickjacking vulnerabilities, especially in a big scope.

## Table of Contents
- [Features](https://github.com/grozdniyandy/tik#features)
- [Usage](https://github.com/grozdniyandy/tik#usage)
- [Installation](https://github.com/grozdniyandy/tik#installation)
- [Example Input File](https://github.com/grozdniyandy/tik#example-input-file)
- [Dependencies](https://github.com/grozdniyandy/tik#dependencies)
- [License](https://github.com/grozdniyandy/tik#license)
- [Author](https://github.com/grozdniyandy/tik#author)
- [Contributing](https://github.com/grozdniyandy/tik#contributing)

## Features
- Concurrent checking of security headers for multiple domains.
- Progress bar to track the status of domain checking.

## Usage
1. **Clone or Download:** Clone this repository or download the code to your local machine.
2. **Run the tool:** Run the tool using the following command:
   ```
   go run main.go -f <filename> -t <thread-count>
   ```

## Installation
You can either check the "Usage" or download already compiled code from "releases".

## Example Input File
The input file should contain a list of domains, one per line, like this:
```
example.com
example2.com
example3.com
example4.com
```

## Dependencies
This code uses the Go standard library, so there are no external dependencies to install.

## License
This code is released under the [MIT License](LICENSE).

## Author
Tik is developed by GrozdniyAndy of [XSS.is](https://xss.is).

## Contributing
Feel free to contribute, report issues, or suggest improvements by creating pull requests or issues in the GitHub repository. Enjoy using this simple clickjacking checker!
