package pkg

import (
	"fmt"
	"os"

	"github.com/ossf/scorecard/v4/docs/checks"
	"github.com/ossf/scorecard/v4/log"
	"github.com/ossf/scorecard/v4/options"
	scpkg "github.com/ossf/scorecard/v4/pkg"
)

// ChangeType is the change type (added, updated, removed) of a dependency.
type ChangeType string

const (
	// Added suggests the dependency is a newly added one.
	Added ChangeType = "added"
	// Updated suggests the dependency is updated from an old version.
	Updated ChangeType = "updated"
	// Removed suggests the dependency is removed.
	Removed ChangeType = "removed"
)

// IsValid determines if a ChangeType is valid.
func (ct *ChangeType) IsValid() bool {
	switch *ct {
	case Added, Updated, Removed:
		return true
	default:
		return false
	}
}

// ScorecardResultWithError is used for the dependency-diff module to record the scorecard result
// and a potential error field if the Scorecard run fails.
type ScorecardResultWithError struct {
	// ScorecardResult is the scorecard result for the dependency repo.
	ScorecardResult *scpkg.ScorecardResult

	// Error is an error returned when running the scorecard checks. A nil Error indicates the run succeeded.
	Error error
}

// DependencyCheckResult is the dependency structure used in the returned results.
type DependencyCheckResult struct {
	// ChangeType indicates whether the dependency is added, updated, or removed.
	ChangeType *ChangeType

	// Package URL is a short link for a package.
	PackageURL *string

	// SourceRepository is the source repository URL of the dependency.
	SourceRepository *string

	// ManifestPath is the path of the manifest file of the dependency, such as go.mod for Go.
	ManifestPath *string

	// Ecosystem is the name of the package management system, such as NPM, GO, PYPI.
	Ecosystem *string

	// Version is the package version of the dependency.
	Version *string

	// ScorecardResultWithError is the scorecard checking results of the dependency.
	ScorecardResultWithError ScorecardResultWithError

	// Name is the name of the dependency.
	Name string
}

// FormatDependencydiffResults formats dependencydiff results.
func FormatDependencydiffResults(
	opts *options.Options,
	depdiffResults []DependencyCheckResult,
	doc checks.Doc,
) error {
	err := DependencydiffResultsAsJSON(depdiffResults, log.ParseLevel(opts.LogLevel), doc, os.Stdout)
	if err != nil {
		return fmt.Errorf("failed to output dependencydiff results: %w", err)
	}
	return nil
}
