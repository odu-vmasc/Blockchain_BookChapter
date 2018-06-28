#!/bin/bash

# Importing useful functions for cc testing
if [ -f ./func.sh ]; then
 source ./func.sh
elif [ -f scripts/func.sh ]; then
 source scripts/func.sh
fi

tot=1000
val=$1
iter=1
end=1
fulltime=$(date +%s)
#while [ $iter -lt $tot ]; do
#	for i in {1..5}
#	do
		echo_b "Channel name : "$CHANNEL_NAME
		
		#echo_b "====================Query the existing value of a===================================="
		#chaincodeQuery 0 $val
		
		if [ $end = 1 ]; then
			echo_b "=====================Invoke a transaction to transfer 1 from a to b=================="
			chaincodeAutoVoke 0
			val=$((val - 1))
			#sleep 1
		else
			echo_b "=====================Invoke a transaction to transfer 1 from b to a=================="
			chaincodAutoVoke 0
			val=$((val + 1))
			#sleep 1
		fi
		echo_b "=====================Check if the result of a is $val==================================="
		#chaincodeCheck 0 $val	
		iter=$((iter + 1))
#	done
	if [ $end = 1 ]; then
		end=2
	else
		end=1
	fi
#done
echo "testcomplete,$(($(date +%s)-fulltime))" >> report.txt
#echo "$cf,$2,$(($(date +%s)-starttime))" >> report.txt
echo
echo_g "=====================All GOOD, MVE Test completed ===================== "
echo
exit 0
