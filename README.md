# Helix
Helix is a command-line tool that issues HTTP requests, then parses, filters, and displays relevant HTTP response headers. It's designed to help developers and network administrators quickly inspect HTTP headers for debugging or analysis.

This was a project to revisit my GoLang skills by building something small to start.

## Features
- Displays all response headers with colored keys by default.
- Supports filtering headers by a comma-separated list.
- Outputs headers in plain text, colored text, or JSON format.
- Displays the HTTP status code of the response.

## Example 
To see an example of the color-coded output, check out the [screenshot in the images folder](https://github.com/gia-lexa/helix/blob/main/images/Helix.png). 
## Installation
### Build from Source
Make sure you have Go installed.

1. Clone the repository:
```
git clone https://github.com/gia-lexa/helix.git
cd helix
```

2. Build the executable:
```
go build -o helix main.go
```

3. Move the executable to a directory in your PATH:
```
sudo mv helix /usr/local/bin/
```

Now you can use helix from anywhere in your terminal.

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

### JSON Output
Output the headers in JSON format:
```
helix -json https://www.example.com
```

### Plain Text Output
Display headers without any ANSI color codes:
```
helix -plain https://www.example.com
```

### Combining Filters and JSON Output
You can combine multiple options, such as filtering and JSON output:
```
helix -filter server,date -json https://www.example.com
```

### Displaying Help Information
To see all available flags and options:
```
helix -help
```

## Examples
### Basic Request
```
helix https://www.example.com
```

### Filter Specific Headers
```
helix -filter cache-control,content-type https://www.example.com
```

### Output in JSON Format
```
helix -json https://www.example.com
```

### Plain Output
```
helix -plain https://www.example.com
```

