$ErrorActionPreference = "Stop"

$Repo = "johnathantam/TodoTerminal"
$BinaryName = "todo.exe"
$InstallDir = "$env:LOCALAPPDATA\TodoTerminal"

$LatestTag = (Invoke-RestMethod "https://api.github.com/repos/$Repo/releases/latest").tag_name
$Archive = "TodoTerminal_Windows_x86_64.zip"
$Url = "https://github.com/$Repo/releases/download/$LatestTag/$Archive"

New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
$TmpZip = "$env:TEMP\$Archive"

Write-Host "Downloading todo $LatestTag..."
Invoke-WebRequest -Uri $Url -OutFile $TmpZip
Expand-Archive -Path $TmpZip -DestinationPath $InstallDir -Force
Remove-Item $TmpZip

$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notlike "*$InstallDir*") {
    [Environment]::SetEnvironmentVariable("Path", "$UserPath;$InstallDir", "User")
    Write-Host "Added $InstallDir to PATH. Restart your terminal for it to take effect."
}

Write-Host "Installed! Run 'todo help' to get started."