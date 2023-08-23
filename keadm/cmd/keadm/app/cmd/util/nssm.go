package util

import (
	"fmt"
	"strings"
)

const installNSSMScript = `
function DownloadFile($destination, $source) {
    Write-Host("Downloading $source to $destination")
    curl.exe --silent --fail -Lo $destination $source

    if (!$?) {
        Write-Error "Download $source failed"
        exit 1
    }
}

$global:NssmInstallDirectory = "$env:ProgramFiles\nssm"

[Environment]::SetEnvironmentVariable("Path", $env:Path, [System.EnvironmentVariableTarget]::Machine)

Write-Host "Installing nssm"
$arch = "win32"
if ([Environment]::Is64BitOperatingSystem) {
    $arch = "win64"
}

mkdir -Force $global:NssmInstallDirectory
DownloadFile nssm.zip https://nssm.cc/ci/nssm-2.24-101-g897c7ad.zip
tar C $global:NssmInstallDirectory -xvf .\nssm.zip --strip-components 2 */$arch/*.exe
Remove-Item -Force .\nssm.zip

$env:path += ";$global:NssmInstallDirectory"
$newPath = "$global:NssmInstallDirectory;" +
        [Environment]::GetEnvironmentVariable("PATH", [EnvironmentVariableTarget]::Machine)

[Environment]::SetEnvironmentVariable("PATH", $newPath, [EnvironmentVariableTarget]::Machine)
`

func IsNSSMInstalled() bool {
	cmd := NewCommand("nssm version")
	err := cmd.Exec()
	return err == nil
}

func InstallNSSM() error {
	cmd := NewCommand(installNSSMScript)
	err := cmd.Exec()
	return err
}

func InstallNSSMService(name, path string, args ...string) error {
	cmd := NewCommand(fmt.Sprintf("nssm install %s %s %s", name, path, strings.Join(args, " ")))
	return cmd.Exec()
}

func IsNSSMServiceExist(service string) bool {
	cmd := NewCommand("nssm status " + service)
	_err := cmd.Exec()
	return _err == nil
}

func IsNSSMServiceRunning(service string) bool {
	cmd := NewCommand("nssm status " + service)
	_err := cmd.Exec()
	if _err != nil || cmd.ExitCode > 0 {
		return false
	}
	if cmd.GetStdOut() == "SERVICE_RUNNING" {
		return true
	}
	return false
}

func StartNSSMService(service string) error {
	cmd := NewCommand("nssm start " + service)
	return cmd.Exec()
}

func StopNSSMService(service string) error {
	cmd := NewCommand("nssm stop " + service)
	return cmd.Exec()
}

func SetNSSMServiceStdout(service string, file string) error {
	cmd := NewCommand(strings.Join([]string{"nssm", "set", service, "AppStdout", file}, " "))
	return cmd.Exec()
}

func SetNSSMServiceStderr(service string, file string) error {
	cmd := NewCommand(strings.Join([]string{"nssm", "set", service, "AppStderr", file}, " "))
	return cmd.Exec()
}

func UninstallNSSMService(service string) error {
	cmd := NewCommand(strings.Join([]string{"nssm", "remove", service, "confirm"}, " "))
	return cmd.Exec()
}
