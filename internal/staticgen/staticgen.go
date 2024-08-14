package staticgen

import (
	"context"
	"fmt"
	"goth/internal/templates"
	"os"
	"path/filepath"
)

func GenerateStaticFiles(outputDir string) error {
	// Ensure the output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Create the index.html file in the output directory
	indexPath := filepath.Join(outputDir, "index.html")
	f, err := os.Create(indexPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer f.Close()

	// Render the template to the file
	err = templates.Layout().Render(context.Background(), f)
	if err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}

	return nil
}
