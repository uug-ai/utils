# uug-ai/utils

## Overview
This repository, `uug-ai/utils`, is a collection of utility functions designed to perform various common tasks in software development. These utilities include string manipulation, date and time formatting, encoding/decoding operations, and random key generation. The repository aims to provide developers with a set of reliable and reusable functions to streamline their development workflow.

## List of Features
- **String Manipulation**:
  - `ToLower`: Convert strings to lowercase.
  - `StringToInt`: Convert strings to integers.
  - `RemoveOrdinalSuffix`: Remove ordinal suffixes from strings.
- **Date and Time Formatting**:
  - `GetHour`, `GetDate`, `GetTime`, `GetDateTime`, `GetDateTimeLong`, `GetDateShort`, `GetTimestamp`: Various functions to get and format the current date and time.
  - `FormatDuration`: Format a duration in a human-readable way.
- **Encoding/Decoding**:
  - `Base64Encode`, `Base64Decode`: Encode and decode strings using Base64.
  - `EncodeURL`, `DecodeURL`: Encode and decode URLs.
- **Random Key Generation**:
  - `GenerateShortLink`, `RandStringBytesRmndr`, `RandKey`, `GenerateKey`: Functions to generate random strings and keys.
- **Set Operations**:
  - `Contains`, `Uniq`, `Difference`: Functions to perform operations on sets.
- **Testing**:
  - Comprehensive test functions for each utility function to ensure reliability and correctness.

## How to Run the Project
To run the project, follow these steps:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/uug-ai/utils.git
   cd utils
   ```

2. **Run the Project**:
   Ensure you have Python installed, then run the main script:
   ```bash
   python main.py
   ```

## Testing Instructions
To run the tests for the project, follow these steps:

1. **Navigate to the Tests Directory**:
   ```bash
   cd tests
   ```

2. **Run the Tests**:
   Use the following command to run all the tests:
   ```bash
   python -m unittest discover
   ```

The test functions include:
- `TestContains`
- `TestUniq`
- `TestDifference`
- `TestGetDate`
- `TestGetHour`
- `TestGetTime`
- `TestGetDateTime`
- `TestGetDateTimeLong`
- `TestGetDateShort`
- `TestGetTimestamp`
- `TestFormatDuration`
- `TestBase64Encode`
- `TestBase64Decode`
- `TestEncodeURL`
- `TestDecodeURL`
- `TestToLower`
- `TestStringToInt`
- `TestRemoveOrdinalSuffix`
- `TestRandStringBytesRmndr`
- `TestRandKey`
- `TestGenerateShortLink`
- `TestRandKeyErrorHandling`
- `TestGenerateKey`

## How to Contribute to the Project
We welcome contributions to the `uug-ai/utils` project! To contribute, please follow these guidelines:

1. **Fork the Repository**:
   Click the "Fork" button on the repository page to create a copy of the repository in your GitHub account.

2. **Clone Your Fork**:
   ```bash
   git clone https://github.com/YOUR_USERNAME/utils.git
   cd utils
   ```

3. **Create a New Branch**:
   ```bash
   git checkout -b feature-branch
   ```

4. **Make Your Changes**:
   Implement your changes or additions.

5. **Commit and Push**:
   ```bash
   git add .
   git commit -m "Description of changes"
   git push origin feature-branch
   ```

6. **Create a Pull Request**:
   Navigate to the original repository and click "New Pull Request". Provide a detailed description of your changes and submit the pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.