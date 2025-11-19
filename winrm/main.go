package main

import (
	"fmt"

	"github.com/masterzen/winrm"
)

func main() {
	// --- Konfigurasi Koneksi ---
	endpoint := winrm.NewEndpoint(
		"IP_TARGET_WINDOWS", // Ganti dengan IP target
		5985,                // Port HTTP untuk WinRM
		false,               // Tidak menggunakan HTTPS
	)

	// --- Kredensial ---
	client, err := winrm.NewClient(
		endpoint,
		"USERNAME", // Ganti dengan username
		"PASSWORD", // Ganti dengan password
	)
	if err != nil {
		panic(err)
	}

	// --- Eksekusi Perintah PowerShell ---
	// Perintah yang akan dijalankan di mesin target
	psCommand := "Get-Process | Select-Object -First 5 | ConvertTo-Json"

	fmt.Printf("Menjalankan perintah: %s\n", psCommand)

	// Jalankan perintah dan tangkap outputnya
	stdout, stderr, exitCode, err := client.RunWithString(psCommand, "")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// --- Tampilkan Hasil ---
	fmt.Printf("Exit Code: %d\n", exitCode)
	fmt.Println("--- STDOUT ---")
	fmt.Println(stdout)
	if stderr != "" {
		fmt.Println("--- STDERR ---")
		fmt.Println(stderr)
	}
}
