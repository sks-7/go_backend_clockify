package dbconfig

import (
    "clockfy_backend/prisma/db"
    "log"
)

var Client *db.PrismaClient

func InitDB() error {
    Client = db.NewClient()
    if err := Client.Prisma.Connect(); err != nil {
        return err
    }
    return nil
}

func CloseDB() {
    if err := Client.Prisma.Disconnect(); err != nil {
        log.Fatalf("Failed to disconnect from the database: %v", err)
    }
}
