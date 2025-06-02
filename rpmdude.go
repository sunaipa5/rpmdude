package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: rpmdude <init <app-name> | build>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "init":
		if len(os.Args) < 3 {
			fmt.Println("Application name required: rpmdude init <app-name>")
			os.Exit(1)
		}
		appName := os.Args[2]
		initProject(appName)
	case "build":
		buildProject()
	default:
		fmt.Println("Unknown command:", command)
	}
}

// initProject creates the RPM build folder structure and template files for the given app name.
func initProject(app string) {
	baseDir := "rpmdude_build"

	// Check if baseDir already exists
	if _, err := os.Stat(baseDir); err == nil {
		fmt.Printf("⚠️  Directory '%s' already exists.\n", baseDir)
		fmt.Println("Creating a new template here may overwrite existing files or cause conflicts.")
		fmt.Println("Please remove the directory first or run this command in a different location.")
		os.Exit(1)
	}

	specsDir := filepath.Join(baseDir, "SPECS")
	sourcesDir := filepath.Join(baseDir, "SOURCES")

	// Create directories
	os.MkdirAll(specsDir, 0755)
	os.MkdirAll(sourcesDir, 0755)

	// RPM spec file content
	spec := fmt.Sprintf(`Name:           %s
Version:        1.0
Release:        1%%{?dist}
Summary:        %s system tray application
License:        GPLv3

Source0:        %s
Source1:        %s.png
Source2:        %s.desktop

%%description
%s is a system tray application.

%%install
rm -rf %%{buildroot}
mkdir -p %%{buildroot}/usr/bin/
mkdir -p %%{buildroot}/usr/share/icons/
mkdir -p %%{buildroot}/usr/share/applications/
install -m 755 %%{SOURCE0} %%{buildroot}/usr/bin/
install -m 644 %%{SOURCE1} %%{buildroot}/usr/share/icons/
install -m 644 %%{SOURCE2} %%{buildroot}/usr/share/applications/

%%files
/usr/bin/%s
/usr/share/icons/%s.png
/usr/share/applications/%s.desktop

%%changelog
* Mon Jun 02 2025 Developer <you@example.com> - 1.0-1
- Initial RPM release
`, app, app, app, app, app, app, app, app, app)

	// .desktop file content
	desktop := fmt.Sprintf(`[Desktop Entry]
Name=%s
Exec=%s
Icon=%s
Type=Application
Categories=Utility;
Terminal=false
`, app, app, app)

	// rpmdude_build.sh content
	buildSh := fmt.Sprintf(`#!/bin/bash
set -e
TOPDIR=$(pwd)
rpmbuild --define "_topdir $TOPDIR" -ba SPECS/%s.spec
`, app)

	// Write files
	os.WriteFile(filepath.Join(specsDir, app+".spec"), []byte(spec), 0644)
	os.WriteFile(filepath.Join(sourcesDir, app+".desktop"), []byte(desktop), 0644)
	os.WriteFile(filepath.Join(baseDir, "rpmdude_build.sh"), []byte(buildSh), 0755)

	fmt.Println("✔ Project initialized in", baseDir)
}

// buildProject runs the rpmdude_build.sh script inside the rpmdude_build directory.
func buildProject() {
	baseDir := "rpmdude_build"
	buildScript := filepath.Join(baseDir, "rpmdude_build.sh")

	if _, err := os.Stat(buildScript); os.IsNotExist(err) {
		fmt.Println("❌ rpmdude_build.sh not found. Run `rpmdude init <app>` first.")
		os.Exit(1)
	}

	cmd := exec.Command("bash", "rpmdude_build.sh")
	cmd.Dir = baseDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("🚧 Starting build process...")
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Build failed:", err)
		os.Exit(1)
	}

	fmt.Println("✅ Build finished successfully.")
}
