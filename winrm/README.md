# WinRM

di client jalankan

```ps
Get-NetConnectionProfile

Set-NetConnectionProfile -InterfaceAlias "Wi-Fi" -NetworkCategory Private

Enable-PSRemoting -Force

New-LocalUser -Name Username -Password (ConvertTo-SecureString "Passwordnya213" -AsPlainText -Force) -PasswordNeverExpires -Description "User untuk akses WinRM"

Add-LocalGroupMember -Group "Administrators" -Member Username

# Perintah 1: Aktifkan WinRM (jika belum)
winrm quickconfig

# Perintah 2: Izinkan autentikasi Basic (wajib untuk username/password)
Set-Item -Path 'WSMan:\localhost\Service\Auth\Basic' -Value $true

# Perintah 3: Izinkan koneksi tidak terenkripsi (wajib untuk HTTP)
Set-Item -Path 'WSMan:\localhost\Service\AllowUnencrypted' -Value $true

Get-ChildItem WSMan:\localhost\Service\Auth
Get-Item WSMan:\localhost\Service\AllowUnencrypted
```
