# Helix
Helix is a command-line tool for fetching and displaying HTTP response headers from a given URL. It allows you to filter specific headers and display the output in various formats, including plain text and JSON, making it useful for developers, testers, and network engineers.

The purpose of this project was to revisit my GoLang skills by practicing building a small example app.

This is a work in progress, the next step being the addition of tests.

## Features
- Displays all response headers with colored keys by default.
- Supports filtering headers by a comma-separated list.
- Outputs headers in plain text, colored text, or JSON format.
- Displays the HTTP status code of the response.

## Commands
Display all headers for a given URL:
```
helix https://www.example.com
```

Show only specified headers using the -filter flag:
```
helix -filter cache-control,content-type https://www.example.com
```

Output the headers in JSON format:
```
helix -json https://www.example.com
```

Display headers without any ANSI color codes:
```
helix -plain https://www.example.com
```

You can combine multiple options, such as filtering and JSON output:
```
helix -filter server,date -json https://www.example.com
```

To see all available flags and options:
```
helix -help
```

