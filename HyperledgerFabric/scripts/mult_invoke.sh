#!/bin/bash

# Importing useful functions for cc testing
if [ -f ./func.sh ]; then
 source ./func.sh
elif [ -f scripts/func.sh ]; then
 source scripts/func.sh
fi

tot=1
val=100
iter=1
end=1

#ids = 
idcount=1;
fulltime=$(date +%s)
echo_b "Channel name : "$CHANNEL_NAME
	
echo_b "====================Query the existing value of a===================================="
#chaincodeQuery 0 $val
	
for q in {2..1000}
do
	currID="a$q"
	bankID="bank"
	echo_b "=====================Invoke a transaction to transfer 1 from "$currID" to "$bankID"=================="
	chaincodeInvokeMult 0 $currID $bankID
	val=$((val - 1))
	#sleep 1
done
echo_b "=====================Check if the result of a is $val==================================="

#chaincodeCheckMult 0 a100 $val	
iter=$((iter + 1))
echo "testcomplete,$(($(date +%s)-fulltime))" >> report.txt
#echo "$cf,$2,$(($(date +%s)-starttime))" >> report.txt
echo
echo_g "=====================All GOOD, MVE Test completed ===================== "
echo
exit 0
