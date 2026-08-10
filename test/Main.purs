module Test.Main where

import Prelude

import Effect (Effect)
import Test.Spec.Reporter (consoleReporter)
import Test.Spec.Runner.Node (runSpecAndExitProcess)

import Test.BasicsSpec as BasicsSpec
import Test.ErrorsSpec as ErrorsSpec
import Test.GenericsSpec as GenericsSpec
import Test.WriteViaSpec as WriteViaSpec

main ∷ Effect Unit
main = runSpecAndExitProcess [ consoleReporter ] do
  BasicsSpec.spec
  ErrorsSpec.spec
  GenericsSpec.spec
  WriteViaSpec.spec

