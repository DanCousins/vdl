package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/bodgit/sevenzip"
)

func extractArchive(archivePath string) error {

	r, err := sevenzip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if strings.HasSuffix(f.Name, ".bin") || strings.HasSuffix(f.Name, ".cue") {

			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			outFile, err := os.Create("/mnt/SDCARD/Roms/PS/" + strings.Split(f.Name, "/")[1])
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return err
			}

			fmt.Printf("Extracted: %s\n", "/mnt/SDCARD/Roms/PS/"+f.Name)
		}
	}

	return nil
}

func parseRom(vaultId string) (mediaId string, romFolder string) {

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	res, err := c.Get("https://vimm.net/vault/" + vaultId)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	mediaId = doc.Find("input[name='mediaId']").AttrOr("value", "")

	img := doc.Find("input[name='system']")
	if img.Length() > 0 {
		console, _ := img.Attr("value")

		if console == "GB" {
			romFolder = "GB/"
		} else if console == "GBC" {
			romFolder = "GBC/"
		} else if console == "GBA" {
			romFolder = "GBA/"
		} else if console == "DS" {
			romFolder = "NDS/"
		} else if console == "Atari2600" {
			romFolder = "ATARI/"
		} else if console == "Atari5200" {
			romFolder = "FIFTYTWOHUNDRED/"
		} else if console == "NES" {
			romFolder = "FC/"
		} else if console == "SMS" {
			romFolder = "MS/"
		} else if console == "Atari7800" {
			romFolder = "SEVENTYEIGHTHUNDRED/"
		} else if console == "Genesis" {
			romFolder = "MD/"
		} else if console == "SNES" {
			romFolder = "SFC/"
		} else if console == "32X" {
			romFolder = "THIRTYTWOX/"
		} else if console == "PS1" {
			romFolder = "PS/"
		} else if console == "Lynx" {
			romFolder = "LYNX/"
		} else if console == "GG" {
			romFolder = "GG/"
		} else if console == "VB" {
			romFolder = "VB/"
		} else {
			fmt.Println("No console match - saving ROM to /mnt/SDCARD/Roms/.")
		}

	} else {
		fmt.Println("No console found - saving ROM to /mnt/SDCARD/Roms/.")
	}

	return mediaId, romFolder
}

func downloadRom(filepath string, romUrl string, downloadUrl string) (err error) {

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	fmt.Println("Downloading... this may take some time for larger game files... (PS1, NDS)")

	req, err := http.NewRequest("GET", downloadUrl, nil)
	req.Header.Add("Referer", romUrl)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:123.0) Gecko/20100101 Firefox/123.0")
	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Error: %s", resp.Status)
	}

	gameName := strings.Split(resp.Header.Get("Content-Disposition"), "\"")[1]

	filepath = filepath + gameName

	rom, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer rom.Close()

	_, err = io.Copy(rom, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Download complete -", filepath)

	if strings.Contains(filepath, "/PS/") {
		fmt.Println("Extracting Archive... Please wait...")
		extractArchive(filepath)
		fmt.Println("Extraction Complete")
		defer os.Remove(filepath) // remove a single file
	}

	return nil
}

func welcome() {
	fmt.Println("Welcome to the (unofficial) Miyoo interface for Vimm's Lair!")
}

func main() {
	welcome()
	fmt.Print("Please input the vault ID of the rom you wish to download: ")
	var inputId string
	fmt.Scanln(&inputId)

	mediaId, romFolder := parseRom(inputId)
	romFilepath := "/mnt/SDCARD/Roms/" + romFolder
	downloadRom(romFilepath, "https://vimm.net/vault/"+inputId, "https://download3.vimm.net/download/?mediaId="+mediaId)
}
