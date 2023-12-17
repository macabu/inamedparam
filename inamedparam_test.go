package inamedparam_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/macabu/inamedparam"
)

func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)

	analysistest.Run(t, testdata, inamedparam.Analyzer, "dummypkg")
}

func TestAnalyzerSkipSingleParam(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)

	analyzer := inamedparam.Analyzer
	analyzer.Flags.Set("skip-single-param", "true")

	analysistest.Run(t, testdata, analyzer, "params")
}
