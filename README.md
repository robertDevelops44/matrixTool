
# matrixTool

matrixTool is a CLI application that takes an Excel file with matrix pricing and produces filtered pricing with custom broker fee and adjustments.

## Installation

To install matrixTool, you can use `go install`. Make sure you have Go installed on your system before proceeding with the installation.

```shell
go install github.com/rh5661/matrixTool@latest
```

## Usage

`matrixTool.exe` can be run in a terminal (Command Prompt, Powershell) to start using

Or

Run `launch.bat`

The basic usage of matrixTool is as follows:

```shell
matrixTool [command]
```

### Available Commands

- `load`            :   Loads matrix excel file into database
- `setStart`        :   Set the parameter for contract start date
- `setUtil`         :   Set the parameter for utility
- `setDualBilling`  :   Set the parameter to include dual billing
- `setTerms`        :   Set the parameter for terms
- `setMils`         :   Set the parameter for the broker fee in mils
- `showParameters`  :   Displays current parameters entered
- `showUtils`       :   Displays all valid utilities with their code & full name
- `generate`        :   Creates a new sheet with filtered pricing and margin
- `help`            :   Help about any command


You can use the `--help` flag with any command to get more information about it. For example:

```shell
matrixTool load --help
```

### Examples

- Load a Matrix Pricing Excel file:

  ```shell
  matrixTool load "C:\Users\Robert\Documents\matrixTool\Daily Matrix Price For All Term.xlsx"
  ```

- Show all parameters:

  ```shell
  matrixTool showParameters
  ```

- Show all valid utility codes/names:

  ```shell
  matrixTool showUtils
  ```

- Set parameter for start date:

  ```shell
  matrixTool setStart "Sep-23"
  ```

- Set parameter for utility code:

  ```shell
  matrixTool setUtil "PPL"
  ```

- Set parameter for dual billing inclusion:

  ```shell
  matrixTool setDualBilling "No"
  ```

- Set parameter for terms to include:

  ```shell
  matrixTool setStart "[12,24,36,48]"
  ```

- Set parameter for amount of broker fee to insert (mils):

  ```shell
  matrixTool setMils "15"
  ``` 

- Generate a sheet with filtered pricing with broker fee:

  ```shell
  matrixTool generate
  ``` 



### Dependencies

matrixTool depends on the following libraries:

- `github.com/spf13/cobra` (CLI framework for app)
- `modernc.org/sqlite` (non-cgo SQLite3 driver)
- `github.com/xuri/excelize` (for reading and writing with .xlsx files)
- `github.com/mailru/easyjson` (for reading and writing parameters to json file)
- `github.com/pquerna/ffjson`  (for marshalling/unmarshalling into strings)

## Acknowledgements

matrixTool was written in Go. Much appreciation to the following projects and their contributors for their work:

- [cobra](https://cobra.dev/)
- [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)
- [excelize](https://xuri.me/excelize/)
- [easyjson](https://pkg.go.dev/github.com/tcolar/easyjson)
- [ffjson](https://pkg.go.dev/github.com/pquerna/ffjson/ffjson)
- [Go](https://go.dev/)

## License

This project is licensed under the [MIT License](LICENSE).

---

Thank you for using matrixTool! If you encounter any issues or have suggestions for improvements, please open an issue on the [GitHub Repository](https://github.com/rh5661/matrixTool).