# Instructions
Install Go:
Ensure you have Go installed on your machine. You can download and install it from the official Go website.

# Save the Code:
Save the above code to a file named cron_parser.go.

# Build the Program:
Open your terminal and navigate to the directory containing cron_parser.go. Then, run the following command to build the program:
go build cron_parser.go

# Run the Program:
Execute the program with a cron expression as an argument:
./cron_parser "*/15 0 1,15 * 1-5 /usr/bin/find"

### Expected Output:
minute         0 15 30 45
hour           0
day of month   1 15
month          1 2 3 4 5 6 7 8 9 10 11 12
day of week    1 2 3 4 5
command        /usr/bin/find


# Cron Parser
A simple Go command line application to parse a cron string and expand each field to show the times at which it will run.

## Usage

```sh
cron_parser "<cron_expression>"

# Building
Ensure you have Go installed.
Clone the repository and navigate to the project directory.
Build the project:
go build cron_parser.go
