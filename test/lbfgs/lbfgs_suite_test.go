package lbfgs_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "testing"
)

func TestLBFGS(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "LBFGS Suite")
}
