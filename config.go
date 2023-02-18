package main

import (
    "os"
    "log"
    "fmt"
    "encoding/json"
)

type Config struct {
    filename string
    data map[string]interface{}
}

func NewConfig(filename string) (Config, error) {
    cfg := Config{}
    cfg.filename = filename

    file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
    check(err, "")
    file.Close()

    content, err := os.ReadFile(filename)
    check(err, "")

    if string(content) == "" {
        os.WriteFile(filename, []byte("{}"), 0644)
    }

    // Re-read the file, as we've done conditional wrote
    // of {} above
    content, err = os.ReadFile(filename)
    check(err, "")

    log.Printf("Auto-load config file: %s", filename)
    err = json.Unmarshal(content, &cfg.data)
    if err != nil {
        return cfg, fmt.Errorf("Failed to load json from file: %w", err)
    }
    return cfg, nil
}

func (cfg Config) ReadConfig() (error) {
    file, err := os.ReadFile(cfg.filename)
    if err != nil {
        return err
    }
    json.Unmarshal(file, &cfg.data)
    return nil
}

func (cfg Config) WriteConfig() (error) {
    json_bytes, err := json.Marshal(cfg.data)
    if err != nil {
        return err
    }

    err = os.WriteFile(cfg.filename, json_bytes, 0644)
    if err != nil {
        return err
    }

    return nil
}
