# SSH Win

```ps
Get-WindowsCapability -Online | Where-Object Name -like 'OpenSSH*'

Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0

# Jalankan di PowerShell (Admin)
Start-Service sshd
# Atur agar otomatis berjalan saat startup
Set-Service -Name sshd -StartupType 'Automatic'

Subsystem powershell C:\Program Files\PowerShell\7\pwsh.exe -sshs -NoLogo -NoProfile

```
