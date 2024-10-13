# Helix
Helix is a command-line tool that issues HTTP requests, then parses, filters, and surfaces relevant HTTP response headers. 

But the ultimate purpose of this project was to revisit my GoLang skills by building something small to start.


## Features
- Displays all response headers with colored keys by default.
- Supports filtering headers by a comma-separated list.
- Outputs headers in plain text, colored text, or JSON format.
- Displays the HTTP status code of the response.

### Basic Usage
Display all headers for a given URL:
```
helix https://www.example.com
```
### Filtering Specific Headers
Show only specified headers using the -filter flag:
```
helix -filter cache-control,content-type https://www.example.com
```
