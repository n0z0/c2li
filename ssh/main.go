package main

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func main() {
	// --- Konfigurasi Koneksi SSH ---
	config := &ssh.ClientConfig{
		User: "USERNAME", // Ganti dengan username
		Auth: []ssh.AuthMethod{
			ssh.Password("PASSWORD"), // Ganti dengan password
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Hanya untuk testing, tidak aman untuk produksi
	}

	// --- Hubungkan ke Server ---
	client, err := ssh.Dial("tcp", "IP_TARGET_WINDOWS:22", config)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// --- Buat Sesi ---
	session, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// --- Eksekusi Perintah PowerShell ---
	// Perintah dipanggil melalui `pwsh`
	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	psCommand := `pwsh -Command "Get-Service | Where-Object {$_.Status -eq 'Running'} | Select-Object -First 3 | ConvertTo-Csv -NoTypeInformation"`
	err = session.Run(psCommand)
	if err != nil {
		fmt.Printf("Error menjalankan perintah: %v\n", err)
	}

	// --- Tampilkan Hasil ---
	fmt.Println("--- STDOUT ---")
	fmt.Println(stdout.String())
	if stderr.Len() > 0 {
		fmt.Println("--- STDERR ---")
		fmt.Println(stderr.String())
	}
}
