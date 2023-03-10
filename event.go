package main

import (
    "time"
)

type Event struct {
    ID                      string      `json:"id"`
    Name                    string      `json:"name"`
    Description             string      `json:"description"`
    CreatedBy               string      `json:"createdBy"`
    StartTime               time.Time   `json:"startTime"`
    Duration                int         `json:"duration"`
    InvitedUsers            []string    `json:"invitedUsers"`
    AcceptedUsers           []string    `json:"acceptedUsers"`
    AcceptanceRatio         float64     `json:"acceptanceRatio"`
}