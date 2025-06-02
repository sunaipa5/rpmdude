# Rpmdude

Simple rpm build tool

## Build

> [GO](https://go.dev) required

```bash
git clone https://github.com/sunaipa5/rpmdude

go build
```

## Usage

```bash
//Init project
rpmdude init my-project

//Build RPM
rpmdude build
```

## Tree

> Modify the .spec file according to your project

```bash
rpmdude_build
├── rpmdude_build.sh
├── SOURCES
│   └── my-project.desktop
└── SPECS
    └── my-project.spec

```
