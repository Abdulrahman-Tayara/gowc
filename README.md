# Go WC (Word Count) Tool

Go WC is a simple command-line tool written in Go for counting words, lines, and characters in a given text file.

## Installation

Make sure you have Go installed on your machine. Then, you can install the Go WC tool using the following command:

```bash
go get -u github.com/Abdulrahman-Tayara/gowc
```
This will download and install the gowc binary into your Go bin directory.

Usage
To use Go WC, run the following command:

```bash
gowc <filename>
```
Replace <filename> with the path to the text file you want to analyze.

Options
- -l: Count lines in the file.
- -w: Count words in the file.
- -c: Count bytes in the file.
- -m: Count characters in the file.
You can combine options to count multiple metrics at once. For example:

```bash
gowc -l -w myfile.txt
```
This will count both lines and words in the myfile.txt file.

Examples
```bash
# Count lines, words, and characters in myfile.txt
gowc myfile.txt

# Count only lines in myfile.txt
gowc -l myfile.txt

# Count words and characters in myfile.txt
gowc -w -c myfile.txt
```

# License
This Go WC tool is open-source and available under the MIT License. Feel free to contribute and improve it!


## Authors

- [@Abdulrahman-Tayara](https://github.com/Abdulrahman-Tayara)

# Issues and Contributions
If you encounter any issues or would like to contribute to the project, please open an issue or create a pull request.

