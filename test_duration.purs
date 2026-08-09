module Test.Duration where
import Prelude
import Effect (Effect)
import Effect.Console (log)
import Data.Time.Duration (Milliseconds(..))
import Yoga.JSON (writeJSON)

main :: Effect Unit
main = do
  log $ writeJSON (Milliseconds 16.67)
