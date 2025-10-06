# oscalctl
A Golang OSCAL Automation Tool

![oscalctl vision](docs/diagrams/oscalctl-concept.drawio-1.png)

## Overview

oscalctl is a command-line tool designed to automate Security Technical Implementation Guides (STIG) compliance checks using the Open Security Controls Assessment Language (OSCAL). Currently, it focuses on generating OSCAL component definitions, with more functionality planned for future releases.

## Installation

### Prerequisites
- Go 1.18 or higher

### Installing from source
```bash
git clone https://github.com/open-automation-construct/oscalctl.git
cd oscalctl
go build
```

## Current Functionality

At present, oscalctl supports the following command:

```bash
oscalctl generate oscal component [flags]
```

```bash
oscalctl generate oscal component -i /path/to/checklist.cklb -o /path/to/my/new/oscalComponent.json
```

### With a custom CCI Mapping
```bash
oscalctl generate oscal component -i /path/to/checklist.cklb -o /path/to/my/new/oscalComponent.json --cci-map /path/to/custom/cci.xml
```

### With a custom title
```bash
./oscalctl generate oscal component -i references/cklb/testdata/aaa-srg.cklb.json -o test
.json --title mynewcomponent
```

This command allows you to generate OSCAL component definition files, which are essential for describing system components that are subject to security controls.

## Using Configuration Files

oscalctl supports configuration files for setting default values and managing complex configurations. The tool will look for configuration files in the following locations:

- The current directory (./config.yaml, ./config.json)
- The user's home directory (~/.oscalctl/config.yaml)
- System-wide configuration (/etc/oscalctl/config.yaml)

### Configuration File Format

Configuration files can be in YAML or JSON format. Here's an example of a YAML configuration file:

```yaml
# config.yaml
oscal:
  title: "My Custom OSCAL Component Title"
  component:
    input: "/path/to/checklist.cklb"
    output: "/path/to/output.json"
    cciMap: "/path/to/custom/cci.xml"
```

The same configuration in JSON format:

```json
{
  "oscal": {
    "title": "My Custom OSCAL Component Title",
    "component": {
      "input": "/path/to/checklist.cklb",
      "output": "/path/to/output.json",
      "cciMap": "/path/to/custom/cci.xml"
    }
  }
}
```

### Configuration Precedence

oscalctl uses the following precedence order when resolving configuration values (from lowest to highest priority):

1. Default values
2. Configuration file values
3. Environment variables (prefixed with OSCALCTL_)
4. Command-line flags

This means that command-line flags will always override values set in configuration files or environment variables.

## Usage Examples

### Generate an OSCAL Component Definition

The basic syntax for generating a component definition is:

```bash
oscalctl generate oscal component -i <input_checklist> -o <output_file> [flags]
```

#### Example with input and output files

```bash
oscalctl generate oscal component -i checklist.cklb -o component.json
```

#### Example with custom title

```bash
oscalctl generate oscal component -i checklist.cklb -o component.json -t "Database Server"
```

#### Example with custom CCI mapping

```bash
oscalctl generate oscal component -i checklist.cklb -o component.json --cci-map custom_cci.xml
```

### Available Flags for Component Generation

- `--title`, `-t`: Custom title for the OSCAL document
- `--input`, `-i`: Path to the STIG checklist (required)
- `--output`, `-o`: Path to the output OSCAL component definition (required)
- `--cci-map`: Path to a custom CCI XML document (optional)

## Command Help

To view detailed help information for the available commands:

```bash
oscalctl --help
oscalctl generate --help
oscalctl generate oscal --help
oscalctl generate oscal component --help
```

## Future Functionality

In future releases, oscalctl will support additional commands for:

- Importing and managing STIG catalogs
- Creating and customizing OSCAL profiles
- Performing assessments against target systems
- Generating compliance reports
- Converting between various OSCAL formats

## Troubleshooting

### Enable Verbose Output
```bash
oscalctl --verbose generate oscal component -i checklist.cklb -o component.json
```

### Common Issues

- **File not found errors**: Ensure all file paths are correct and files exist
- **Format errors**: Verify your STIG checklist is in valid CKLB format
- **Permission issues**: Check that you have appropriate read/write permissions

## Contributing

Contributions to oscalctl are welcome! As this tool is in active development, feedback and contributions can help shape its direction. Please see our contribution guidelines for more information on how to get involved.

## License

oscalctl is open source software licensed under [LICENSE].