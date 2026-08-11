# r3turn

This Golang script is used alongside a file which contains URLs to get the return codes of each URL. I created this for my own sake of VDPs and bug finding, as it's useful to see which URLs from a tool like subfinder are valid and which aren't.

# Installation

`https://github.com/raaz4n/r3turn.git`

# Usage

This program is meant to be used alongside another file which contains URLs.
./r3turn <file> <ms>

If the responses coming through aren't what's expected, you can increase the amount of ms you want the request to stay for. I recommend 200ms personally