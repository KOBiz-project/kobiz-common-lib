@echo off
setlocal enabledelayedexpansion

rem Get latest tag
for /f "tokens=*" %%i in ('git describe --tags --abbrev^=0 2^>nul') do set "LATEST_TAG=%%i"

if not defined LATEST_TAG (
    echo No tags found. Creating first tag v0.0.1
    set "NEW_TAG=v0.0.1"
    goto create_tag
)

echo Latest tag: !LATEST_TAG!

rem Parse version numbers (assuming format v0.0.0)
set "VERSION=!LATEST_TAG:v=!"
for /f "tokens=1,2,3 delims=." %%a in ("!VERSION!") do (
    set "MAJOR=%%a"
    set "MINOR=%%b"
    set "PATCH=%%c"
)

rem Increment patch version
set /a "PATCH=!PATCH!+1"
set "NEW_TAG=v!MAJOR!.!MINOR!.!PATCH!"

:create_tag
echo Creating new tag: !NEW_TAG!
git tag "!NEW_TAG!"

if !errorlevel! equ 0 (
    echo Pushing tag !NEW_TAG! to origin...
    git push origin "!NEW_TAG!"
    echo Done! Created and pushed tag: !NEW_TAG!
) else (
    echo Failed to create tag !NEW_TAG!
    exit /b 1
)