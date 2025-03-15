package inamedparam_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/macabu/inamedparam"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), inamedparam.Analyzer, "dummypkg")
}

func TestAnalyzerSkipSingleParam(t *testing.T) {
	analyzer := inamedparam.Analyzer

	err := analyzer.Flags.Set("skip-single-param", "true")
	if err != nil {
		t.Fatalf("failed to set skip-single-param flag: %v", err)
	}

	analysistest.Run(t, analysistest.TestData(), analyzer, "params")
}
