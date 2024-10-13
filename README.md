# Helix
Helix is a command-line tool that issues HTTP requests, then parses, filters, and surfaces relevant HTTP response headers. 

The purpose of this project was to revisit my GoLang skills by practicing building a small example app.

## Examples

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
