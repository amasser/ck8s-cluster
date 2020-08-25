#!/bin/bash
#This script should contain extra tests (not the status command in ckctl) to run after a cluster is setup.

set -e

here="$(dirname "$(readlink -f "$0")")"

"${here}/e2e-tests.bash"
