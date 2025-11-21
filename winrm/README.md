# WinRM

di client jalankan

```ps
Get-NetConnectionProfile

Set-NetConnectionProfile -InterfaceAlias "Wi-Fi" -NetworkCategory Private

Enable-PSRemoting -Force

New-LocalUser -Name Username -Password (ConvertTo-SecureString "Passwordnya213" -AsPlainText -Force) -PasswordNeverExpires -Description "User untuk akses WinRM"

Add-LocalGroupMember -Group "Administrators" -Member Astrajingga
```
