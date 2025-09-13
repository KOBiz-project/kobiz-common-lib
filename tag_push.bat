@echo off
setlocal enabledelayedexpansion

for /f "tokens=*" %%i in ('git describe --tags --abbrev^=0 2^>nul') do set "LATEST_TAG=%%i"

if not defined LATEST_TAG (
    echo No tags found in repository
    exit /b 1
)

echo Latest tag found: !LATEST_TAG!
echo Pushing tag !LATEST_TAG! to origin...
git push origin "!LATEST_TAG!"

echo Done!