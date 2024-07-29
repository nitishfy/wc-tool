# WC Tool 
---

## 1. Overview

WC is a command line tool written in Go that matches the functionality of the Unix Command Line Tool `wc`. It is used to count number of words, bytes, lines, characters that is passed to it, either as a file or Standard Input (Stdin).

## 2. How to use the wc tool?

For testing purposes, there is ![test.txt](/test.txt) file present at the root of the repository that you can use to test the functionalities. 

Here are some examples to use the wc tool:

**Note: If you pass multiple arguments to the tool, only the first argument will be considered.**

- Build the binary using the following command to test the tool against various inputs:
  ```
  go build -o wc main.go
  ```

**Calculating the details of the file input**

- Calculate the number of bytes in a file
  ```
  ./wc -c test.txt
  ```
- Calculate the number of lines in a file
  ```
  ./wc -l test.txt
  ```
- Calculate the number of words in a file
  ```
  ./wc -w test.txt
  ```  
- Calcuate the number of characters in a file
  ```
  ./wc -m test.txt
  ```
- Calculate all the things at once
  ```
  ./wc test.txt
  ```    
**Calculating the details from Standard Input (Stdin)**

- Calculate the number of bytes 
  ```
  echo "nitish" | ./wc -c
  ```
- Calculate the number of lines
  ```
  cat test.txt | ./wc -l
  ```  
- Calculate the number of words
  ```
  cat test.txt | ./wc -w
  ```
- Calculate the number of characters
  ```
  cat test.txt | ./wc -m
  ```

## 3. How to contribute?

Thank you for considering contributing to this project! We welcome your contributions and support in making this project better.

Contributions are welcome in the form of bug reports, feature requests, code changes, documentation updates, and more. To get started, please follow these steps:

1. **Create an Issue:**

   - If you find a bug or have a feature request, please create an issue.
   - Provide as much detail as possible, including the version of the project where you encountered the issue, your operating system, and steps to reproduce the problem.

2. **Fork the Repository:**

   - If you plan to make code contributions, fork the repository to your GitHub account.

3. **Make Changes:**

   - Create a new branch for your changes: `git checkout -b feature/your-feature` or `fix/your-fix`.
   - Make your changes, ensuring that your code follows the project's coding guidelines and standards.

4. **Test Your Changes:**

   - Ensure that your changes work as expected and do not introduce new issues.

5. **Submit a Pull Request:**

   - When you're ready to submit your changes, create a pull request from your fork to the main repository's `master` branch (or the relevant target branch).
   - In the pull request, provide a clear description of your changes and reference any related issues.

6. **Code Review:**

   - Your pull request will be reviewed by project maintainers. Please be responsive to any feedback or requested changes.

7. **Merge and Release:**
   - Once your changes are approved, they will be merged into the project. 
