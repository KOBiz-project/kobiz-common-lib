@echo off
if "%1"=="" (
    echo Usage: tag_push.bat ^<tag_name^>
    echo Example: tag_push.bat v1.0.0
    exit /b 1
)

echo Creating git tag: %1
git tag %1

echo Pushing tag %1 to origin...
git push origin %1

echo Done!