# G-init - Project Initialization CLI Tool

Govial-init (G-init) is a user-friendly CLI tool written in Go, designed to speedup the setup of new projects across various languages, including Terraform, Python, Go, and others. 

## Key Features

- **Predefined Structures and Configurations**: G-init generates project skeletons with industry-standard best practices, tailored to each technology.
- **Customizable Templates**: Each language or technology is supported with a default template file that can be easily customized to fit your unique workflows and preferences.
- **Shortcut Commands**: Use simple commands to initialize projects in seconds. For example:
  - `ginit tf .` sets up a Terraform project with a minimal recommended structure, a custom pre-commit hook, and a .gitignore file.
  - `ginit tf --module` initializes a Terraform module with the necessary configurations.
  - `ginit py` creates a Python project with a virtual environment and other essential files.
- **Written in Go**: Built for performance and simplicity, G-init ensures a fast, lightweight, and reliable experience.
- **Extensibility**: Easily add support for new technologies or customize existing setups to align with evolving requirements.

## Why Use G-init?

G-init saves developers time by automating project setup, enforcing consistent structures, and providing a solid foundation for various technologies. Whether you're starting a Terraform module, a Python application, or a Kubernetes configuration, G-init ensures your projects start with best practices baked in.

## Installation

To install G-init, clone the repository and build the project:

```bash
git clone https://github.com/vakintosh/Govial_Init.git
cd govial
go build -o ginit ./cmd/ginit/main.go
```

## Usage

After building the project, you can use the `ginit` command followed by the desired shortcut to initialize your project. For example:

```bash
ginit tf .
```

This command will set up a new Terraform project within the current work directory with the recommended structure.

To create a new Terraform project in a specific directory, you can use:
```bash
ginit tf my-project
```
This command will set up a new Terraform project in the my-tf-project directory with the recommended structure.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.