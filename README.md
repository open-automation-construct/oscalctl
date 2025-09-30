# stigctl
A Golang STIG/OSCAL Automation Tool

![stigctl vision](stigctl-concept.drawio-1.png)

## Overview

stigctl is a command-line tool designed to automate Security Technical Implementation Guides (STIG) compliance checks using the Open Security Controls Assessment Language (OSCAL). Currently, it focuses on generating OSCAL component definitions, with more functionality planned for future releases.

## Installation

### Prerequisites
- Go 1.18 or higher

### Installing from source
```bash
git clone https://github.com/yourusername/stigctl.git
cd stigctl
go build
```

## Current Functionality

At present, stigctl supports the following command:

```bash
stigctl generate oscal component [flags]
```

```bash
stigctl generate oscal component -i /path/to/checklist.cklb -o /path/to/my/new/oscalComponent.json
```

```bash
stigctl generate oscal component -i /path/to/checklist.cklb -o /path/to/my/new/oscalComponent.json --cci-map /path/to/custom/cci.xml
```

This command allows you to generate OSCAL component definition files, which are essential for describing system components that are subject to security controls.

## Usage Examples (Future)

### Generate an OSCAL Component Definition

The basic syntax for generating a component definition is:

```bash
stigctl generate oscal component --title "Component Name" --id component-id
```

#### Example with minimal parameters

```bash
stigctl generate oscal component --title "Web Server" --id web-server-01
```

This creates a basic OSCAL component definition for a web server.

#### Example with additional metadata

```bash
stigctl generate oscal component --title "Database Server" --id db-server-01 --version 1.0 --description "Primary PostgreSQL database server"
```

#### Example with output file specification

```bash
stigctl generate oscal component --title "Application Server" --id app-server-01 --output app-server-component.json
```

This generates the component definition and saves it to the specified file.

#### Example with component type

```bash
stigctl generate oscal component --title "Network Firewall" --id firewall-01 --type "infrastructure"
```

### Available Flags for Component Generation

- `--title`: The name of the component (required)
- `--id`: Unique identifier for the component (required)
- `--description`: Detailed description of the component
- `--version`: Version information for the component
- `--type`: Type of component (e.g., software, hardware, service, infrastructure)
- `--output`: Destination file for the generated OSCAL component definition

## Command Help

To view detailed help information for the available commands:

```bash
stigctl --help
stigctl generate --help
stigctl generate oscal --help
stigctl generate oscal component --help
```

## Future Functionality

In future releases, stigctl will support additional commands for:

- Importing and managing STIG catalogs
- Creating and customizing OSCAL profiles
- Performing assessments against target systems
- Generating compliance reports

## Troubleshooting

### Enable Verbose Output
```bash
stigctl --verbose generate oscal component --title "Web Server" --id web-server-01
```

## Contributing

Contributions are welcome! As this tool is in early development, feedback and contributions can help shape its direction.