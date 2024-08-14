# TIG: A Lightweight Git Clone

**tig** is a simple version control system inspired by Git, offering basic functionalities like adding, committing, cloning, initializing repositories, and viewing logs. Use `tig` to manage your project’s version control with ease.

## Features

- **Add**: Stage files for commit.
     - Usage: `tig add <file_name>` or `tig add .` to stage all changes.
- **Commit**: Commit staged changes with a message.

     - Usage: `tig commit -m "Your commit message"`

- **Clone**: Clone a repository from a remote source.

     - Usage: `tig clone <repo_url>`

- **Init**: Initialize a new repository.

     - Usage: `tig init`

- **Log**: View the commit history.
     - Usage: `tig log`

## Installation

Follow these steps to install and use **tig**:

1. **Install Go**:

      - Make sure you have Go installed on your system. You can download and install it from [golang.org](https://golang.org/dl/).

2. **Clone the Repository**:

      - Clone the **tig** repository:
           ```bash
           git clone https://github.com/Prayag2003/tig-version-control-system
           ```

3. **Navigate to the `tig` Directory**:

      - Change directory to the `tig` command source:
           ```bash
           cd cmd/tig
           ```

4. **Set Up an Alias**:

      - Create an alias for `tig` to run it easily from the command line:
           ```bash
           alias tig="go run main.go"
           ```

5. **Start Using `tig`**:
      - Now you can use `tig` just like Git:
           ```bash
           tig init
           tig add <file_name>
           tig commit -m "Your commit message"
           ```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve **tig**.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
