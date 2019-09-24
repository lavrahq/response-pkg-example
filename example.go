package example

import (
	"fmt"

	"github.com/lavrahq/response-pkg/sdk"
)

type examplePackage struct {
}

// GetName returns the human-friendly name of the package.
func (p *examplePackage) GetName() string {
	return "Base"
}

// GetAuthor returns the Marketplace username of the author.
func (p *examplePackage) GetAuthor() string {
	return "lavrahq"
}

// GetVersion returns the package version.
func (p *examplePackage) GetVersion() string {
	return "0.0.1"
}

func (p *examplePackage) Install(sdk *sdk.SDK) error {
	fmt.Println("installed")

	return nil
}

func init() {
	sdk.RegisterPackage(&examplePackage{})
}
