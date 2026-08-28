$InstallDir = "$env:LOCALAPPDATA\TodoTerminal"

if (Test-Path $InstallDir) {
    Remove-Item -Recurse -Force $InstallDir
    Write-Host "todo has been uninstalled." -ForegroundColor Green
} else {
    Write-Host "todo is not installed." -ForegroundColor Yellow
}