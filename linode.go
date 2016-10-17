package main

import "time"
import "net/http"
import "encoding/json"

type LinodeClient struct {
    Token string
}

func (c LinodeClient) Request(snippet string, result interface{}) (error) {
    endpoint := "https://api.alpha.linode.com/v4/"

    client := &http.Client{}

    req, err := http.NewRequest("GET", endpoint + snippet, nil)
    req.Header.Add("Authorization", "token " + c.Token)
    resp, err := client.Do(req)

    if (err != nil) {
        return err
    }

    defer resp.Body.Close()

    err = json.NewDecoder(resp.Body).Decode(&result)

    if (err != nil) {
        return err
    }

    return nil
}

type Datacenter struct {
    ID      string `json:"id"`
    Label   string `json:"label"`
    Country string `json:"country"`
}

type DatacentersResult struct {
    TotalPages   uint         `json:"total_pages"`
    TotalResults uint         `json:"total_results"`
    Page         uint         `json:"page"`
    Datacenters  []Datacenter `json:"datacenters"`
}

type Time struct {
    time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
    // TODO: actually parse the time
    *t = Time{time.Now()}
    return nil
}

type Distribution struct {
    ID                 string    `json:"id"`
    Created            Time      `json:"created"`
    Label              string    `json:"label"`
    MinimumStorageSize uint      `json:"minimum_storage_size"`
    Recommended        bool      `json:"recommended"`
    Vendor             string    `json:"vendor"`
    X64                bool      `json:"X64"`
}

type DistributionResult struct {
    TotalPages    uint           `json:"total_pages"`
    TotalResults  uint           `json:"total_results"`
    Page          uint           `json:"page"`
    Distributions []Distribution `json:"distributions"`
}

type Kernel struct {
    ID          string `json:"id"`
    Created     Time   `json:"created"`
    Deprecated  bool   `json:"deprecated"`
    Description string `json:"description"`
    Xen         bool   `json:"xen"`
    KVM         bool   `json:"kvm"`
    Label       string `json:"label"`
    Version     string `json:"version"`
    X64         bool   `json:"x64"`
}

type KernelResult struct {
    TotalPages   uint     `json:"total_pages"`
    TotalResults uint     `json:"total_results"`
    Page         uint     `json:"page"`
    Kernels      []Kernel `json:"kernels"`
}

type Alert struct {
    Enabled   bool `json:"enabled"`
    Threshold uint `json:"threshold"`
}

type IPv4 struct {
    Address string `json:"address"`
    Type    string `json:"type"`
    RDNS    string `json:"rdns"`
}

type IPv6 struct {
    Range string `json:"range"`
    Type  string `json:"type"`
}

type Linode struct {
    ID            uint             `json:"id"`
    Alerts        map[string]Alert `json:"alerts"`
    Backups       BackupInfo       `json:"backups"`
    Created       Time             `json:"created"`
    Datacenter    Datacenter       `json:"datacenter"`
    Distribution  Distribution     `json:"distribution"`
    Group         string           `json:"group"`
    IPv4          IPv4             `json:"ipv4"`
    IPv6          []IPv6           `json:"ipv6"`
    Label         string           `json:"label"`
    Type          []Service        `json:"type"`
    Status        string           `json:"status"`
    TotalTransfer uint             `json:"total_transfer"`
    Updated       Time             `json:"updated"`
}

type Disk struct {
    ID         uint   `json:"id"`
    Label      string `json:"label"`
    Status     string `json:"status"`
    Size       uint   `json:"size"`
    Filesystem string `json:"filesystem"`
    Created    Time   `json:"created"`
    Updated    Time   `json:"updated"`
}

type Backup struct {
    ID         uint       `json:"id"`
    Label      string     `json:"label"`
    Status     string     `json:"status"`
    Type       string     `json:"type"`
    Datacenter Datacenter `json:"datacenter"`
    Created    Time       `json:"created"`
    Updated    Time       `json:"updated"`
    Finished   Time       `json:"finished"`
    Configs    []string   `json:"configs"`
    Disks      []Disk     `json:"disks"`
}

type Schedule struct {
    Day    string `json:"day"`
    Window string `json:"window"`
}

type BackupInfo struct {
    Enabled    bool     `json:"enabled"`
    Schedule   Schedule `json:"schedule"`
    LastBackup Backup   `json:"last_backup"`
    Snapshot   Backup   `json:"snapshot"`
}

type Service struct {
    ID           string      `json:"id"`
    Storage      uint        `json:"storage"`
    BackupsPrice uint        `json:"backups_price"`
    HourlyPrice  uint        `json:"hourly_price"`
    Label        string      `json:"label"`
    MBitsOut     uint        `json:"mbits_out"`
    MonthlyPrice uint        `json:"monthly_price"`
    RAM          uint        `json:"ram"`
    ServiceType  string      `json:"service_type"`
    Transfer     uint        `json:"transfer"`
    VCPUs        uint        `json:"vcpus"`
}

type LinodeResult struct {
    TotalPages   uint     `json:"total_pages"`
    TotalResults uint     `json:"total_results"`
    Page         uint     `json:"page"`
    Linodes      []Linode `json:"linodes"`
}
