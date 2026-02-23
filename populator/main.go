// populator reads a CSV file with headers "name", "count", and "address",
// splits names on " and ", creates an invitation per row, then creates a guest
// per split name linked to that invitation.
//
// Usage:
//
//	go run main.go -file guests.csv
//	go run main.go -file guests.csv -base http://localhost:8080
package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// ── API payload types ─────────────────────────────────────────────────────────

type createInvitationRequest struct {
	Address   string `json:"address"`
	NumGuests int    `json:"num_guests"`
}

type invitationResponse struct {
	ID int `json:"id"`
}

type createGuestRequest struct {
	Name         string `json:"name"`
	InvitationID int    `json:"invitation_id"`
}

// ── HTTP helpers ──────────────────────────────────────────────────────────────

func postJSON(client *http.Client, url string, payload any, out any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	resp, err := client.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("POST %s: %w", url, err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("POST %s returned %d: %s", url, resp.StatusCode, respBody)
	}

	if out != nil {
		if err := json.Unmarshal(respBody, out); err != nil {
			return fmt.Errorf("unmarshal response from %s: %w", url, err)
		}
	}
	return nil
}

// ── Main ──────────────────────────────────────────────────────────────────────

func main() {
	file := flag.String("file", "guests.csv", "path to the CSV file")
	base := flag.String("base", "http://localhost:8081", "base URL of the API (no trailing slash)")
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		log.Fatalf("open %s: %v", *file, err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.TrimLeadingSpace = true

	// Read and validate headers.
	headers, err := reader.Read()
	if err != nil {
		log.Fatalf("read headers: %v", err)
	}
	colIndex := map[string]int{}
	for i, h := range headers {
		colIndex[strings.ToLower(strings.TrimSpace(h))] = i
	}
	for _, required := range []string{"name", "count", "address"} {
		if _, ok := colIndex[required]; !ok {
			log.Fatalf("CSV is missing required column %q", required)
		}
	}

	client := &http.Client{}
	invURL := *base + "/api/v1/invitations"
	guestURL := *base + "/api/v1/guests"

	rowNum := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("row %d: read error: %v", rowNum, err)
		}
		rowNum++

		rawName := strings.TrimSpace(record[colIndex["name"]])
		rawCount := strings.TrimSpace(record[colIndex["count"]])
		address := strings.TrimSpace(record[colIndex["address"]])

		count, err := strconv.Atoi(rawCount)
		if err != nil {
			log.Printf("row %d: invalid count %q, skipping: %v", rowNum, rawCount, err)
			continue
		}

		// Split names on " and " (case-insensitive).
		parts := splitOnAnd(rawName)

		// 1. Create the invitation.
		var inv invitationResponse
		if err := postJSON(client, invURL, createInvitationRequest{
			Address:   address,
			NumGuests: count,
		}, &inv); err != nil {
			log.Printf("row %d: create invitation: %v", rowNum, err)
			continue
		}
		fmt.Printf("row %d: created invitation id=%d for %q\n", rowNum, inv.ID, address)

		// 2. Create a guest for each name.
		for _, name := range parts {
			if err := postJSON(client, guestURL, createGuestRequest{
				Name:         name,
				InvitationID: inv.ID,
			}, nil); err != nil {
				log.Printf("row %d: create guest %q: %v", rowNum, name, err)
				continue
			}
			fmt.Printf("row %d:   created guest %q -> invitation %d\n", rowNum, name, inv.ID)
		}
	}
}

// splitOnAnd splits s on the word "and" (surrounded by whitespace) and
// returns a slice of trimmed, non-empty name strings.
func splitOnAnd(s string) []string {
	// Split on " and " variants, case-insensitive.
	lower := strings.ToLower(s)
	var parts []string
	for {
		idx := strings.Index(lower, " and ")
		if idx == -1 {
			parts = append(parts, strings.TrimSpace(s))
			break
		}
		parts = append(parts, strings.TrimSpace(s[:idx]))
		s = s[idx+5:]
		lower = lower[idx+5:]
	}
	// Filter empty strings.
	result := parts[:0]
	for _, p := range parts {
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}
