module Test.Main where

import Prelude

import Effect (Effect)
import Effect.Aff (launchAff_)
import Test.Spec.Reporter (consoleReporter)
import Test.Spec.Runner (runSpec)
import Test.BasicsSpec as BasicsSpec
import Test.ErrorsSpec as ErrorsSpec
import Test.GenericsSpec as GenericsSpec
import Test.WriteViaSpec as WriteViaSpec

main ∷ Effect Unit
main = launchAff_ $ runSpec [ consoleReporter ] do
  BasicsSpec.spec
  ErrorsSpec.spec
  GenericsSpec.spec
  WriteViaSpec.spec

