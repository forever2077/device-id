@echo off
setlocal enabledelayedexpansion

:: 输入目标文件名称
set /p output_base_name="Enter the base name of the output files: "

:: 定义目标平台
set platforms=windows/amd64 linux/amd64 darwin/amd64

:: 分割和编译
for %%i in (%platforms%) do (
    for /f "tokens=1,2 delims=/" %%a in ("%%i") do (
        set GOOS=%%a
        set GOARCH=%%b
    )
    set output_name=%output_base_name%_!GOOS!_!GOARCH!
    if "!GOOS!"=="windows" (
        set output_name=!output_name!.exe
    )
    echo Building for !GOOS!, !GOARCH! into !output_name!...

    set GOARCH=!GOARCH!
    set GOOS=!GOOS!
    go build -ldflags="-s -w" -o !output_name!

    if errorlevel 1 (
        echo An error has occurred! Aborting.
        exit /b 1
    )
)
