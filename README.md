# Sleigh

![icon of a sleigh](./assets/sleigh.png "Sleigh")

Fetches advent of code puzzle input and saves it to a `<day>-input.txt` file so you can start
cooking before Mrs. Claus gets home.

### Releases

[Here's the release page](https://github.com/jacktrusler/sleigh/releases), pick your OS and
architechture for the raw binary/executable. 

### Building from source

Clone the repository then build the binary by going into the `main.go` directory `cd cmd/sleigh` and
run `go build .` 

You can run directly with `./sleigh`, copy the binary to your system binaries, or create a path in
your `.bashrc` | `.zshrc` as follows: 

`export PATH="$HOME/path/to/sleigh"`  
or  
`export PATH="/full/system/path/to/sleigh"`  

# Usage

You must provide a session cookie to get the input. Advent of Code needs a way to verify that you've
logged into the website. To make it easy, its recommended to create an `.env` file with
`SESSION_COOKIE=<your-session-cookie>` and run `sleigh` from that directory. Other options are:

- add to your environment directly: `export SESSION_COOKIE=<your-session-cookie>`  
- pass it using the -c flag `sleigh -c <your-session-cookie>`  

To see all options run `sleigh -h`

### Finding your Session Cookie

1. Log in to Advent of Code however you want, AoC will provide a Session Cookie.

2. Open dev tools and navigate to the "Application" tab, and find Cookies under "Storage".

3. Copy the Value of the cookie with the Name `session`.
