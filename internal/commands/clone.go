package commands

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Clone(repoURL, destinationPath string) {
	fmt.Println("Cloning from", repoURL, "to", destinationPath)

	owner, repo, err := parseGitHubURL(repoURL)
	if err != nil {
		fmt.Println("Error parsing GitHub URL:", err)
		return
	}

	if err := os.MkdirAll(destinationPath, 0755); err != nil {
		fmt.Println("Error creating destination directory:", err)
		return
	}

	zipURL := fmt.Sprintf("https://github.com/%s/%s/archive/refs/heads/main.zip", owner, repo)
	zipFilePath := filepath.Join(destinationPath, "repo.zip")
	if err := downloadFile(zipURL, zipFilePath); err != nil {
		fmt.Println("Error downloading repository:", err)
		return
	}

	if err := extractZip(zipFilePath, destinationPath); err != nil {
		fmt.Println("Error extracting ZIP file:", err)
		return
	}

	os.Remove(zipFilePath)
	fmt.Println("Clone completed")
}

func parseGitHubURL(url string) (string, string, error) {
	parts := strings.Split(strings.TrimPrefix(url, "https://github.com/"), "/")
	if len(parts) < 2 {
		return "", "", fmt.Errorf("invalid GitHub URL")
	}
	return parts[0], parts[1], nil
}

func downloadFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func extractZip(zipFilePath, destinationPath string) error {
	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, file := range r.File {
		filePath := filepath.Join(destinationPath, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		rc, err := file.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
