@echo off

rem Change directory to the directory containing the executable
cd /d "%~dp0"

rem Run the executable with the provided arguments
tig.exe %*
