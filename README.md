# empty_files_logger
A Go program that reads a specified directory, identifies empty files, and logs their names to an output file (out.txt). The program also calculates the total byte size required to store these names.
Features
Directory Scan
Reads all file names from a specified directory (given as a command-line argument).

Empty File Identification
Identifies and collects the names of files that are empty (size = 0 bytes).

Byte Calculation
Calculates the total bytes needed to store the names of all empty files, including a newline character after each name.

File Logging
Writes all empty file names into out.txt and prints the total bytes required to the console.

Usage
To run the program, provide a directory path as a command-line argument:

bash
Copy code
go run main.go <directory_path>
Example
If the directory contains the files "file1.txt" (empty), "file2.txt" (not empty), and "file3.txt" (empty), the program will output the total byte size for the names of empty files and create out.txt containing:

Copy code
file1.txt
file3.txt
Code Flow
Command-line Argument

Checks if a directory path is provided; exits if not.
Directory Reading

Uses os.ReadDir() to get file information in the specified directory.
Empty File Identification

Loops through each file; if the file size is 0, it collects the file name and updates the total byte size.
Logging and Output

Stores each empty file name in out.txt, printing the total byte size needed for storage.
This program is great for managing file information in any directory, quickly finding and logging all empty files. It’s efficient and simple, making it an easy tool to use for file management.
