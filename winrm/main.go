package main

import (
	"bufio" // Untuk membaca input
	"fmt"
	"os"      // Untuk mengakses standard input (keyboard)
	"strings" // Untuk membersihkan input dari spasi atau karakter baru
	"time"

	"github.com/masterzen/winrm"
)

func main() {
	// Buat reader untuk membaca input dari keyboard
	reader := bufio.NewReader(os.Stdin)

	// --- Minta Input dari User ---
	fmt.Print("Masukkan Host Target: ")
	host, _ := reader.ReadString('\n')
	host = strings.TrimSpace(host) // Hapus spasi dan karakter newline di akhir

	fmt.Print("Masukkan Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Masukkan Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Println("\n--- Menghubungkan ke", host, "---")

	// --- Konfigurasi Koneksi ---
	endpoint := winrm.NewEndpoint(
		host,           // Gunakan host dari input user
		5985,           // Port HTTP untuk WinRM
		false,          // Tidak menggunakan HTTPS
		false,          // insecure (bool)
		nil,            // caCert ([]byte)
		nil,            // cert ([]byte)
		nil,            // key ([]byte)
		time.Second*60, // timeout (time.Duration)
	)

	// --- Kredensial & Parameter ---
	// Gunakan username dan password dari input user
	// PERBAIKAN: Hapus tanda kurung () karena DefaultParameters adalah variabel, bukan fungsi
	params := winrm.DefaultParameters
	params.TransportDecorator = func() winrm.Transporter {
		return &winrm.ClientAuthRequest{} // Menggunakan Basic Auth
	}

	client, err := winrm.NewClientWithParameters(endpoint, username, password, params)
	if err != nil {
		panic(err)
	}

	// --- Eksekusi Perintah PowerShell ---
	psCommand := "Get-Process | Select-Object -First 5 | ConvertTo-Json"

	fmt.Printf("Menjalankan perintah: %s\n", psCommand)

	// Jalankan perintah dan tangkap outputnya
	stdout, stderr, exitCode, err := client.RunWithString(psCommand, "")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// --- Tampilkan Hasil ---
	fmt.Printf("\nExit Code: %d\n", exitCode)
	fmt.Println("--- STDOUT ---")
	fmt.Println(stdout)
	if stderr != "" {
		fmt.Println("--- STDERR ---")
		fmt.Println(stderr)
	}
}
