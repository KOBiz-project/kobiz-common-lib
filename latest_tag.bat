@echo off
chcp 65001 >nul
echo === 최신 Git Tag 조회 ===
echo.

echo [최신 태그 1개]
git describe --tags --abbrev=0
echo.