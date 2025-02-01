package main

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/fatih/color"
    "github.com/gorilla/handlers"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/host"
    "github.com/shirou/gopsutil/v3/mem"
    "log"
)

// Сервисы и связанные с ними хосты
var services = map[string][]string{
    "Service A": {"host1.example.com", "host2.example.com"},
    "Service B": {"host3.example.com", "host4.example.com"},
}

// Информация о сертификате
type CertificateInfo struct {
    Host           string    `json:"host"`
    Issuer         string    `json:"issuer"`
    ValidFrom      time.Time `json:"valid_from"` // Добавляем запятую
    ValidUntil     time.Time `json:"valid_until"` // Добавляем запятую
    SerialNumber   string    `json:"serial_number"` // Добавляем запятую
    SystemTime     string    `json:"system_time"` // Добавляем запятую
    Uptime         uint64    `json:"uptime"` // Добавляем запятую
    CPUUsage       float64   `json:"cpu_usage"` // Добавляем запятую
    RAMUsage       float64   `json:"ram_usage"` // Добавляем запятую
    TargetPath     string    `json:"target_path"` // Добавляем запятую
}

// Информация о системной нагрузке
type HostInfo struct {
    SystemTime string  `json:"system_time"` // Добавляем запятую
    Uptime     uint64  `json:"uptime"` // Добавляем запятую
    CPUUsage   float64 `json:"cpu_usage"` // Добавляем запятую
    RAMUsage   float64 `json:"ram_usage"` // Добавляем запятую
}

// Моковые данные для примера
func getCertificateInfo(host string) CertificateInfo {
    info := hostInfo()
    return CertificateInfo{
        Host:        host,
        Issuer:      "Let's Encrypt",
        ValidFrom:   time.Now().Add(-time.Hour * 24 * 30),
        ValidUntil:  time.Now().Add(time.Hour * 24 * 60),
        SerialNumber: "1234567890",
        SystemTime: info.SystemTime,
        Uptime:     info.Uptime,
        CPUUsage:   info.CPUUsage,
        RAMUsage:   info.RAMUsage,
        TargetPath: "/path/to/certificates/" + host + ".crt",
    }
}

func hostInfo() HostInfo {
    sysTime, _ := host.BootTime()
    bootTime := time.Unix(int64(sysTime), 0)
    upTime := time.Since(bootTime).Seconds()
    cpuPercent, _ := cpu.Percent(time.Second, false)
    memInfo, _ := mem.VirtualMemory()
    return HostInfo{
        SystemTime: time.Now().Format(time.RFC1123),
        Uptime:     uint64(upTime),
        CPUUsage:   cpuPercent[0],
        RAMUsage:   memInfo.UsedPercent,
    }
}

// Обработчик запроса для получения информации о сервисах
func servicesHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(services)
}

// Обработчик запроса для получения информации о сертификате по хосту
func certificateHandler(w http.ResponseWriter, r *http.Request) {
    host := r.URL.Query().Get("host")
    if host == "" {
        http.Error(w, "Host is required", http.StatusBadRequest)
        return
    }
    certInfo := getCertificateInfo(host)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(certInfo)
}

// Обработчик запроса для импорта данных из JSON файла
func importServicesHandler(w http.ResponseWriter, r *http.Request) {
    var importedServices map[string][]string
    err := json.NewDecoder(r.Body).Decode(&importedServices)
    if err != nil {
        http.Error(w, "Invalid JSON data", http.StatusBadRequest)
        return
    }
    services = importedServices
    w.WriteHeader(http.StatusOK)
}

// Обработчик запроса для экспорта данных в JSON файл
func exportServicesHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(services)
}

// Обработчик запроса для получения информации о системной нагрузке по хосту
func systemInfoHandler(w http.ResponseWriter, r *http.Request) {
    host := r.URL.Query().Get("host")
    if host == "" {
        http.Error(w, "Host is required", http.StatusBadRequest)
        return
    }
    systemInfo := hostInfo()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(systemInfo)
}

func main() {
    color.Cyan("Starting SSL Certificate Checker...")

    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type"}),
    )

    http.Handle("/services", corsHandler(http.HandlerFunc(servicesHandler)))
    http.Handle("/certificate", corsHandler(http.HandlerFunc(certificateHandler)))
    http.Handle("/import", corsHandler(http.HandlerFunc(importServicesHandler)))
    http.Handle("/export", corsHandler(http.HandlerFunc(exportServicesHandler)))
    http.Handle("/systeminfo", corsHandler(http.HandlerFunc(systemInfoHandler)))

    log.Println("Server started at :8082")
    if err := http.ListenAndServe(":8082", nil); err != nil {
        log.Fatalf("Error starting server: %s\n", err)
    }
}
