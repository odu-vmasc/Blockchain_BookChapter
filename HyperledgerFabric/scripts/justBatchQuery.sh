#!/bin/bash

# Importing useful functions for cc testing
if [ -f ./func.sh ]; then
 source ./func.sh
elif [ -f scripts/func.sh ]; then
 source scripts/func.sh
fi

#tot=5
#val=90
#end=$((val-5))

echo_b "Channel name : "$CHANNEL_NAME

echo_b "====================Query the existing value of a $1===================================="
multQuery 0 $1 $2

echo
echo_g "=====================All GOOD, MVE Test completed ===================== "
echo
exit 0
