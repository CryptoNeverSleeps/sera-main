#!/bin/bash
rm -rf /tmp/serad-temp

serad --devnet --appdir=/tmp/serad-temp --profile=6061 &
SEDRAD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:22611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $SEDRAD_PID

wait $SEDRAD_PID
SEDRAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "serad exit code: $SEDRAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SEDRAD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
